package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"iter"
)

type MovieSearchResult struct {
	MovieShort
	GenereIds []GenereId `json:"genre_ids"`
}

func SearchMovies(client Client, query string, options ...SearchMoviesOption) iter.Seq2[*MovieSearchResult, error] {
	return func(yield func(*MovieSearchResult, error) bool) {
		o := searchMoviesOptions{}
		for _, opt := range options {
			opt.applyToSearchMoviesOptions(&o)
		}

		type rawReply struct {
			Page         int                  `json:"page"`
			Results      []*MovieSearchResult `json:"results"`
			TotalPages   int                  `json:"total_pages"`
			TotalResults int                  `json:"total_results"`
		}
		ctx := context.Background()
		if o.useContext != nil {
			ctx = *o.useContext
		}
		for page := 1; ; page++ {
			params := GetParams{
				"query":         query,
				"page":          fmt.Sprintf("%d", page),
				"include_adult": fmt.Sprintf("%t", o.wantAdult),
			}
			data, err := checkCode(client.Get(ctx, "/search/movie", params))
			if err != nil {
				yield(nil, err)
				return
			}
			if o.rawReply != nil {
				*o.rawReply = data
			}
			r := &rawReply{}
			if err := json.Unmarshal(data, r); err != nil {
				yield(nil, fmt.Errorf("unmarshalling search result: %w", err))
				return
			}
			for _, movie := range r.Results {
				if !yield(movie, nil) {
					return
				}
			}
			if r.Page >= r.TotalPages {
				break
			}
		}
	}
}

type searchMoviesOptions struct {
	baseOptions
	wantAdult bool
}

type SearchMoviesOption interface {
	applyToSearchMoviesOptions(o *searchMoviesOptions)
}
