package tmdb

import (
	"context"
	"net/http"

	"github.com/krelinga/go-tmdb/tmdbseries"
)

type GetSeasonDetailsOptions = tmdbseries.GetDetailsOptions
type GetSeasonDetailsReply = tmdbseries.GetDetailsReply

func GetSeasonDetails(ctx context.Context, client *http.Client, id int32, options GetSeasonDetailsOptions) (*http.Response, error) {
	return tmdbseries.GetDetails(ctx, client, id, options)
}
func ParseGetSeasonDetailsReply(httpReply *http.Response) (*GetSeasonDetailsReply, error) {
	return tmdbseries.ParseGetDetailsReply(httpReply)
}