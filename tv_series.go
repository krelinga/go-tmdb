package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
)

type TvSeriesId int

type TvSeries struct {
	TvSeriesSum
	CreatedBy        []*TvSeriesCreator `json:"created_by"`
	EpisodeRunTimes  []int              `json:"episode_run_time"`
	Genres           []*Genre           `json:"genres"`
	Homepage         string             `json:"homepage"`
	InProduction     bool               `json:"in_production"`
	Languages        []string           `json:"languages"`
	LastAirDate      DateYYYYMMDD       `json:"last_air_date"`
	LastEpisodeToAir *TvEpisodeSum      `json:"last_episode_to_air"`
	NextEpisodeToAir string             `json:"next_episode_to_air"`
	TvNetworks       []*TvNetwork       `json:"networks"`
	NumberOfEpisodes int                `json:"number_of_episodes"`
	NumberOfSeasons  int                `json:"number_of_seasons"`
	// TODO: ProductionCompanies
	// TODO: ProductionCountries
	TvSeasons []*TvSeasonSum `json:"seasons"`
	// TODO: SpokenLanguages
	Status  string `json:"status"`
	Tagline string `json:"tagline"`
	Type    string `json:"type"`
}

type TvSeriesSum struct {
	Adult            bool               `json:"adult"`
	BackdropImage    BackdropImage      `json:"backdrop_path"`
	TvSeriesId       TvSeriesId         `json:"id"`
	OriginCountries  []CountryIso3166_1 `json:"origin_country"`
	OriginalLanguage string             `json:"original_language"`
	OriginalName     string             `json:"original_name"`
	Overview         string             `json:"overview"`
	Popularity       float64            `json:"popularity"`
	PosterImage      PosterImage        `json:"poster_path"`
	FirstAirDate     DateYYYYMMDD       `json:"first_air_date"`
	Name             string             `json:"name"`
	VoteAverage      float64            `json:"vote_average"`
	VoteCount        int                `json:"vote_count"`
}

type TvSeriesCreator struct {
	PersonCore
	CreditId CreditId `json:"credit_id"`
}

func GetTvSeries(client Client, id TvSeriesId, options ...GetTvSeriesOption) (*TvSeries, error) {
	o := getTvSeriesOptions{}
	for _, opt := range options {
		opt.applyToGetTvSeriesOptions(&o)
	}

	ctx := context.Background()
	if o.useContext != nil {
		ctx = *o.useContext
	}

	data, err := checkCode(client.Get(ctx, fmt.Sprintf("/tv/%d", id), nil))
	if err != nil {
		return nil, err
	}
	if o.rawReply != nil {
		*o.rawReply = data
	}
	r := &TvSeries{}
	if err := json.Unmarshal(data, r); err != nil {
		return nil, err
	}
	return r, nil
}

type getTvSeriesOptions struct {
	baseOptions
}

type GetTvSeriesOption interface {
	applyToGetTvSeriesOptions(o *getTvSeriesOptions)
}
