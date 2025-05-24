package tmdb

import "context"

type baseOptions struct {
	useContext *context.Context
	rawReply *[]byte
}

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

func WithCredits() creditsOption {
	return creditsOption{}
}

type creditsOption struct {}

func (c creditsOption) applyToGetMovieOptions(o *getMovieOptions) {
	o.wantCredits = true
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

func (c contextOption) applyToGetConfigurationOptions(o *getConfigruationOptions) {
	o.useContext = &c.ctx
}

func WithRawReply(reply *[]byte) rawReplyOption {
	return rawReplyOption{reply: reply}
}

type rawReplyOption struct {
	reply *[]byte
}

func (r rawReplyOption) applyToGetMovieOptions(o *getMovieOptions) {
	o.rawReply = r.reply
}

func (r rawReplyOption) applyToGetConfigurationOptions(o *getConfigruationOptions) {
	o.rawReply = r.reply
}