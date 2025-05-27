package tmdb

type GenreId int

type Genere struct {
	GenreId GenreId `json:"id"`
	Name    string  `json:"name"`
}
