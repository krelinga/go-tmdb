package util

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestAPIKeyFromContext(t *testing.T) {
	tests := []struct {
		name        string
		ctx         context.Context
		expectedKey string
	}{
		{
			name:        "retrieve API key from context",
			ctx:         ContextWithAPIKey(context.Background(), "my-api-key"),
			expectedKey: "my-api-key",
		},
		{
			name:        "retrieve empty API key",
			ctx:         ContextWithAPIKey(context.Background(), ""),
			expectedKey: "",
		},
		{
			name:        "nil context returns empty string",
			ctx:         nil,
			expectedKey: "",
		},
		{
			name:        "context without API key returns empty string",
			ctx:         context.Background(),
			expectedKey: "",
		},
		{
			name:        "context with wrong value type returns empty string",
			ctx:         context.WithValue(context.Background(), apiKeyContextKey{}, 123),
			expectedKey: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retrievedKey := APIKeyFromContext(tt.ctx)
			if retrievedKey != tt.expectedKey {
				t.Errorf("Expected API key %q, got %q", tt.expectedKey, retrievedKey)
			}
		})
	}
}

func TestAPIReadAccessTokenFromContext(t *testing.T) {
	tests := []struct {
		name          string
		ctx           context.Context
		expectedToken string
	}{
		{
			name:          "retrieve token from context",
			ctx:           ContextWithAPIReadAccessToken(context.Background(), "my-token"),
			expectedToken: "my-token",
		},
		{
			name:          "retrieve empty token",
			ctx:           ContextWithAPIReadAccessToken(context.Background(), ""),
			expectedToken: "",
		},
		{
			name:          "nil context returns empty string",
			ctx:           nil,
			expectedToken: "",
		},
		{
			name:          "context without token returns empty string",
			ctx:           context.Background(),
			expectedToken: "",
		},
		{
			name:          "context with wrong value type returns empty string",
			ctx:           context.WithValue(context.Background(), apiReadAccessTokenContextKey{}, 456),
			expectedToken: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retrievedToken := APIReadAccessTokenFromContext(tt.ctx)
			if retrievedToken != tt.expectedToken {
				t.Errorf("Expected token %q, got %q", tt.expectedToken, retrievedToken)
			}
		})
	}
}

func TestHTTPClientFromContext(t *testing.T) {
	customClient := &http.Client{Timeout: 15 * time.Second}

	tests := []struct {
		name           string
		ctx            context.Context
		expectedClient *http.Client
	}{
		{
			name:           "retrieve custom client from context",
			ctx:            ContextWithHTTPClient(context.Background(), customClient),
			expectedClient: customClient,
		},
		{
			name:           "nil context returns default client",
			ctx:            nil,
			expectedClient: http.DefaultClient,
		},
		{
			name:           "context without client returns default client",
			ctx:            context.Background(),
			expectedClient: http.DefaultClient,
		},
		{
			name:           "context with wrong value type returns default client",
			ctx:            context.WithValue(context.Background(), httpClientContextKey{}, "not-a-client"),
			expectedClient: http.DefaultClient,
		},
		{
			name:           "context with nil client value returns nil",
			ctx:            context.WithValue(context.Background(), httpClientContextKey{}, (*http.Client)(nil)),
			expectedClient: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retrievedClient := HTTPClientFromContext(tt.ctx)
			if retrievedClient != tt.expectedClient {
				t.Errorf("Expected client %v, got %v", tt.expectedClient, retrievedClient)
			}
		})
	}
}

// Integration tests
func TestContextIntegration(t *testing.T) {
	t.Run("multiple values in same context", func(t *testing.T) {
		apiKey := "test-api-key"
		token := "test-token"
		client := &http.Client{Timeout: 20 * time.Second}

		// Build context with all values
		ctx := context.Background()
		ctx = ContextWithAPIKey(ctx, apiKey)
		ctx = ContextWithAPIReadAccessToken(ctx, token)
		ctx = ContextWithHTTPClient(ctx, client)

		// Verify all values can be retrieved
		if retrievedKey := APIKeyFromContext(ctx); retrievedKey != apiKey {
			t.Errorf("Expected API key %q, got %q", apiKey, retrievedKey)
		}

		if retrievedToken := APIReadAccessTokenFromContext(ctx); retrievedToken != token {
			t.Errorf("Expected token %q, got %q", token, retrievedToken)
		}

		if retrievedClient := HTTPClientFromContext(ctx); retrievedClient != client {
			t.Errorf("Expected client %v, got %v", client, retrievedClient)
		}
	})

	t.Run("overwriting values", func(t *testing.T) {
		ctx := context.Background()

		// Add initial values
		ctx = ContextWithAPIKey(ctx, "first-key")
		ctx = ContextWithAPIReadAccessToken(ctx, "first-token")

		// Overwrite with new values
		ctx = ContextWithAPIKey(ctx, "second-key")
		ctx = ContextWithAPIReadAccessToken(ctx, "second-token")

		// Verify we get the latest values
		if retrievedKey := APIKeyFromContext(ctx); retrievedKey != "second-key" {
			t.Errorf("Expected API key %q, got %q", "second-key", retrievedKey)
		}

		if retrievedToken := APIReadAccessTokenFromContext(ctx); retrievedToken != "second-token" {
			t.Errorf("Expected token %q, got %q", "second-token", retrievedToken)
		}
	})
}
