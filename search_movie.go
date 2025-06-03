package tmdb

import (
	"context"
	"fmt"
	"iter"
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
			var language Language
			if options != nil {
				if options.IncludeAdult {
					params.Set("include_adult", fmt.Sprint(options.IncludeAdult))
				}
				if options.Language != "" {
					params.Set("language", string(options.Language))
					language = options.Language
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
			var result raw.SearchMovie
			if err := get(ctx, c, "search/movie", params, &result); err != nil {
				yield(nil, fmt.Errorf("searching for movie %q: %w", query, err))
				return
			}

			for _, smrMovie := range result.Results {
				m := &movie{
					client:   c,
					id:       MovieId(smrMovie.Id),
					language: language,
					MovieData: &searchMovieResultData{
						raw:       smrMovie,
						MovieData: movieNoData{},
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

type searchMovieResultData struct {
	client *Client
	raw *raw.SearchMovieResult
	MovieData
}

func (s *searchMovieResultData) upgrade(in *getMovieData) MovieData {
	s.MovieData = s.MovieData.upgrade(in)
	return s
}

func (s *searchMovieResultData) Adult() bool {
	return *s.raw.Adult
}

func (s *searchMovieResultData) Backdrop() Image {
	return image{
		client:   s.client,
		raw:      s.raw.BackdropPath,
	}
}
