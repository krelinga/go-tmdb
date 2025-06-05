package tmdb

type KeywordId int

type Keyword struct {
	Id KeywordId

	// TODO: does this really need to be a pointer?
	Name *string
}
