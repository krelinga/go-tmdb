package util

import (
	"context"
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

// Request builds the HTTP request using the provided context and parameters
func (rb *RequestBuilder) Request() *http.Request {
	// Set API key if available
	setIfNotZero(&rb.values, "api_key", APIKeyFromContext(rb.ctx))

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
	setAuthIfNotZero(request, APIReadAccessTokenFromContext(rb.ctx))
	// Set the context for the request
	return request.WithContext(rb.ctx)
}

// Do executes the HTTP request and returns the response
func (rb *RequestBuilder) Do() (*http.Response, error) {

	// Execute the request
	return HTTPClientFromContext(rb.ctx).Do(rb.Request())
}
