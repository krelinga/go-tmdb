package util

import (
	"net/url"
	"strings"
)

type URLBuilder struct {
	appends []string
	values  url.Values
	path    string
}

func NewURLBuilder(path string) *URLBuilder {
	b := &URLBuilder{
		path:   path,
		values: make(url.Values),
	}
	return b
}

func (b *URLBuilder) SetApiKey(apiKey string) *URLBuilder {
	SetIfNotZero(&b.values, "api_key", apiKey)
	return b
}

func (b *URLBuilder) AppendToResponse(value string, do bool) *URLBuilder {
	if do {
		b.appends = append(b.appends, value)
	}
	return b
}

func (b *URLBuilder) SetValue(key, value string) *URLBuilder {
	SetIfNotZero(&b.values, key, value)
	return b
}

func (b *URLBuilder) URL() *url.URL {
	SetIfNotZero(&b.values, "append_to_response", strings.Join(b.appends, ","))
	return &url.URL{
		Scheme:   "https",
		Host:     "api.themoviedb.org",
		Path:     b.path,
		RawQuery: b.values.Encode(),
	}
}
