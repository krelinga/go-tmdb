package tmdb

import (
	"context"
	"net/http"

	tmdbsearch "github.com/krelinga/go-tmdb/search"
)

type SearchMoviesOptions = tmdbsearch.FindMoviesOptions
type SearchMoviesReply = tmdbsearch.FindMoviesReply

func SearchMovies(ctx context.Context, client *http.Client, query string, options SearchMoviesOptions) (*http.Response, error) {
	return tmdbsearch.FindMovies(ctx, client, query, options)
}
func ParseSearchMoviesReply(httpReply *http.Response) (*tmdbsearch.FindMoviesReply, error) {
	return tmdbsearch.ParseFindMoviesReply(httpReply)
}
