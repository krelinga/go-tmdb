package tmdb

type GenreId int

type Genere struct {
	Id   GenreId `json:"id"`
	Name string  `json:"name"`
}
