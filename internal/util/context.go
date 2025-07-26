package util

import (
	"context"
	"net/http"
)

// tmdbContextKey is an unexported type for context keys to avoid collisions
type tmdbContextKey struct{}

// TMDBContextValue holds the TMDB-specific values stored in context
type TMDBContextValue struct {
	Key             string
	ReadAccessToken string
	Client          *http.Client
}

// WithTMDBContext stores TMDB-specific values in the context
func WithTMDBContext(ctx context.Context, key, readAccessToken string, client *http.Client) context.Context {
	value := &TMDBContextValue{
		Key:             key,
		ReadAccessToken: readAccessToken,
		Client:          client,
	}
	return context.WithValue(ctx, tmdbContextKey{}, value)
}

// FromTMDBContext retrieves TMDB-specific values from the context
func FromTMDBContext(ctx context.Context) (TMDBContextValue, bool) {
	value, ok := ctx.Value(tmdbContextKey{}).(*TMDBContextValue)
	if !ok || value == nil {
		return TMDBContextValue{}, false
	}
	return *value, true
}