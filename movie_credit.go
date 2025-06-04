package tmdb

type MovieCastId int

type MovieCast interface {
	Credit
	Character() string
	CastId() MovieCastId
	Order() int
}

type MovieCrew interface {
	Credit
	Department() string
	Job() string
}