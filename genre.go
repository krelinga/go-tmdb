package tmdb

type GenreId int

type GenreKey struct {
	Id GenreId
}

type Genre struct {
	GenreKey
	Name *string
}
