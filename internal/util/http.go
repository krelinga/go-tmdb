package util

import (
	"context"
	"net/http"
	"net/url"
)

func MakeRequest(ctx context.Context, client *http.Client, theUrl *url.URL, readAccessToken string) (*http.Response, error) {
	request := &http.Request{
		Method: http.MethodGet,
		URL:    theUrl,
	}
	SetAuthIfNotZero(request, readAccessToken)
	httpReply, err := client.Do(request.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	return httpReply, nil
}