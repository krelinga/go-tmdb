package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

type MovieId int
type ImdbId string

type GetMovieReply struct {
	*Movie

	// Additional bits that can be fetched at the same time.
	Keywords *MovieKeywords     `json:"keywords,omitempty"`
	Credits  *MovieCredits      `json:"credits,omitempty"`
	Releases *MovieReleaseDates `json:"release_dates,omitempty"`
}

type Movie struct {
	MovieSum
	BelongsToCollection        string                  `json:"belongs_to_collection,omitempty"`
	Budget                     int                     `json:"budget"`
	Genres                     []*Genre                `json:"genres"`
	Homepage                   string                  `json:"homepage,omitempty"`
	ImdbId                     ImdbId                  `json:"imdb_id,omitempty"`
	ProductionCompanyShorts    []*ProductionCompanySum `json:"production_companies"`
	ProductionCountrySummaries []*CountrySum           `json:"production_countries"`
	Revenue                    int                     `json:"revenue"`
	Runtime                    Minutes                 `json:"runtime"`
	SpokenLanguages            []*Language             `json:"spoken_languages"`
	Status                     string                  `json:"status"`
	Tagline                    string                  `json:"tagline"`
}

type MovieSum struct {
	Adult            bool          `json:"adult"`
	BackdropImage    BackdropImage `json:"backdrop_path"`
	MovieId          MovieId       `json:"id"`
	OriginalLanguage string        `json:"original_language"`
	OriginalTitle    string        `json:"original_title"`
	Overview         string        `json:"overview"`
	Popularity       float64       `json:"popularity"`
	PosterImage      PosterImage   `json:"poster_path"`
	ReleaseDate      DateYYYYMMDD  `json:"release_date"`
	Title            string        `json:"title"`
	Video            bool          `json:"video"`
	VoteAverage      float64       `json:"vote_average"`
	VoteCount        int           `json:"vote_count"`
}

type MovieKeywords struct {
	MovieId  MovieId    `json:"id"`
	Keywords []*Keyword `json:"keywords"`
}

type MovieCredits struct {
	MovieId MovieId      `json:"id"`
	Cast    []*MovieCast `json:"cast"`
	Crew    []*MovieCrew `json:"crew"`
}

type MovieCastId int
type CreditId string

type MovieCast struct {
	CastPerson
	CastId    MovieCastId `json:"cast_id"`
}

type MovieCrew struct {
	CreditPerson
	Department string `json:"department"`
	Job        string `json:"job"`
}

type MovieReleaseType int

const (
	MovieReleaseTypeUnknown           MovieReleaseType = 0
	MovieReleaseTypePremiere          MovieReleaseType = 1
	MovieReleaseTypeTheatricalLimited MovieReleaseType = 2
	MovieReleaseTypeTheatrical        MovieReleaseType = 3
	MovieReleaseTypeDigital           MovieReleaseType = 4
	MovieReleaseTypePhysical          MovieReleaseType = 5
	MovieReleaseTypeTv                MovieReleaseType = 6
)

type MovieReleaseDates struct {
	MovieId               MovieId                `json:"id"`
	MovieReleaseCountries []*MovieReleaseCountry `json:"results"`
}

type MovieReleaseCountry struct {
	Iso3166_1         CountryIso3166_1    `json:"iso_3166_1"`
	MovieReleaseDates []*MovieReleaseDate `json:"release_dates"`
}

type MovieReleaseDate struct {
	Certification string `json:"certification"`
	// TODO: Descriptors?
	Language         LanguageIso639_1 `json:"iso_639_1"`
	Note             string           `json:"note"`
	ReleaseDate      DateRfc3339      `json:"release_date"`
	MovieReleaseType MovieReleaseType `json:"type"`
}

func GetMovie(client Client, movieId MovieId, options ...GetMovieOption) (*GetMovieReply, error) {
	o := getMovieOptions{}
	for _, opt := range options {
		opt.applyToGetMovieOptions(&o)
	}

	appends := []string{}
	if o.wantKeywords {
		appends = append(appends, "keywords")
	}
	if o.wantCredits {
		appends = append(appends, "credits")
	}
	if o.wantReleases {
		appends = append(appends, "release_dates")
	}

	params := GetParams{}
	if len(appends) > 0 {
		params["append_to_response"] = strings.Join(appends, ",")
	}
	ctx := context.Background()
	if o.useContext != nil {
		ctx = *o.useContext
	}
	endpoint := fmt.Sprintf("/movie/%d", movieId)
	data, err := checkCode(client.Get(ctx, endpoint, params))
	if err != nil {
		return nil, err
	}
	if o.rawReply != nil {
		*o.rawReply = data
	}
	m := &GetMovieReply{}
	if err := json.Unmarshal(data, m); err != nil {
		return nil, fmt.Errorf("unmarshalling movie: %w", err)
	}

	return m, nil
}

type getMovieOptions struct {
	baseOptions
	wantKeywords bool
	wantCredits  bool
	wantReleases bool
}

type GetMovieOption interface {
	applyToGetMovieOptions(o *getMovieOptions)
}
