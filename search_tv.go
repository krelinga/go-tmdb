package tmdb

import (
	"context"
	"fmt"
	"iter"
	"net/url"

	"github.com/krelinga/go-tmdb/internal/raw"
)

type SearchTvResult struct {
	Tv           *Show
	Page         int
	TotalResults int
	TotalPages   int
}

func SearchTv(ctx context.Context, client *Client, query string, options ...Option) iter.Seq2[*SearchTvResult, error] {
	return func(yield func(*SearchTvResult, error) bool) {
		callOpts := client.globalOpts
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

			var result raw.SearchTv
			if err := get(ctx, client, "search/tv", params, callOpts, &result); err != nil {
				yield(nil, fmt.Errorf("searching for tv %q: %w", query, err))
				return
			}

			for _, resultTv := range result.Results {
				genres := make([]Genre, len(resultTv.GenreIds))
				for i, id := range resultTv.GenreIds {
					genres[i] = Genre{
						Key: GenreId(id),
					}
				}
				out := &SearchTvResult{
					Page:         page,
					TotalResults: result.TotalResults,
					TotalPages:   result.TotalPages,
					Tv: &Show{
						Key: ShowId(resultTv.Id),
						Data: ShowData{
							Adult:            resultTv.Adult,
							Backdrop:         NewPtr(Image(resultTv.BackdropPath)),
							Genres:           genres,
							OriginCountry:    resultTv.OriginCountry,
							OriginalLanguage: &resultTv.OriginalLanguage,
							OriginalName:     &resultTv.OriginalName,
							Overview:         &resultTv.Overview,
							Popularity:       &resultTv.Popularity,
							Poster:           NewPtr(Image(resultTv.PosterPath)),
							FirstAirDate:     NewPtr(DateYYYYMMDD(resultTv.FirstAirDate)),
							Name:             &resultTv.Name,
							VoteAverage:      &resultTv.VoteAverage,
							VoteCount:        &resultTv.VoteCount,
						},
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
