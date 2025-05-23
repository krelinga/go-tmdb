package tmdb

type Gender int

var (
	GenderUnknown   Gender = 0
	GenderFemale           = 1
	GenderMale             = 2
	GenderNonBinary        = 3
)

type PersonId int
type ImdbPersonId string

type PersonSummary struct {
	Adult              bool         `json:"adult"`
	Gender             Gender       `json:"gender"`
	PersonId           PersonId     `json:"id"`
	KnownForDepartment string       `json:"known_for_department"`
	Name               string       `json:"name"`
	Popularity         float64      `json:"popularity"`
	ProfileImage       ProfileImage `json:"profile_path"`
}

type Person struct {
	PersonSummary
	AlsoKnownAs  []string     `json:"also_known_as"`
	Biography    string       `json:"biography"`
	Birthday     Date         `json:"birthday"`
	Deathday     Date         `json:"deathday"`
	Homepage     string       `json:"homepage"`
	ImdbId       ImdbPersonId `json:"imdb_id"`
	PlaceOfBirth string       `json:"place_of_birth"`
}
