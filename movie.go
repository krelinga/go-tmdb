package tmdb

import (
	"context"
	"net/http"

	tmdbmovie "github.com/krelinga/go-tmdb/movie"
)

type GetMovieDetailsOptions = tmdbmovie.GetDetailsOptions
type GetMovieDetailsReply = tmdbmovie.GetDetailsReply

func GetMovieDetails(ctx context.Context, client *http.Client, id int32, options GetMovieDetailsOptions) (*GetMovieDetailsReply, error) {
	return tmdbmovie.GetDetails(ctx, client, id, options)
}

type GetMovieCreditsOptions = tmdbmovie.GetCreditsOptions
type GetMovieCreditsReply = tmdbmovie.GetCreditsReply

func GetMovieCredits(ctx context.Context, client *http.Client, id int32, options GetMovieCreditsOptions) (*http.Response, error) {
	return tmdbmovie.GetCredits(ctx, client, id, options)
}
func ParseGetMovieCreditsReply(httpReply *http.Response) (*tmdbmovie.GetCreditsReply, error) {
	return tmdbmovie.ParseGetCreditsReply(httpReply)
}

type GetMovieExternalIDsOptions = tmdbmovie.GetExternalIDsOptions
type GetMovieExternalIDsReply = tmdbmovie.GetExternalIDsReply

func GetMovieExternalIDs(ctx context.Context, client *http.Client, id int32, options GetMovieExternalIDsOptions) (*http.Response, error) {
	return tmdbmovie.GetExternalIDs(ctx, client, id, options)
}
func ParseGetMovieExternalIDsReply(httpReply *http.Response) (*tmdbmovie.GetExternalIDsReply, error) {
	return tmdbmovie.ParseGetExternalIDsReply(httpReply)
}

type GetMovieReleaseDatesOptions = tmdbmovie.GetReleaseDatesOptions
type GetMovieReleaseDatesReply = tmdbmovie.GetReleaseDatesReply

func GetMovieReleaseDates(ctx context.Context, client *http.Client, id int32, options GetMovieReleaseDatesOptions) (*http.Response, error) {
	return tmdbmovie.GetReleaseDates(ctx, client, id, options)
}
func ParseGetMovieReleaseDatesReply(httpReply *http.Response) (*tmdbmovie.GetReleaseDatesReply, error) {
	return tmdbmovie.ParseGetReleaseDatesReply(httpReply)
}
