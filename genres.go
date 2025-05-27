package tmdb

type GenreId int

type Genre struct {
	GenreId GenreId `json:"id"`
	Name    string  `json:"name"`
}
