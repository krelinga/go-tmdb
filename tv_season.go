package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
)

type TvSeasonId int
type TvSeasonNumber int

type TvSeason struct {
	// TODO: what is this?
	UnderscoreId   string         `json:"_id"`
	AirDate        DateYYYYMMDD   `json:"air_date"`
	TvEpisodes     []*TvEpisode   `json:"episodes"`
	Name           string         `json:"name"`
	Overview       string         `json:"overview"`
	TvSeasonId     TvSeasonId     `json:"id"`
	PosterImage    PosterImage    `json:"poster_path"`
	TvSeasonNumber TvSeasonNumber `json:"season_number"`
	VoteAverage    float64        `json:"vote_average"`
}

func GetTvSeason(client Client, id TvSeasonId, number TvSeasonNumber, options ...GetTvSeasonOption) (*TvSeason, error) {
	o := getTvSeasonOptions{}
	for _, opt := range options {
		opt.applyToGetTvSeasonOptions(&o)
	}

	ctx := context.Background()
	if o.useContext != nil {
		ctx = *o.useContext
	}

	data, err := checkCode(client.Get(ctx, fmt.Sprintf("/tv/%d/season/%d", id, number), nil))
	if err != nil {
		return nil, err
	}
	if o.rawReply != nil {
		*o.rawReply = data
	}
	r := &TvSeason{}
	if err := json.Unmarshal(data, r); err != nil {
		return nil, err
	}
	return r, nil
}

type getTvSeasonOptions struct {
	baseOptions
}

type GetTvSeasonOption interface {
	applyToGetTvSeasonOptions(o *getTvSeasonOptions)
}
