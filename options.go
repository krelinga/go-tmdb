package tmdb

import (
	"fmt"
	"net/url"
)

type allOptions struct {
	Language *string
	IncludeAdult *bool
	StartPage *int
	LimitPage *int
	MovieDataColumns []MovieDataColumn

	// Auth
	ApiKey *string
	BearerToken *string
}

func (o *allOptions) apply(options []Option) {
	for _, opt := range options {
		opt(o)
	}
}

func (o *allOptions) applyIncludeAdult(v url.Values) {
	if o.IncludeAdult != nil {
		v.Set("include_adult", fmt.Sprint(*o.IncludeAdult))
	}
}

func (o *allOptions) applyLanguage(v url.Values) {
	if o.Language != nil {
		v.Set("language", *o.Language)
	}
}

func (o *allOptions) applyMovieDataColumns(v url.Values) {
	if len(o.MovieDataColumns) > 0 {
		v.Set("append_to_response", appendToResponse(o.MovieDataColumns))
	}
}

type Option func(*allOptions)

func WithLanguage(language string) Option {
	return func(o *allOptions) {
		o.Language = &language
	}
}

func WithoutLanguage() Option {
	return func(o *allOptions) {
		o.Language = nil
	}
}

func WithIncludeAdult(include bool) Option {
	return func(o *allOptions) {
		o.IncludeAdult = &include
	}
}

func WithoutIncludeAdult() Option {
	return func(o *allOptions) {
		o.IncludeAdult = nil
	}
}

func WithStartPage(page int) Option {
	return func(o *allOptions) {
		o.StartPage = &page
	}
}

func WithoutStartPage() Option {
	return func(o *allOptions) {
		o.StartPage = nil
	}
}

func WithLimitPage(page int) Option {
	return func(o *allOptions) {
		o.LimitPage = &page
	}
}

func WithoutLimitPage() Option {
	return func(o *allOptions) {
		o.LimitPage = nil
	}
}

func WithMovieDataColumns(columns ...MovieDataColumn) Option {
	return func(o *allOptions) {
		o.MovieDataColumns = columns
	}
}

func WithoutMovieDataColumns() Option {
	return func(o *allOptions) {
		o.MovieDataColumns = nil
	}
}

func WithApiKey(apiKey string) Option {
	return func(o *allOptions) {
		o.ApiKey = &apiKey
	}
}

func WithoutApiKey() Option {
	return func(o *allOptions) {
		o.ApiKey = nil
	}
}

func WithBearerToken(token string) Option {
	return func(o *allOptions) {
		o.BearerToken = &token
	}
}

func WithoutBearerToken() Option {
	return func(o *allOptions) {
		o.BearerToken = nil
	}
}