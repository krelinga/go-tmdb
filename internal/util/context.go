package util

import (
	"context"
	"net/http"
)

// tmdbContextKey is an unexported type for context keys to avoid collisions
type tmdbContextKey struct{}

// tmdbContextValue holds the TMDB-specific values stored in context
type tmdbContextValue struct {
	Key             string
	ReadAccessToken string
	Client          *http.Client
}

// WithTMDBContext stores TMDB-specific values in the context
func WithTMDBContext(ctx context.Context, key, readAccessToken string, client *http.Client) context.Context {
	value := &tmdbContextValue{
		Key:             key,
		ReadAccessToken: readAccessToken,
		Client:          client,
	}
	return context.WithValue(ctx, tmdbContextKey{}, value)
}

// FromTMDBContext retrieves TMDB-specific values from the context
func FromTMDBContext(ctx context.Context) (key, readAccessToken string, client *http.Client, ok bool) {
	value, ok := ctx.Value(tmdbContextKey{}).(*tmdbContextValue)
	if !ok || value == nil {
		return "", "", nil, false
	}
	return value.Key, value.ReadAccessToken, value.Client, true
}