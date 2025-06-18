package tmdb

type GenreId int

type Genre struct {
	Key  GenreId
	Data GenreData
}

type GenreData struct {
	Name *string
}
