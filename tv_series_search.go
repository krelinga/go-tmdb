package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"iter"
)

type TvSeriesSearchResult struct {
	TvSeriesSum
	GenreIds []GenreId `json:"genre_ids"`
}

func SearchTvSeries(client Client, query string, options ...SearchTvSeriesOption) iter.Seq2[*TvSeriesSearchResult, error] {
	return func(yield func(*TvSeriesSearchResult, error) bool) {
		o := searchTvSeriesOptions{}
		for _, opt := range options {
			opt.applyToSearchTvSeriesOptions(&o)
		}

		type rawReply struct {
			Page         int                     `json:"page"`
			Results      []*TvSeriesSearchResult `json:"results"`
			TotalPages   int                     `json:"total_pages"`
			TotalResults int                     `json:"total_results"`
		}
		ctx := context.Background()
		if o.useContext != nil {
			ctx = *o.useContext
		}
		for page := 1; ; page++ {
			params := GetParams{
				"query":         query,
				"include_adult": fmt.Sprintf("%t", o.wantAdult),
				"page":          fmt.Sprintf("%d", page),
			}
			if o.firstAirDateYear > 0 {
				params["first_air_date_year"] = fmt.Sprintf("%d", o.firstAirDateYear)
			}
			if o.year > 0 {
				params["year"] = fmt.Sprintf("%d", o.year)
			}
			data, err := checkCode(client.Get(ctx, "/search/tv", params))
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
			for _, tvSeries := range r.Results {
				if !yield(tvSeries, nil) {
					return
				}
			}
			if r.Page >= r.TotalPages {
				break
			}
		}
	}
}

type SearchTvSeriesOption interface {
	applyToSearchTvSeriesOptions(o *searchTvSeriesOptions)
}

type searchTvSeriesOptions struct {
	baseOptions
	wantAdult        bool
	firstAirDateYear int
	year int
}
