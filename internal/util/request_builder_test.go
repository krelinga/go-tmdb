package util

import (
	"context"
	"testing"
)

func TestRequestBuilder_URLConstruction(t *testing.T) {
	ctx := SetContext(context.Background(), Context{
		Key: "test-api-key",
	})
	
	rb := NewRequestBuilder(ctx).
		SetPath("/3/movie/123").
		SetValue("language", "en-US").
		AppendToResponse("credits", true).
		AppendToResponse("external_ids", false).
		AppendToResponse("reviews", true)

	// Test that we can access the URL construction logic by accessing the internal state
	// In a real scenario, we'd test this by intercepting the HTTP call
	
	// Verify path is set correctly
	if rb.path != "/3/movie/123" {
		t.Errorf("Expected path '/3/movie/123', got '%s'", rb.path)
	}
	
	// Verify language is set
	if rb.values.Get("language") != "en-US" {
		t.Errorf("Expected language 'en-US', got '%s'", rb.values.Get("language"))
	}
	
	// Verify appends array has correct values
	expectedAppends := []string{"credits", "reviews"}
	if len(rb.appends) != len(expectedAppends) {
		t.Errorf("Expected %d appends, got %d", len(expectedAppends), len(rb.appends))
	}
	for i, expected := range expectedAppends {
		if rb.appends[i] != expected {
			t.Errorf("Expected append[%d] to be '%s', got '%s'", i, expected, rb.appends[i])
		}
	}
}

func TestRequestBuilder_ChainableMethods(t *testing.T) {
	ctx := SetContext(context.Background(), Context{
		Key: "key",
		ReadAccessToken: "token",
	})

	// Test that all methods return *RequestBuilder for chaining
	rb := NewRequestBuilder(ctx)
	
	result := rb.SetPath("/test").
		SetValue("param", "value").
		AppendToResponse("test", true)
	
	if result != rb {
		t.Error("Methods should return the same RequestBuilder instance for chaining")
	}
}

func TestRequestBuilder_EmptyValues(t *testing.T) {
	ctx := SetContext(context.Background(), Context{
	})
	
	rb := NewRequestBuilder(ctx).
		SetPath("/test").
		SetValue("empty", "").  // Empty value should not be set
		SetValue("nonempty", "value")
	
	// Verify empty api_key is not set
	if rb.values.Has("api_key") {
		t.Error("Empty api_key should not be set in values")
	}
	
	// Verify empty custom value is not set
	if rb.values.Has("empty") {
		t.Error("Empty custom value should not be set in values")
	}
	
	// Verify non-empty value is set
	if rb.values.Get("nonempty") != "value" {
		t.Errorf("Expected 'value', got '%s'", rb.values.Get("nonempty"))
	}
}