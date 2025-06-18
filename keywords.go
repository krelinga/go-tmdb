package tmdb

type KeywordId int

type Keyword struct {
	Key  KeywordId
	Data KeywordData
}

type KeywordData struct {
	// TODO: does this really need to be a pointer?
	Name *string
}
