package tmdb

import (
	"context"
	"net/http"

	"github.com/krelinga/go-tmdb/tmdbmovie"
)

type GetMovieDetailsOptions = tmdbmovie.GetDetailsOptions
type GetMovieDetailsReply = tmdbmovie.GetDetailsReply

func GetMovieDetails(ctx context.Context, id int32, options GetMovieDetailsOptions) (*http.Response, error) {
	return tmdbmovie.GetDetails(ctx, id, options)
}
func ParseGetMovieDetailsReply(httpReply *http.Response) (*GetMovieDetailsReply, error) {
	return tmdbmovie.ParseGetDetailsReply(httpReply)
}

type GetMovieCreditsOptions = tmdbmovie.GetCreditsOptions
type GetMovieCreditsReply = tmdbmovie.GetCreditsReply

func GetMovieCredits(ctx context.Context, id int32, options GetMovieCreditsOptions) (*http.Response, error) {
	return tmdbmovie.GetCredits(ctx, id, options)
}
func ParseGetMovieCreditsReply(httpReply *http.Response) (*GetMovieCreditsReply, error) {
	return tmdbmovie.ParseGetCreditsReply(httpReply)
}

type GetMovieExternalIDsOptions = tmdbmovie.GetExternalIDsOptions
type GetMovieExternalIDsReply = tmdbmovie.GetExternalIDsReply

func GetMovieExternalIDs(ctx context.Context, id int32, options GetMovieExternalIDsOptions) (*http.Response, error) {
	return tmdbmovie.GetExternalIDs(ctx, id, options)
}
func ParseGetMovieExternalIDsReply(httpReply *http.Response) (*GetMovieExternalIDsReply, error) {
	return tmdbmovie.ParseGetExternalIDsReply(httpReply)
}

type GetMovieReleaseDatesOptions = tmdbmovie.GetReleaseDatesOptions
type GetMovieReleaseDatesReply = tmdbmovie.GetReleaseDatesReply

func GetMovieReleaseDates(ctx context.Context, id int32, options GetMovieReleaseDatesOptions) (*http.Response, error) {
	return tmdbmovie.GetReleaseDates(ctx, id, options)
}
func ParseGetMovieReleaseDatesReply(httpReply *http.Response) (*GetMovieReleaseDatesReply, error) {
	return tmdbmovie.ParseGetReleaseDatesReply(httpReply)
}
