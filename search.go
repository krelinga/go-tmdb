package tmdb

import (
	"context"
	"net/http"

	tmdbsearch "github.com/krelinga/go-tmdb/tmdbsearch"
)

type SearchMoviesOptions = tmdbsearch.FindMoviesOptions
type SearchMoviesReply = tmdbsearch.FindMoviesReply

func SearchMovies(ctx context.Context, client *http.Client, query string, options SearchMoviesOptions) (*http.Response, error) {
	return tmdbsearch.FindMovies(ctx, client, query, options)
}
func ParseSearchMoviesReply(httpReply *http.Response) (*tmdbsearch.FindMoviesReply, error) {
	return tmdbsearch.ParseFindMoviesReply(httpReply)
}

type SearchSeriesOptions = tmdbsearch.FindSeriesOptions
type SearchSeriesReply = tmdbsearch.FindSeriesReply

func SearchSeries(ctx context.Context, client *http.Client, query string, options SearchSeriesOptions) (*http.Response, error) {
	return tmdbsearch.FindSeries(ctx, client, query, options)
}
func ParseSearchSeriesReply(httpReply *http.Response) (*tmdbsearch.FindSeriesReply, error) {
	return tmdbsearch.ParseFindSeriesReply(httpReply)
}
