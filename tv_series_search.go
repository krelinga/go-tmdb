package tmdb

import "iter"

type TvSeriesSearchResult struct {
	TvSeriesSum
	GenreIds []GenreId `json:"genre_ids"`
}

func SearchTvSeries(client Client, query string, options ...SearchTvSeriesOption) iter.Seq2[*TvSeriesSearchResult, error] {
	return nil
}

type SearchTvSeriesOption interface {
	applyToSearchTvSeriesOptions(o *searchTvSeriesOptions)
}

type searchTvSeriesOptions struct {
	baseOptions
}