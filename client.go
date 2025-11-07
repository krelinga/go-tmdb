package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client interface {
	GetObject(ctx context.Context, path string, options ...RequestOption) (Object, error)
	GetArray(ctx context.Context, path string, options ...RequestOption) (Array, error)
}

type ClientOptions struct {
	APIKey             string
	APIReadAccessToken string
	HttpClient         *http.Client
}

func (co ClientOptions) NewClient() Client {
	if co.HttpClient == nil {
		co.HttpClient = http.DefaultClient
	}
	return &clientImpl{
		options: co,
	}
}

type clientImpl struct {
	options ClientOptions
}

func (c *clientImpl) getRaw(ctx context.Context, path string, options ...RequestOption) (io.ReadCloser, error) {
	if c.options.APIKey != "" {
		options = append(options, WithQueryParam("api_key", c.options.APIKey))
	}
	if c.options.APIReadAccessToken != "" {
		options = append(options, WithRequestHeader("Authorization", "Bearer "+c.options.APIReadAccessToken))
	}
	urlValues := url.Values{}
	for _, opt := range options {
		if opt.ChangeValues != nil {
			opt.ChangeValues(&urlValues)
		}
	}
	reqUrl := &url.URL{
		Scheme:   "https",
		Host:     "api.themoviedb.org",
		Path:     path,
		RawQuery: urlValues.Encode(),
	}
	reqHeader := http.Header{}
	for _, opt := range options {
		if opt.ChangeHeader != nil {
			opt.ChangeHeader(&reqHeader)
		}
	}
	req := &http.Request{
		Method: http.MethodGet,
		URL:    reqUrl,
		Header: reqHeader,
	}
	for _, opt := range options {
		if opt.ChangeRequest != nil {
			opt.ChangeRequest(req)
		}
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	response, err := c.options.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	for _, opt := range options {
		if opt.ChangeResponse != nil {
			opt.ChangeResponse(response)
		}
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
	contentType := response.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		return nil, fmt.Errorf("unexpected content type: %s", contentType)
	}
	return response.Body, nil
}

func (c *clientImpl) GetObject(ctx context.Context, path string, options ...RequestOption) (Object, error) {
	body, err := c.getRaw(ctx, path, options...)
	if err != nil {
		return nil, err
	}
	defer body.Close()
	o := Object{}
	if err := json.NewDecoder(body).Decode(&o); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return o, nil
}

func (c *clientImpl) GetArray(ctx context.Context, path string, options ...RequestOption) (Array, error) {
	body, err := c.getRaw(ctx, path, options...)
	if err != nil {
		return nil, err
	}
	defer body.Close()
	var arr Array
	if err := json.NewDecoder(body).Decode(&arr); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return arr, nil
}
