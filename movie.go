package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

type MovieId int
type ImdbId string

type Movie struct {
	Adult                      bool                        `json:"adult"`
	BackdropImage              BackdropImage               `json:"backdrop_path"`
	BelongsToCollection        string                      `json:"belongs_to_collection,omitempty"`
	Budget                     int                         `json:"budget"`
	Genres                     []*Genere                   `json:"genres"`
	Homepage                   string                      `json:"homepage,omitempty"`
	MovieId                    MovieId                     `json:"id"`
	ImdbId                     ImdbId                      `json:"imdb_id,omitempty"`
	OriginalLanguage           string                      `json:"original_language"`
	OriginalTitle              string                      `json:"original_title"`
	Overview                   string                      `json:"overview"`
	Popularity                 float64                     `json:"popularity"`
	PosterImage                PosterImage                 `json:"poster_path"`
	ProductionCompanySummaries []*ProductionCompanySummary `json:"production_companies"`
	ProductionCountrySummaries []*CountrySummary           `json:"production_countries"`
	RelaseDate                 Date                        `json:"release_date"`
	Revenue                    int                         `json:"revenue"`
	Runtime                    Minutes                     `json:"runtime"`
	SpokenLanguages            []*Language                 `json:"spoken_languages"`
	Status                     string                      `json:"status"`
	Tagline                    string                      `json:"tagline"`
	Title                      string                      `json:"title"`
	Video                      bool                        `json:"video"`
	VoteAverage                float64                     `json:"vote_average"`
	VoteCount                  int                         `json:"vote_count"`

	// Additional bits that can be fetched at the same time.
	Keywords *MovieKeywords `json:"keywords,omitempty"`
	Credits *MovieCredits `json:"credits,omitempty"`
}

type MovieKeywords struct {
	MovieId  MovieId    `json:"id"`
	Keywords []*Keyword `json:"keywords"`
}

type MovieCredits struct {
	MovieId MovieId `json:"id"`
}

func GetMovie(client Client, movieId MovieId, options ...GetMovieOption) (*Movie, error) {
	o := getMovieOptions{}
	for _, opt := range options {
		opt.applyToGetMovieOptions(&o)
	}

	appends := []string{}
	if o.wantKeywords {
		appends = append(appends, "keywords")
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
	data, err := client.Get(ctx, endpoint, params)
	if err != nil {
		return nil, err
	}
	if o.rawReply != nil {
		*o.rawReply = data
	}
	m := &Movie{}
	if err := json.Unmarshal(data, m); err != nil {
		return nil, fmt.Errorf("unmarshalling movie: %w", err)
	}

	return m, nil
}

type getMovieOptions struct {
	baseOptions
	wantDetails  bool
	wantKeywords bool
}

type GetMovieOption interface {
	applyToGetMovieOptions(o *getMovieOptions)
}
