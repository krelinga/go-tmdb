package tmdb

type MovieId int

type Movie interface {
	Id() MovieId
}

type movie struct {
	id MovieId
}

func (m *movie) Id() MovieId {
	return m.id
}