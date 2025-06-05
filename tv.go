package tmdb

import "time"

type TvId int

type Tv struct {
	Id TvId

	Adult *bool
	Backdrop *Image
	CreatedBy []*Credit
	EpisodeRunTime []time.Duration
	FirstAirDate *DateYYYYMMDD
	Genres []Genre
	Homepage *string
	InProduction *bool
	Languages []string  // TODO: use Language type?
	LastAirDate *DateYYYYMMDD
	LastEpisodeToAir *Episode
	Name *string
	NextEpisodeToAir *string // TODO: what's up with this?
	Networks []*Network
	NumberOfEpisodes *int
	NumberOfSeasons *int
	OriginCountry []string
	OriginalLanguage *string
	OriginalName *string
	Overview *string
	Popularity *float64
	Poster *Image
	ProductionCompanies []*Company
	ProductionCountries []*Country
	Seasons []*Season
	SpokenLanguages []*Language
	Status *string
	Tagline *string
	Type *string
	VoteAverage *float64
	VoteCount *int
}