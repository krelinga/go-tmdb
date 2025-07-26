package util

import (
	"context"
	"net/http"
)

// contextKey is an unexported type for context keys to avoid collisions
type contextKey struct{}

// Context holds the TMDB-specific values stored in context
type Context struct {
	Key             string
	ReadAccessToken string
	Client          *http.Client
}

// WithContext stores TMDB-specific values in the context
func WithContext(ctx context.Context, key, readAccessToken string, client *http.Client) context.Context {
	value := &Context{
		Key:             key,
		ReadAccessToken: readAccessToken,
		Client:          client,
	}
	return context.WithValue(ctx, contextKey{}, value)
}

// GetContext retrieves TMDB-specific values from the context
func GetContext(ctx context.Context) (Context, bool) {
	value, ok := ctx.Value(contextKey{}).(*Context)
	if !ok || value == nil {
		return Context{}, false
	}
	return *value, true
}