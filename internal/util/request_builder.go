package util

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// RequestBuilder combines URL building and HTTP request functionality
type RequestBuilder struct {
	ctx     context.Context
	path    string
	appends []string
	values  url.Values
}

// NewRequestBuilder creates a new RequestBuilder with the provided context
func NewRequestBuilder(ctx context.Context) *RequestBuilder {
	return &RequestBuilder{
		ctx:    ctx,
		values: make(url.Values),
	}
}

// SetPath sets the API path for the request
func (rb *RequestBuilder) SetPath(path string) *RequestBuilder {
	rb.path = path
	return rb
}

// AppendToResponse adds a value to the append_to_response parameter if the condition is true
func (rb *RequestBuilder) AppendToResponse(value string, do bool) *RequestBuilder {
	if do {
		rb.appends = append(rb.appends, value)
	}
	return rb
}

// SetValueString sets a custom query parameter
func (rb *RequestBuilder) SetValueString(key, value string) *RequestBuilder {
	setIfNotZero(&rb.values, key, value)
	return rb
}

// SetValueInt32 sets a custom query parameter with an int32 value
func (rb *RequestBuilder) SetValueInt32(key string, value int32) *RequestBuilder {
	setIfNotZero(&rb.values, key, value)
	return rb
}

// Do executes the HTTP request and returns the response
func (rb *RequestBuilder) Do() (*http.Response, error) {
	// Get configuration from context
	tmdbCtx, ok := GetContext(rb.ctx)
	if !ok {
		return nil, fmt.Errorf("TMDB context not found")
	}

	// Set API key if available
	setIfNotZero(&rb.values, "api_key", tmdbCtx.Key)

	// Build the URL
	setIfNotZero(&rb.values, "append_to_response", strings.Join(rb.appends, ","))
	requestURL := &url.URL{
		Scheme:   "https",
		Host:     "api.themoviedb.org",
		Path:     rb.path,
		RawQuery: rb.values.Encode(),
	}

	// Create the HTTP request
	request := &http.Request{
		Method: http.MethodGet,
		URL:    requestURL,
	}

	// Set authorization if provided
	setAuthIfNotZero(request, tmdbCtx.ReadAccessToken)

	// Use client from context, fallback to default client
	client := tmdbCtx.Client
	if client == nil {
		client = http.DefaultClient
	}

	// Execute the request
	return client.Do(request.WithContext(rb.ctx))
}
