package tmdb

import (
	"context"
	"net/http"

	tmdbsearch "github.com/krelinga/go-tmdb/tmdbsearch"
)

type SearchMoviesOptions = tmdbsearch.FindMoviesOptions
type SearchMoviesReply = tmdbsearch.FindMoviesReply

func SearchMovies(ctx context.Context, query string, options SearchMoviesOptions) (*http.Response, error) {
	return tmdbsearch.FindMovies(ctx, query, options)
}
func ParseSearchMoviesReply(httpReply *http.Response) (*SearchMoviesReply, error) {
	return tmdbsearch.ParseFindMoviesReply(httpReply)
}

type SearchSeriesOptions = tmdbsearch.FindSeriesOptions
type SearchSeriesReply = tmdbsearch.FindSeriesReply

func SearchSeries(ctx context.Context, query string, options SearchSeriesOptions) (*http.Response, error) {
	return tmdbsearch.FindSeries(ctx, query, options)
}
func ParseSearchSeriesReply(httpReply *http.Response) (*SearchSeriesReply, error) {
	return tmdbsearch.ParseFindSeriesReply(httpReply)
}
