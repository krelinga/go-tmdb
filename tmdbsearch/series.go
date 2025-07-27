package tmdbsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/krelinga/go-tmdb/internal/util"
)

type FindSeriesOptions struct {
	FirstAirDateYear int32
	IncludeAdult     bool
	Language         string
	Page             int32
	Year             int32
}

func FindSeries(ctx context.Context, query string, options FindSeriesOptions) (*http.Response, error) {
	rb := util.NewRequestBuilder().
		SetPath("/3/search/tv").
		SetValueString("query", query).
		SetValueInt32("first_air_date_year", options.FirstAirDateYear).
		SetValueString("language", options.Language).
		SetValueInt32("page", options.Page).
		SetValueInt32("year", options.Year)

	if options.FirstAirDateYear > 0 {
		rb.SetValueString("first_air_date_year", fmt.Sprint(options.FirstAirDateYear))
	}
	if options.IncludeAdult {
		rb.SetValueString("include_adult", "true")
	}
	if options.Page > 0 {
		rb.SetValueString("page", fmt.Sprint(options.Page))
	}
	if options.Year > 0 {
		rb.SetValueString("year", fmt.Sprint(options.Year))
	}

	return rb.Do(ctx)
}

func ParseFindSeriesReply(httpReply *http.Response) (*FindSeriesReply, error) {
	defer httpReply.Body.Close()
	if httpReply.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", httpReply.StatusCode)
	}
	if httpReply.Header.Get("Content-Type") != "application/json;charset=utf-8" {
		return nil, fmt.Errorf("unexpected content type: %s", httpReply.Header.Get("Content-Type"))
	}

	reply := &FindSeriesReply{}
	if err := json.NewDecoder(httpReply.Body).Decode(reply); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return reply, nil
}

type FindSeriesReply struct {
	Page         *int32    `json:"page"`
	Series       []*Series `json:"results"`
	TotalPages   *int32    `json:"total_pages"`
	TotalResults *int32    `json:"total_results"`
}

func (sr *FindSeriesReply) SetDefaults() {
	if sr == nil {
		return
	}
	util.SetIfNil(&sr.Page, 0)
	for _, series := range sr.Series {
		series.SetDefaults()
	}
	util.SetIfNil(&sr.TotalPages, 0)
	util.SetIfNil(&sr.TotalResults, 0)
}

func (sr *FindSeriesReply) String() string {
	if sr == nil {
		return "<nil>"
	}
	var b strings.Builder
	b.WriteString("{")
	fmt.Fprintf(&b, "Page: %s", util.FmtOrNil(sr.Page))
	fmt.Fprintf(&b, " Series: %v", sr.Series)
	fmt.Fprintf(&b, " TotalPages: %s", util.FmtOrNil(sr.TotalPages))
	fmt.Fprintf(&b, " TotalResults: %s", util.FmtOrNil(sr.TotalResults))
	b.WriteString("}")
	return b.String()
}

type Series struct {
	Adult            *bool    `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	GenreIDs         []int32  `json:"genre_ids"`
	ID               *int32   `json:"id"`
	OriginCountries  []string `json:"origin_country"`
	OriginalLanguage *string  `json:"original_language"`
	OriginalName     *string  `json:"original_name"`
	Overview         *string  `json:"overview"`
	Popularity       *float32 `json:"popularity"`
	PosterPath       *string  `json:"poster_path"`
	FirstAirDate     *string  `json:"first_air_date"`
	Name             *string  `json:"name"`
	VoteAverage      *float32 `json:"vote_average"`
	VoteCount        *int32   `json:"vote_count"`
}

func (s *Series) SetDefaults() {
	if s == nil {
		return
	}
	util.SetIfNil(&s.Adult, true)
	util.SetIfNil(&s.ID, 0)
	util.SetIfNil(&s.Popularity, 0.0)
	util.SetIfNil(&s.VoteAverage, 0.0)
	util.SetIfNil(&s.VoteCount, 0)
}

func (s *Series) String() string {
	if s == nil {
		return "<nil>"
	}
	var b strings.Builder
	b.WriteString("{")
	fmt.Fprintf(&b, "Adult: %s", util.FmtOrNil(s.Adult))
	fmt.Fprintf(&b, " BackdropPath: %s", util.FmtOrNil(s.BackdropPath))
	fmt.Fprintf(&b, " GenreIDs: %v", s.GenreIDs)
	fmt.Fprintf(&b, " ID: %s", util.FmtOrNil(s.ID))
	fmt.Fprintf(&b, " OriginCountries: %v", s.OriginCountries)
	fmt.Fprintf(&b, " OriginalLanguage: %s", util.FmtOrNil(s.OriginalLanguage))
	fmt.Fprintf(&b, " OriginalName: %s", util.FmtOrNil(s.OriginalName))
	fmt.Fprintf(&b, " Overview: %s", util.FmtOrNil(s.Overview))
	fmt.Fprintf(&b, " Popularity: %s", util.FmtOrNil(s.Popularity))
	fmt.Fprintf(&b, " PosterPath: %s", util.FmtOrNil(s.PosterPath))
	fmt.Fprintf(&b, " FirstAirDate: %s", util.FmtOrNil(s.FirstAirDate))
	fmt.Fprintf(&b, " Name: %s", util.FmtOrNil(s.Name))
	fmt.Fprintf(&b, " VoteAverage: %s", util.FmtOrNil(s.VoteAverage))
	fmt.Fprintf(&b, " VoteCount: %s", util.FmtOrNil(s.VoteCount))
	b.WriteString("}")
	return b.String()
}
