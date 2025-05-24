package tmdb

import "context"

func WithDetails() detailsOption {
	return detailsOption{}
}

type detailsOption struct {}

func (d detailsOption) applyToGetMovieOptions(o *getMovieOptions) {
	o.wantDetails = true
}

func WithKeywords() keywordsOption {
	return keywordsOption{}
}
type keywordsOption struct {}

func (k keywordsOption) applyToGetMovieOptions(o *getMovieOptions) {
	o.wantKeywords = true
}

func WithContext(ctx context.Context) contextOption {
	return contextOption{ctx: ctx}
}

type contextOption struct {
	ctx context.Context
}

func (c contextOption) applyToGetMovieOptions(o *getMovieOptions) {
	o.useContext = &c.ctx
}