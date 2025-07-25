package tmdb

import (
	"context"
	"net/http"

	"github.com/krelinga/go-tmdb/movies"
)

type GetMovieDetailsOptions = tmdbmovie.GetDetailsOptions
type GetMovieDetailsReply = tmdbmovie.GetDetailsReply

func GetMovieDetails(ctx context.Context, client *http.Client, id int32, options GetMovieDetailsOptions) (*GetMovieDetailsReply, error) {
	return tmdbmovie.GetDetails(ctx, client, id, options)
}

type GetMovieCreditsOptions = tmdbmovie.GetCreditsOptions
type GetMovieCreditsReply = tmdbmovie.GetCreditsReply

func GetMovieCredits(ctx context.Context, client *http.Client, id int32, options GetMovieCreditsOptions) (*GetMovieCreditsReply, error) {
	return tmdbmovie.GetCredits(ctx, client, id, options)
}

type GetMovieExternalIDsOptions = tmdbmovie.GetExternalIDsOptions
type GetMovieExternalIDsReply = tmdbmovie.GetExternalIDsReply

func GetMovieExternalIDs(ctx context.Context, client *http.Client, id int32, options GetMovieExternalIDsOptions) (*GetMovieExternalIDsReply, error) {
	return tmdbmovie.GetExternalIDs(ctx, client, id, options)
}

type GetMovieReleaseDatesOptions = tmdbmovie.GetReleaseDatesOptions
type GetMovieReleaseDatesReply = tmdbmovie.GetReleaseDatesReply

func GetMovieReleaseDates(ctx context.Context, client *http.Client, id int32, options GetMovieReleaseDatesOptions) (*GetMovieReleaseDatesReply, error) {
	return tmdbmovie.GetReleaseDates(ctx, client, id, options)
}
