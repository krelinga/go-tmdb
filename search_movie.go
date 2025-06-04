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
	Language           LanguageId
	PrimaryReleaseYear int
	Region             string
	Year               int
}

// The resulting Movie instances will have the following data available:
// - Adult()
// - Backdrop()
// - GenreIds()
// - OriginalLanguage()
// - OriginalTitle()
// - Popularity()
// - Poster()
func SearchMovie(ctx context.Context, c *Client, query string, options *SearchMovieOptions) iter.Seq2[Movie, error] {
	return func(yield func(Movie, error) bool) {
		for page := 1; ; page++ {
			params := url.Values{
				"query": []string{query},
				"page":  []string{fmt.Sprint(page)},
			}
			var language LanguageId
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
	raw    *raw.SearchMovieResult
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
		client: s.client,
		raw:    s.raw.BackdropPath,
	}
}

func (s *searchMovieResultData) GenreIds() iter.Seq[GenreId] {
	return func(yield func(GenreId) bool) {
		for _, id := range s.raw.GenreIds {
			if !yield(GenreId(id)) {
				return
			}
		}
	}
}

func (s *searchMovieResultData) OriginalLanguage() LanguageId {
	return LanguageId(s.raw.OriginalLanguage)
}

func (s *searchMovieResultData) OriginalTitle() string {
	return s.raw.OriginalTitle
}

func (s *searchMovieResultData) Overview() string {
	return s.raw.Overview
}

func (s *searchMovieResultData) Popularity() float64 {
	return s.raw.Popularity
}

func (s *searchMovieResultData) Poster() Image {
	return image{
		client: s.client,
		raw:    s.raw.PosterPath,
	}
}

func (s *searchMovieResultData) ReleaseDate() Date {
	return newDateYYYYMMDD(s.raw.ReleaseDate)
}
