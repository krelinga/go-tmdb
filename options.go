package tmdb

type allOptions struct {
	Language *string
	IncludeAdult *bool
	StartPage *int
	LimitPage *int
	MovieDataColumns []MovieDataColumn
}

func (o *allOptions) apply(options []Option) {
	for _, opt := range options {
		opt(o)
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