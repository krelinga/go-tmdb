package search

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/krelinga/go-tmdb/internal/util"
)

type MovieOptions struct {
	Key                string
	ReadAccessToken    string
	IncludeAdult       bool
	Language           string
	PrimaryReleaseYear string
	Page               int32
	Region             string
	Year               string
}

func Movie(ctx context.Context, client *http.Client, query string, options MovieOptions) (*MovieReply, error) {
	values := url.Values{
		"query": []string{query},
	}
	util.SetIfNotZero(&values, "api_key", options.Key)
	util.SetIfNotZero(&values, "include_adult", options.IncludeAdult)
	util.SetIfNotZero(&values, "language", options.Language)
	util.SetIfNotZero(&values, "primary_release_year", options.PrimaryReleaseYear)
	util.SetIfNotZero(&values, "page", options.Page)
	util.SetIfNotZero(&values, "region", options.Region)
	util.SetIfNotZero(&values, "year", options.Year)
	url := &url.URL{
		Scheme:   "https",
		Host:     "api.themoviedb.org",
		Path:     "/3/search/movie",
		RawQuery: values.Encode(),
	}
	request := &http.Request{
		Method: http.MethodGet,
		URL:    url,
	}
	util.SetAuthIfNotZero(request, options.ReadAccessToken)
	httpReply, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer httpReply.Body.Close()
	if httpReply.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", httpReply.StatusCode)
	}
	if httpReply.Header.Get("Content-Type") != "application/json;charset=utf-8" {
		return nil, fmt.Errorf("unexpected content type: %s", httpReply.Header.Get("Content-Type"))
	}

	reply := &MovieReply{}
	if err := json.NewDecoder(httpReply.Body).Decode(reply); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return reply, nil
}

type MovieReply struct {
	Page         *int32         `json:"page"`
	MovieResults []*MovieResult `json:"results"`
	TotalPages   *int32         `json:"total_pages"`
	TotalResults *int32         `json:"total_results"`
}

func (mr *MovieReply) SetDefaults() {
	if mr == nil {
		return
	}
	util.SetIfNil(&mr.Page, 0)
	for _, movie := range mr.MovieResults {
		movie.SetDefaults()
	}
	util.SetIfNil(&mr.TotalPages, 0)
	util.SetIfNil(&mr.TotalResults, 0)
}

func (mr *MovieReply) String() string {
	if mr == nil {
		return "<nil>"
	}
	var b strings.Builder
	b.WriteString("{")
	fmt.Fprintf(&b, "Page: %s", util.FmtOrNil(mr.Page))
	fmt.Fprintf(&b, " MovieResults: %v", mr.MovieResults)
	fmt.Fprintf(&b, " TotalPages: %s", util.FmtOrNil(mr.TotalPages))
	fmt.Fprintf(&b, " TotalResults: %s", util.FmtOrNil(mr.TotalResults))
	b.WriteString("}")
	return b.String()
}

type MovieResult struct {
	Adult            *bool    `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	GenreIDs         []int32  `json:"genre_ids"`
	ID               *int32   `json:"id"`
	OriginalLanguage *string  `json:"original_language"`
	OriginalTitle    *string  `json:"original_title"`
	Overview         *string  `json:"overview"`
	Popularity       *float32 `json:"popularity"`
	PosterPath       *string  `json:"poster_path"`
	ReleaseDate      *string  `json:"release_date"`
	Title            *string  `json:"title"`
	Video            *bool    `json:"video"`
	VoteAverage      *float32 `json:"vote_average"`
	VoteCount        *int32   `json:"vote_count"`
}

func (mr *MovieResult) SetDefaults() {
	if mr == nil {
		return
	}
	util.SetIfNil(&mr.Adult, true)
	util.SetIfNil(&mr.ID, 0)
	util.SetIfNil(&mr.Popularity, 0.0)
	util.SetIfNil(&mr.Video, true)
	util.SetIfNil(&mr.VoteAverage, 0.0)
	util.SetIfNil(&mr.VoteCount, 0)
}

func (mr *MovieResult) String() string {
	if mr == nil {
		return "<nil>"
	}
	var b strings.Builder
	b.WriteString("{")
	fmt.Fprintf(&b, "Adult: %s", util.FmtOrNil(mr.Adult))
	fmt.Fprintf(&b, " BackdropPath: %s", util.FmtOrNil(mr.BackdropPath))
	fmt.Fprintf(&b, " GenreIDs: %v", mr.GenreIDs)
	fmt.Fprintf(&b, " ID: %s", util.FmtOrNil(mr.ID))
	fmt.Fprintf(&b, " OriginalLanguage: %s", util.FmtOrNil(mr.OriginalLanguage))
	fmt.Fprintf(&b, " OriginalTitle: %s", util.FmtOrNil(mr.OriginalTitle))
	fmt.Fprintf(&b, " Overview: %s", util.FmtOrNil(mr.Overview))
	fmt.Fprintf(&b, " Popularity: %s", util.FmtOrNil(mr.Popularity))
	fmt.Fprintf(&b, " PosterPath: %s", util.FmtOrNil(mr.PosterPath))
	fmt.Fprintf(&b, " ReleaseDate: %s", util.FmtOrNil(mr.ReleaseDate))
	fmt.Fprintf(&b, " Title: %s", util.FmtOrNil(mr.Title))
	fmt.Fprintf(&b, " Video: %s", util.FmtOrNil(mr.Video))
	fmt.Fprintf(&b, " VoteAverage: %s", util.FmtOrNil(mr.VoteAverage))
	fmt.Fprintf(&b, " VoteCount: %s", util.FmtOrNil(mr.VoteCount))
	b.WriteString("}")
	return b.String()
}
