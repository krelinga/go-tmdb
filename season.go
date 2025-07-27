package tmdb

import (
	"context"
	"net/http"

	"github.com/krelinga/go-tmdb/tmdbseason"
)

type GetSeasonDetailsOptions = tmdbseason.GetDetailsOptions
type GetSeasonDetailsReply = tmdbseason.GetDetailsReply

func GetSeasonDetails(ctx context.Context, seriesID int32, seasonNumber int32, options GetSeasonDetailsOptions) (*http.Response, error) {
	return tmdbseason.GetDetails(ctx, seriesID, seasonNumber, options)
}
func ParseGetSeasonDetailsReply(httpReply *http.Response) (*GetSeasonDetailsReply, error) {
	return tmdbseason.ParseGetDetailsReply(httpReply)
}