package tmdb

import (
	"context"
	"fmt"
	"iter"
	"net/url"

	"github.com/krelinga/go-tmdb/internal/raw"
)

type SearchMovieResult struct {
	rawResult    *raw.SearchMovieResult
	Page         int
	TotalPages   int
	TotalResults int
}

func (smr *SearchMovieResult) Upsert(g *Graph) *Movie {
	movie := g.EnsureMovie(MovieId(smr.rawResult.Id))
	movie.Data = MovieData{
		Adult:            smr.rawResult.Adult,
		Backdrop:         NewPtr(Image(smr.rawResult.BackdropPath)),
		OriginalLanguage: &smr.rawResult.OriginalLanguage,
		OriginalTitle:    &smr.rawResult.OriginalTitle,
		Overview:         &smr.rawResult.Overview,
		Popularity:       &smr.rawResult.Popularity,
		Poster:           NewPtr(Image(smr.rawResult.PosterPath)),
		ReleaseDate:      NewPtr(DateYYYYMMDD(smr.rawResult.ReleaseDate)),
		Title:            &smr.rawResult.Title,
		VoteAverage:      &smr.rawResult.VoteAverage,
		VoteCount:        &smr.rawResult.VoteCount,
	}

	for _, id := range smr.rawResult.GenreIds {
		genre := g.EnsureGenre(GenreId(id))
		movie.AddGenre(genre)
	}

	return movie
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
					rawResult:    smrMovie,
					Page:         page,
					TotalPages:   result.TotalPages,
					TotalResults: result.TotalResults,
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
