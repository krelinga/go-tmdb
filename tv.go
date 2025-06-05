package tmdb

type TvId int

type Tv struct {
	Id TvId

	Adult *bool
	Backdrop *Image
	Genres []Genre
	OriginCountry []string
	OriginalLanguage *string
	OriginalName *string
	Overview *string
	Popularity *float64
	Poster *Image
	FirstAirDate *DateYYYYMMDD
	Name *string
	VoteAverage *float64
	VoteCount *int
}