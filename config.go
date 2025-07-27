package tmdb

import (
	"context"
	"net/http"

	"github.com/krelinga/go-tmdb/tmdbconfig"
)

type GetConfigDetailsOptions = tmdbconfig.GetDetailsOptions
type GetConfigDetailsReply = tmdbconfig.GetDetailsReply

func GetConfigDetails(ctx context.Context, options GetConfigDetailsOptions) (*http.Response, error) {
	return tmdbconfig.GetDetails(ctx, options)
}
func ParseGetConfigDetailsReply(httpReply *http.Response) (*GetConfigDetailsReply, error) {
	return tmdbconfig.ParseGetDetailsReply(httpReply)
}
