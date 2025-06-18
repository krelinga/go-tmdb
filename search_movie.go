package tmdb

import (
	"context"
	"fmt"
	"iter"
	"net/url"

	"github.com/krelinga/go-tmdb/internal/raw"
)

type SearchMovieResult struct {
	Movie        *Movie
	Page         int
	TotalPages   int
	TotalResults int
}

// The resulting Movie instances will have the following fields set:
// - Id
// - Adult
// - Backdrop
// - Genres.Id (NOT Genres.Name)
// - OriginalLanguage
// - OriginalTitle
// - Popularity
// - Poster
// - ReleaseDate
// - Title
// - Video
// - VoteAverage
// - VoteCount
func SearchMovie(ctx context.Context, c *Client, query string, options ...Option) iter.Seq2[*SearchMovieResult, error] {
	return func(yield func(*SearchMovieResult, error) bool) {
		callOpts := c.globalOpts
		callOpts.apply(options)

		if callOpts.StartPage == nil {
			callOpts.StartPage = NewPtr(1)
		}

		for page := *callOpts.StartPage; ; page++ {
			params := url.Values{
				"query": []string{query},
				"page":  []string{fmt.Sprint(page)},
			}
			callOpts.applyIncludeAdult(params)
			callOpts.applyLanguage(params)

			var result raw.SearchMovie
			if err := get(ctx, c, "search/movie", params, callOpts, &result); err != nil {
				yield(nil, fmt.Errorf("searching for movie %q: %w", query, err))
				return
			}

			for _, smrMovie := range result.Results {
				genres := make([]Genre, len(smrMovie.GenreIds))
				for i, id := range smrMovie.GenreIds {
					genres[i] = Genre{
						Key: GenreId(id),
					}
				}
				out := &SearchMovieResult{
					Page:         page,
					TotalPages:   result.TotalPages,
					TotalResults: result.TotalResults,
					Movie: &Movie{
						Key:              MovieId(smrMovie.Id),
						Adult:            smrMovie.Adult,
						Backdrop:         NewPtr(Image(smrMovie.BackdropPath)),
						Genres:           genres,
						OriginalLanguage: &smrMovie.OriginalLanguage,
						OriginalTitle:    &smrMovie.OriginalTitle,
						Overview:         &smrMovie.Overview,
						Popularity:       &smrMovie.Popularity,
						Poster:           NewPtr(Image(smrMovie.PosterPath)),
						ReleaseDate:      NewPtr(DateYYYYMMDD(smrMovie.ReleaseDate)),
						Title:            &smrMovie.Title,
						VoteAverage:      &smrMovie.VoteAverage,
						VoteCount:        &smrMovie.VoteCount,
					},
				}
				if !yield(out, nil) {
					return
				}
			}
			overLimitPage := callOpts.LimitPage != nil && page >= *callOpts.LimitPage
			if page >= result.TotalPages || overLimitPage {
				return
			}
		}
	}
}
