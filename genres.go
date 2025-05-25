package tmdb

type GenereId int

type Genere struct {
	Id   GenereId    `json:"id"`
	Name string `json:"name"`
}