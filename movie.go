package tmdb

import (
	"fmt"
	"time"

	"github.com/krelinga/go-views"
)

type MovieId int

type WikidataMovieId string

type MovieDataColumn int

const (
	movieDataNone MovieDataColumn = iota
	MovieDataCredits
	MovieDataExternalIds
	MovieDataKeywords
	movieDataMax
)

func (d MovieDataColumn) String() string {
	switch d {
	case movieDataNone:
		return "movieDataNone"
	case MovieDataCredits:
		return "MovieDataCredits"
	case MovieDataExternalIds:
		return "MovieDataExternalIds"
	case MovieDataKeywords:
		return "MovieDataKeywords"
	case movieDataMax:
		return "movieDataMax"
	default:
		return fmt.Sprintf("MovieDataColumn(%d)", d)
	}
}

func (d MovieDataColumn) Endpoint() string {
	switch d {
	case MovieDataCredits:
		return "credits"
	case MovieDataExternalIds:
		return "external_ids"
	case MovieDataKeywords:
		return "keywords"
	default:
		panic(fmt.Sprintf("no endpoint for %s", d))
	}
}

type Movie struct {
	Key  MovieId
	Data MovieData

	// MovieDataCredits
	Cast []*Credit
	Crew []*Credit

	// MovieDataKeywords
	Keywords []*Keyword

	genres []*Genre
}

func (m *Movie) AddGenre(g *Genre) {
	if m.Genres().Has(g) {
		return // already added
	}
	if m.genres == nil {
		m.genres = make([]*Genre, 0, 1)
	}
	m.genres = append(m.genres, g)

	if g.movies == nil {
		g.movies = make([]*Movie, 0, 1)
	}
	g.movies = append(g.movies, m)
}

func (m *Movie) Genres() views.Bag[*Genre] {
	return views.BagOfSlice[*Genre]{S: m.genres}
}

type MovieData struct {
	Adult               *bool
	Backdrop            *Image
	BelongsToCollection *string
	Budget              *int
	Homepage            *string
	ImdbId              *string
	OriginalLanguage    *string
	OriginalTitle       *string
	Overview            *string
	Popularity          *float64
	Poster              *Image
	ProductionCompanies []*Company
	ProductionCountries []*Country
	ReleaseDate         *DateYYYYMMDD
	Revenue             *int
	Runtime             *time.Duration
	SpokenLanguages     []*Language
	Status              *string
	Tagline             *string
	Title               *string
	Video               *bool
	VoteAverage         *float64
	VoteCount           *int

	// MovieDataExternalIds
	WikidataId  *string
	FacebookId  *string
	InstagramId *string
	TwitterId   *string
}
