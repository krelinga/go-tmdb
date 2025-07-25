package tmdb

import (
	"context"
	"net/http"

	"github.com/krelinga/go-tmdb/search"
)

type SearchMovieOptions = tmdbsearch.MovieOptions
type SearchMovieReply = tmdbsearch.MovieReply

func SearchMovie(ctx context.Context, client *http.Client, query string, options SearchMovieOptions) (*SearchMovieReply, error) {
	return tmdbsearch.Movie(ctx, client, query, options)
}
