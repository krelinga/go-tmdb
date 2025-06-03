package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"iter"
	"net/http"
	"net/url"

	"github.com/krelinga/go-tmdb/internal/raw"
)

type SearchMovieOptions struct {
	IncludeAdult       bool
	Language           Language
	PrimaryReleaseYear int
	Region             string
	Year               int
}

func SearchMovie(ctx context.Context, c *Client, query string, options *SearchMovieOptions) iter.Seq2[Movie, error] {
	return func(yield func(Movie, error) bool) {
		for page := 1; ; page++ {
			params := url.Values{
				"query": []string{query},
				"page":  []string{fmt.Sprint(page)},
			}
			if options != nil {
				if options.IncludeAdult {
					params.Set("include_adult", fmt.Sprint(options.IncludeAdult))
				}
				if options.Language != "" {
					params.Set("language", string(options.Language))
				}
				if options.PrimaryReleaseYear > 0 {
					params.Set("primary_release_year", fmt.Sprint(options.PrimaryReleaseYear))
				}
				if options.Region != "" {
					params.Set("region", options.Region)
				}
				if options.Year > 0 {
					params.Set("year", fmt.Sprint(options.Year))
				}
			}

			theUrl := &url.URL{
				Path:     "3/search/movie",
				RawQuery: params.Encode(),
			}
			c.prepUrl(theUrl)
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, theUrl.String(), nil)
			if err != nil {
				yield(nil, err)
				return
			}
			c.prepRequest(req)
			if ctx.Err() != nil {
				return
			}
			resp, err := c.httpClient().Do(req)
			if err != nil {
				yield(nil, err)
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				yield(nil, fmt.Errorf("TMDB API returned status code %d", resp.StatusCode))
				return
			}

			decoder := json.NewDecoder(resp.Body)
			var result raw.SearchMovie
			if err := decoder.Decode(&result); err != nil {
				yield(nil, fmt.Errorf("decoding search result: %w", err))
				return
			}
			result.SetDefaults()
			for _, smrMovie := range result.Results {
				m := &movie{
					id: MovieId(smrMovie.Id),
					MovieParts: &searchMovieResultParts{
						raw: smrMovie,
					},
				}
				if !yield(m, nil) {
					return
				}
			}
			if page >= result.TotalPages {
				return
			}
		}
	}
}

type searchMovieResultParts struct {
	raw *raw.SearchMovieResult
	movieNoParts
}

func (s *searchMovieResultParts) Adult() bool {
	return *s.raw.Adult
}
