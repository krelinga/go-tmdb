package tmdb

import (
	"context"
	"net/http"

	"github.com/krelinga/go-tmdb/search"
)

type SearchMovieOptions = search.MovieOptions
type SearchMovieReply = search.MovieReply

func SearchMovie(ctx context.Context, client *http.Client, query string, options SearchMovieOptions) (*SearchMovieReply, error) {
	return search.Movie(ctx, client, query, options)
}
