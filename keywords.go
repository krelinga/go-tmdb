package tmdb

type KeywordId int

type Keyword interface {
	// These methods never panic.
	Id() KeywordId
	Name() string
}

type keyword struct {
	id   KeywordId
	name string
}

func (k keyword) Id() KeywordId {
	return k.id
}

func (k keyword) Name() string {
	return k.name
}