package tmdb

import "fmt"

type Image interface {
	Raw() string
	Url(size string) (string, error)
}

type image struct {
	raw    string
	client *Client
}

func (i image) Raw() string {
	return i.raw
}

func (i image) Url(size string) (string, error) {
	if i.raw == "" {
		return "", fmt.Errorf("image URL is empty")
	}
	baseUrl, err := i.client.getSecureImageBaseUrl()
	if err != nil {
		return "", fmt.Errorf("getting secure image base URL: %w", err)
	}
	return baseUrl + size + i.raw, nil
}
