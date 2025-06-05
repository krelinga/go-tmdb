package tmdb

type options struct {
	Language *string
	IncludeAdult *bool
	StartPage *int
	LimitPage *int
}

type Option func(*options)

func WithLanguage(language string) Option {
	return func(o *options) {
		o.Language = &language
	}
}

func WithoutLanguage() Option {
	return func(o *options) {
		o.Language = nil
	}
}

func WithIncludeAdult(include bool) Option {
	return func(o *options) {
		o.IncludeAdult = &include
	}
}

func WithoutIncludeAdult() Option {
	return func(o *options) {
		o.IncludeAdult = nil
	}
}

func WithStartPage(page int) Option {
	return func(o *options) {
		o.StartPage = &page
	}
}

func WithoutStartPage() Option {
	return func(o *options) {
		o.StartPage = nil
	}
}

func WithLimitPage(page int) Option {
	return func(o *options) {
		o.LimitPage = &page
	}
}

func WithoutLimitPage() Option {
	return func(o *options) {
		o.LimitPage = nil
	}
}