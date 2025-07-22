package tmdb

import (
	"time"

	"github.com/krelinga/go-views"
)

type ShowId int

type Show struct {
	Key  ShowId
	Data ShowData

	lastEpisodeToAir    *Episode
	networks            []*Network
	productionCompanies []*Company
	ProductionCountries []*Country // TODO: threat this the same as other edges.
	seasons             []*Season
	SpokenLanguages     []*Language // TODO: threat this the same as other edges.
}

func (s *Show) SetLastEpisodeToAir(e *Episode) {
	if s.lastEpisodeToAir != nil {
		s.lastEpisodeToAir.lastToAir = false
	}
	s.lastEpisodeToAir = e
	e.lastToAir = true
}

func (s *Show) LastEpisodeToAir() *Episode {
	return s.lastEpisodeToAir
}

func (s *Show) AddNetwork(n *Network) {
	if s.Networks().Has(n) {
		return
	}
	if s.networks == nil {
		s.networks = make([]*Network, 0, 1)
	}
	s.networks = append(s.networks, n)

	if n.shows == nil {
		n.shows = make([]*Show, 0, 1)
	}
	n.shows = append(n.shows, s)
}

func (s *Show) Networks() views.Bag[*Network] {
	return views.BagOfSlice[*Network]{S: s.networks}
}

func (s *Show) AddProductionCompany(c *Company) {
	if s.ProductionCompanies().Has(c) {
		return
	}
	if s.productionCompanies == nil {
		s.productionCompanies = make([]*Company, 0, 1)
	}
	s.productionCompanies = append(s.productionCompanies, c)

	if c.shows == nil {
		c.shows = make([]*Show, 0, 1)
	}
	c.shows = append(c.shows, s)
}

func (s *Show) ProductionCompanies() views.Bag[*Company] {
	return views.BagOfSlice[*Company]{S: s.productionCompanies}
}

func (s *Show) Seasons() views.Bag[*Season] {
	return views.BagOfSlice[*Season]{S: s.seasons}
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
