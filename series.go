package tmdb

import (
	"context"
	"net/http"

	"github.com/krelinga/go-tmdb/series"
)

type GetSeriesDetailsOptions = tmdbseries.GetDetailsOptions
type GetSeriesDetailsReply = tmdbseries.GetDetailsReply

func GetSeriesDetails(ctx context.Context, client *http.Client, id int32, options GetSeriesDetailsOptions) (*GetSeriesDetailsReply, error) {
	return tmdbseries.GetDetails(ctx, client, id, options)
}