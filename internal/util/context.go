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

// SetContext stores TMDB-specific values in the context
func SetContext(ctx context.Context, value Context) context.Context {
	return context.WithValue(ctx, contextKey{}, &value)
}

// GetContext retrieves TMDB-specific values from the context
func GetContext(ctx context.Context) (Context, bool) {
	value, ok := ctx.Value(contextKey{}).(*Context)
	if !ok || value == nil {
		return Context{}, false
	}
	return *value, true
}