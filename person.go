package tmdb

type Gender int

var (
	GenderUnknown   Gender = 0
	GenderFemale    Gender = 1
	GenderMale      Gender = 2
	GenderNonBinary Gender = 3
)

type PersonId int
type ImdbPersonId string

type PersonCore struct {
	Gender             Gender       `json:"gender"`
	PersonId           PersonId     `json:"id"`
	Name               string       `json:"name"`
	ProfileImage       ProfileImage `json:"profile_path"`
}

type PersonSum struct {
	PersonCore
	Adult              bool         `json:"adult"`
	KnownForDepartment string       `json:"known_for_department"`
	Popularity         float64      `json:"popularity"`
}

type Person struct {
	PersonSum
	AlsoKnownAs  []string     `json:"also_known_as"`
	Biography    string       `json:"biography"`
	Birthday     DateYYYYMMDD `json:"birthday"`
	Deathday     DateYYYYMMDD `json:"deathday"`
	Homepage     string       `json:"homepage"`
	ImdbId       ImdbPersonId `json:"imdb_id"`
	PlaceOfBirth string       `json:"place_of_birth"`
}

type CreditId string

type CreditPerson struct {
	PersonSum
	OriginalName string   `json:"original_name"`
	CreditId     CreditId `json:"credit_id"`
}

type CastPerson struct {
	CreditPerson
	Character string `json:"character"`
	Order     int    `json:"order"`
}

type CrewPerson struct {
	CreditPerson
	Department string `json:"department"`
	Job        string `json:"job"`
}
