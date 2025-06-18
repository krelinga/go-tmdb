package tmdb

import "time"

type ShowId int

type Show struct {
	Key  ShowId
	Data ShowData

	LastEpisodeToAir    *Episode
	Networks            []*Network
	ProductionCompanies []*Company
	ProductionCountries []*Country
	Seasons             []*Season
	SpokenLanguages     []*Language
}

type ShowData struct {
	Adult            *bool
	Backdrop         *Image
	CreatedBy        []*Credit
	EpisodeRunTime   []time.Duration
	FirstAirDate     *DateYYYYMMDD
	Genres           []Genre
	Homepage         *string
	InProduction     *bool
	Languages        []string // TODO: use Language type?
	LastAirDate      *DateYYYYMMDD
	Name             *string
	NextEpisodeToAir *string // TODO: what's up with this?
	NumberOfEpisodes *int
	NumberOfSeasons  *int
	OriginCountry    []string
	OriginalLanguage *string
	OriginalName     *string
	Overview         *string
	Popularity       *float64
	Poster           *Image
	Status           *string
	Tagline          *string
	Type             *string
	VoteAverage      *float64
	VoteCount        *int
}
