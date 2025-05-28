package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
)

type TvEpisodeId int
type TvEpisodeNumber int

type TvEpisode struct {
	TvEpisodeSum
	Crew       []*CrewPerson `json:"crew"`
	GuestStars []*CastPerson `json:"guest_stars"`
}

type TvEpisodeSum struct {
	AirDate         DateYYYYMMDD    `json:"air_date"`
	TvEpisodeNumber TvEpisodeNumber `json:"episode_number"`
	EpisodeType     string          `json:"episode_type"`
	TvEpisodeId     TvEpisodeId     `json:"id"`
	Name            string          `json:"name"`
	Overview        string          `json:"overview"`
	ProductionCode  string          `json:"production_code"`
	Runtime         Minutes         `json:"runtime"`
	TvSeasonNumber  TvSeasonNumber  `json:"season_number"`
	TvSeriesId      TvSeriesId      `json:"show_id"`
	StillImage      PosterImage     `json:"still_path"`
	VoteAverage     float64         `json:"vote_average"`
	VoteCount       int             `json:"vote_count"`
}

type GetTvEpisodeReply struct {
	*TvEpisode

	// TODO: add fields that can be appended to the reply.
}

func GetTvEpisode(client Client, tvSeriesId TvSeriesId, seasonNumber TvSeasonNumber, episodeNumber TvEpisodeNumber, options ...GetTvEpisodeOption) (*GetTvEpisodeReply, error) {
	o := getTvEpisodeOptions{}
	for _, opt := range options {
		opt.applyToGetTvEpisodeOptions(&o)
	}

	ctx := context.Background()
	if o.useContext != nil {
		ctx = *o.useContext
	}

	data, err := checkCode(client.Get(ctx, fmt.Sprintf("/tv/%d/season/%d/episode/%d", tvSeriesId, seasonNumber, episodeNumber), nil))
	if err != nil {
		return nil, err
	}
	if o.rawReply != nil {
		*o.rawReply = data
	}
	r := &GetTvEpisodeReply{}
	if err := json.Unmarshal(data, r); err != nil {
		return nil, err
	}
	return r, nil
}

type GetTvEpisodeOption interface {
	applyToGetTvEpisodeOptions(*getTvEpisodeOptions)
}

type getTvEpisodeOptions struct {
	baseOptions
}
