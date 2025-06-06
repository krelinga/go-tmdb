package tmdb

type KeywordId int

type KeywordKey struct {
	Id KeywordId
}

type Keyword struct {
	KeywordKey

	// TODO: does this really need to be a pointer?
	Name *string
}
