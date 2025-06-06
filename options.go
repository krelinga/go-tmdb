package tmdb

import "context"

type baseOptions struct {
	useContext *context.Context
	rawReply   *[]byte
}

func WithKeywords() keywordsOption {
	return keywordsOption{}
}

type keywordsOption struct{}

func (k keywordsOption) applyToGetMovieOptions(o *getMovieOptions) {
	o.wantKeywords = true
}

func WithCredits() creditsOption {
	return creditsOption{}
}

type creditsOption struct{}

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

func (r rawReplyOption) applyToGetTvSeriesOptions(o *getTvSeriesOptions) {
	o.rawReply = r.reply
}

func (r rawReplyOption) applyToGetTvSeasonOptions(o *getTvSeasonOptions) {
	o.rawReply = r.reply
}

func WithAdult() adultOption {
	return adultOption{}
}

type adultOption struct{}

func (a adultOption) applyToSearchMoviesOptions(o *searchMoviesOptions) {
	o.wantAdult = true
}

func (a adultOption) applyToSearchTvSeriesOptions(o *searchTvSeriesOptions) {
	o.wantAdult = true
}

func WithReleaseDates() releaseDatesOption {
	return releaseDatesOption{}
}

type releaseDatesOption struct{}

func (r releaseDatesOption) applyToGetMovieOptions(o *getMovieOptions) {
	o.wantReleases = true
}

func WithExternalIds() externalIdsOption {
	return externalIdsOption{}
}

type externalIdsOption struct{}

func (e externalIdsOption) applyToGetTvEpisodeOptions(o *getTvEpisodeOptions) {
	o.wantExternalIds = true
}

func (e externalIdsOption) applyToGetTvSeriesOptions(o *getTvSeriesOptions) {
	o.wantExternalIds = true
}

func WithFirstAirDateYear(year int) firstAirDateYearOption {
	return firstAirDateYearOption(year)
}

type firstAirDateYearOption int

func (f firstAirDateYearOption) applyToSearchTvSeriesOptions(o *searchTvSeriesOptions) {
	o.firstAirDateYear = int(f)
}

func WithYear(year int) yearOption {
	return yearOption(year)
}

type yearOption int

func (y yearOption) applyToSearchTvSeriesOptions(o *searchTvSeriesOptions) {
	o.year = int(y)
}
