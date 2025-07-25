package tmdb

import (
	"context"
	"net/http"

	"github.com/krelinga/go-tmdb/tmdbseries"
)

type GetSeriesDetailsOptions = tmdbseries.GetDetailsOptions
type GetSeriesDetailsReply = tmdbseries.GetDetailsReply

func GetSeriesDetails(ctx context.Context, client *http.Client, id int32, options GetSeriesDetailsOptions) (*http.Response, error) {
	return tmdbseries.GetDetails(ctx, client, id, options)
}
func ParseGetSeriesDetailsReply(httpReply *http.Response) (*GetSeriesDetailsReply, error) {
	return tmdbseries.ParseGetDetailsReply(httpReply)
}
