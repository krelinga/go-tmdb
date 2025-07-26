package util

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// RequestBuilder combines URL building and HTTP request functionality
type RequestBuilder struct {
	ctx             context.Context
	client          *http.Client
	path            string
	appends         []string
	values          url.Values
	readAccessToken string
}

// NewRequestBuilder creates a new RequestBuilder with the provided context and HTTP client
func NewRequestBuilder(ctx context.Context, client *http.Client) *RequestBuilder {
	return &RequestBuilder{
		ctx:    ctx,
		client: client,
		values: make(url.Values),
	}
}

// SetPath sets the API path for the request
func (rb *RequestBuilder) SetPath(path string) *RequestBuilder {
	rb.path = path
	return rb
}

// SetApiKey sets the API key parameter
func (rb *RequestBuilder) SetApiKey(apiKey string) *RequestBuilder {
	SetIfNotZero(&rb.values, "api_key", apiKey)
	return rb
}

// SetReadAccessToken sets the read access token for authorization
func (rb *RequestBuilder) SetReadAccessToken(token string) *RequestBuilder {
	rb.readAccessToken = token
	return rb
}

// AppendToResponse adds a value to the append_to_response parameter if the condition is true
func (rb *RequestBuilder) AppendToResponse(value string, do bool) *RequestBuilder {
	if do {
		rb.appends = append(rb.appends, value)
	}
	return rb
}

// SetValue sets a custom query parameter
func (rb *RequestBuilder) SetValue(key, value string) *RequestBuilder {
	SetIfNotZero(&rb.values, key, value)
	return rb
}

// Do executes the HTTP request and returns the response
func (rb *RequestBuilder) Do() (*http.Response, error) {
	// Build the URL
	SetIfNotZero(&rb.values, "append_to_response", strings.Join(rb.appends, ","))
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
	SetAuthIfNotZero(request, rb.readAccessToken)

	// Execute the request
	return rb.client.Do(request.WithContext(rb.ctx))
}