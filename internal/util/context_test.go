package util

import (
	"context"
	"net/http"
	"testing"
)

func TestSetContext_Success(t *testing.T) {
	ctx := context.Background()
	client := &http.Client{}

	value := Context{
		Key:             "test-api-key",
		ReadAccessToken: "test-read-token",
		Client:          client,
	}

	newCtx := SetContext(ctx, value)

	if newCtx == ctx {
		t.Error("SetContext should return a new context")
	}
}

func TestGetContext_Success(t *testing.T) {
	ctx := context.Background()
	client := &http.Client{}

	original := Context{
		Key:             "test-api-key",
		ReadAccessToken: "test-read-token",
		Client:          client,
	}

	ctxWithValue := SetContext(ctx, original)
	retrieved, ok := GetContext(ctxWithValue)

	if !ok {
		t.Error("Expected GetContext to return true")
	}

	if retrieved.Key != original.Key {
		t.Errorf("Expected Key '%s', got '%s'", original.Key, retrieved.Key)
	}

	if retrieved.ReadAccessToken != original.ReadAccessToken {
		t.Errorf("Expected ReadAccessToken '%s', got '%s'", original.ReadAccessToken, retrieved.ReadAccessToken)
	}

	if retrieved.Client != original.Client {
		t.Error("Expected Client to be the same instance")
	}
}

func TestGetContext_EmptyContext(t *testing.T) {
	ctx := context.Background()

	_, ok := GetContext(ctx)

	if ok {
		t.Error("Expected GetContext to return false for empty context")
	}
}

func TestGetContext_NilContext(t *testing.T) {
	var ctx context.Context

	_, ok := GetContext(ctx)

	if ok {
		t.Error("Expected GetContext to return false for nil context")
	}
}

func TestSetGetContext_RoundTrip(t *testing.T) {
	ctx := context.Background()
	client := &http.Client{}

	original := Context{
		Key:             "test-api-key",
		ReadAccessToken: "test-read-token",
		Client:          client,
	}

	// Set and get multiple times to ensure consistency
	ctx1 := SetContext(ctx, original)
	_, ok1 := GetContext(ctx1)

	if !ok1 {
		t.Error("First retrieval should succeed")
	}

	// Set again with different values
	updated := Context{
		Key:             "updated-api-key",
		ReadAccessToken: "updated-read-token",
		Client:          client,
	}

	ctx2 := SetContext(ctx1, updated)
	retrieved2, ok2 := GetContext(ctx2)

	if !ok2 {
		t.Error("Second retrieval should succeed")
	}

	if retrieved2.Key != updated.Key {
		t.Errorf("Expected updated Key '%s', got '%s'", updated.Key, retrieved2.Key)
	}

	// Original context should still have original values
	retrievedOriginal, okOriginal := GetContext(ctx1)
	if !okOriginal {
		t.Error("Original context retrieval should succeed")
	}

	if retrievedOriginal.Key != original.Key {
		t.Errorf("Expected original Key '%s', got '%s'", original.Key, retrievedOriginal.Key)
	}
}

func TestContext_EmptyValues(t *testing.T) {
	ctx := context.Background()

	value := Context{
		Key:             "",
		ReadAccessToken: "",
		Client:          nil,
	}

	ctxWithValue := SetContext(ctx, value)
	retrieved, ok := GetContext(ctxWithValue)

	if !ok {
		t.Error("Expected GetContext to return true even with empty values")
	}

	if retrieved.Key != "" {
		t.Errorf("Expected empty Key, got '%s'", retrieved.Key)
	}

	if retrieved.ReadAccessToken != "" {
		t.Errorf("Expected empty ReadAccessToken, got '%s'", retrieved.ReadAccessToken)
	}

	if retrieved.Client != nil {
		t.Error("Expected nil Client")
	}
}

func TestContext_MultipleContextKeys(t *testing.T) {
	ctx := context.Background()
	client := &http.Client{}

	// Add other values to context to ensure no collisions
	ctx = context.WithValue(ctx, "other-key", "other-value")
	ctx = context.WithValue(ctx, struct{}{}, "struct-key-value")

	value := Context{
		Key:             "test-api-key",
		ReadAccessToken: "test-read-token",
		Client:          client,
	}

	ctxWithValue := SetContext(ctx, value)
	retrieved, ok := GetContext(ctxWithValue)

	if !ok {
		t.Error("Expected GetContext to return true")
	}

	if retrieved.Key != value.Key {
		t.Errorf("Expected Key '%s', got '%s'", value.Key, retrieved.Key)
	}

	// Ensure other context values are still there
	if ctxWithValue.Value("other-key") != "other-value" {
		t.Error("Other context values should not be affected")
	}

	if ctxWithValue.Value(struct{}{}) != "struct-key-value" {
		t.Error("Struct key context values should not be affected")
	}
}
