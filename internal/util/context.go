package util

import (
	"context"
	"net/http"
)

type apiKeyContextKey struct{}

func ContextWithAPIKey(ctx context.Context, apiKey string) context.Context {
	return context.WithValue(ctx, apiKeyContextKey{}, apiKey)
}

func APIKeyFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	apiKey, ok := ctx.Value(apiKeyContextKey{}).(string)
	if !ok {
		return ""
	}
	return apiKey
}

type apiReadAccessTokenContextKey struct{}

func ContextWithAPIReadAccessToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, apiReadAccessTokenContextKey{}, token)
}

func APIReadAccessTokenFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	token, ok := ctx.Value(apiReadAccessTokenContextKey{}).(string)
	if !ok {
		return ""
	}
	return token
}

type httpClientContextKey struct{}

func ContextWithHTTPClient(ctx context.Context, client *http.Client) context.Context {
	return context.WithValue(ctx, httpClientContextKey{}, client)
}

func HTTPClientFromContext(ctx context.Context) *http.Client {
	if ctx == nil {
		return http.DefaultClient
	}
	client, ok := ctx.Value(httpClientContextKey{}).(*http.Client)
	if !ok {
		return http.DefaultClient
	}
	return client
}
