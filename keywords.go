package tmdb

type KeywordId int

type Keyword interface {
	// These methods never panic.
	Id() KeywordId
	Name() string
}
