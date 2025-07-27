package tmdb

import (
	"context"
	"net/http"

	"github.com/krelinga/go-tmdb/internal/util"
)

func ContextWithAPIKey(ctx context.Context, key string) context.Context {
	return util.ContextWithAPIKey(ctx, key)
}

func ContextWithAPIReadAccessToken(ctx context.Context, token string) context.Context {
	return util.ContextWithAPIReadAccessToken(ctx, token)
}

func ContextWithHTTPClient(ctx context.Context, client *http.Client) context.Context {
	return util.ContextWithHTTPClient(ctx, client)
}
