package tmdb

type GenreId int

type Genre interface {
	Id() GenreId
	Name() string
}

type genre struct {
	id   GenreId
	name string
}

func (g genre) Id() GenreId {
	return g.id
}

func (g genre) Name() string {
	return g.name
}