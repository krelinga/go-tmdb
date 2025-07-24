package tmdb

import (
	"context"
	"net/http"

	"github.com/krelinga/go-tmdb/namespaces/movies"
)

type GetMovieDetailsOptions = movies.GetDetailsOptions
type GetMovieDetailsReply = movies.GetDetailsReply

func GetMovieDetails(ctx context.Context, client *http.Client, id int32, options GetMovieDetailsOptions) (*GetMovieDetailsReply, error) {
	return movies.GetDetails(ctx, client, id, options)
}

type GetMovieCreditsOptions = movies.GetCreditsOptions
type GetMovieCreditsReply = movies.GetCreditsReply

func GetMovieCredits(ctx context.Context, client *http.Client, id int32, options GetMovieCreditsOptions) (*GetMovieCreditsReply, error) {
	return movies.GetCredits(ctx, client, id, options)
}

type GetMovieMultiOptions = movies.GetMultiOptions
type GetMovieMultiReply = movies.GetMultiReply

func GetMovieMulti(ctx context.Context, client *http.Client, id int32, options GetMovieMultiOptions) (*GetMovieMultiReply, error) {
	return movies.GetMulti(ctx, client, id, options)
}