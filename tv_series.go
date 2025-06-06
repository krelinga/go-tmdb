package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
)

type TvSeriesId int

type TvSeries struct {
	TvSeriesSum
	CreatedBy             []*TvSeriesCreator      `json:"created_by"`
	EpisodeRunTimes       []int                   `json:"episode_run_time"`
	Genres                []*Genre                `json:"genres"`
	Homepage              string                  `json:"homepage"`
	InProduction          bool                    `json:"in_production"`
	Languages             []string                `json:"languages"`
	LastAirDate           DateYYYYMMDD            `json:"last_air_date"`
	LastEpisodeToAir      *TvEpisodeSum           `json:"last_episode_to_air"`
	NextEpisodeToAir      *TvEpisodeSum           `json:"next_episode_to_air"`
	TvNetworks            []*TvNetwork            `json:"networks"`
	NumberOfEpisodes      int                     `json:"number_of_episodes"`
	NumberOfSeasons       int                     `json:"number_of_seasons"`
	ProductionCompanySums []*ProductionCompanySum `json:"production_companies"`
	ProductionCountries   []*CountrySum           `json:"production_countries"`
	TvSeasons             []*TvSeasonSum          `json:"seasons"`
	SpokenLanguages       []*Language             `json:"spoken_languages"`
	Status                string                  `json:"status"`
	Tagline               string                  `json:"tagline"`
	Type                  string                  `json:"type"`
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

type GetTvSeriesReply struct {
	*TvSeries

	ExternalIds *TvSeriesExternalIds `json:"external_ids,omitempty"`
}

type FacebookTvSeriesId string
type ImdbTvSeriesId string
type InstagramTvSeriesId string
type TheTvdbTvSeriesId int
type TwitterTvSeriesId string
type WikidataTvSeriesId string

type TvSeriesExternalIds struct {
	FacebookTvSeriesId  FacebookTvSeriesId  `json:"facebook_id"`
	ImdbTvSeriesId      ImdbTvSeriesId      `json:"imdb_id"`
	InstagramTvSeriesId InstagramTvSeriesId `json:"instagram_id"`
	TheTvdbTvSeriesId   TheTvdbTvSeriesId   `json:"tvdb_id"`
	TwitterTvSeriesId   TwitterTvSeriesId   `json:"twitter_id"`
	WikidataTvSeriesId  WikidataTvSeriesId  `json:"wikidata_id"`
}

func GetTvSeries(client Client, id TvSeriesId, options ...GetTvSeriesOption) (*GetTvSeriesReply, error) {
	o := getTvSeriesOptions{}
	for _, opt := range options {
		opt.applyToGetTvSeriesOptions(&o)
	}

	ctx := context.Background()
	if o.useContext != nil {
		ctx = *o.useContext
	}

	params := make(GetParams)
	if o.wantExternalIds {
		params["append_to_response"] = "external_ids"
	}

	data, err := checkCode(client.Get(ctx, fmt.Sprintf("/tv/%d", id), params))
	if err != nil {
		return nil, err
	}
	if o.rawReply != nil {
		*o.rawReply = data
	}
	r := &GetTvSeriesReply{}
	if err := json.Unmarshal(data, r); err != nil {
		return nil, err
	}
	return r, nil
}

type getTvSeriesOptions struct {
	baseOptions
	wantExternalIds bool
}

type GetTvSeriesOption interface {
	applyToGetTvSeriesOptions(o *getTvSeriesOptions)
}
