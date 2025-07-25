package tmdb

import (
	"context"
	"net/http"

	tmdbsearch "github.com/krelinga/go-tmdb/search"
)

type SearchMovieOptions = tmdbsearch.MovieOptions
type SearchMovieReply = tmdbsearch.MovieReply

func SearchMovie(ctx context.Context, client *http.Client, query string, options SearchMovieOptions) (*http.Response, error) {
	return tmdbsearch.Movie(ctx, client, query, options)
}
func ParseSearchMovieReply(httpReply *http.Response) (*tmdbsearch.MovieReply, error) {
	return tmdbsearch.ParseMovieReply(httpReply)
}
