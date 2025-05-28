package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
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
	TvSeriesId      TvSeriesId      `json:"show_id"` // This isn't always present in the json, but in all of those cases we can add it.
	StillImage      StillImage      `json:"still_path"`
	VoteAverage     float64         `json:"vote_average"`
	VoteCount       int             `json:"vote_count"`
}

type GetTvEpisodeReply struct {
	*TvEpisode

	// TODO: add fields that can be appended to the reply.
	ExternalIds *TvEpisodeExternalIds `json:"external_ids,omitempty"`
}

type ImdbEpisodeId string
type TheTvdbEpisodeId int
type WikidataEpisodeId string

type TvEpisodeExternalIds struct {
	TvEpisodeId       TvEpisodeId       `json:"id"`
	ImdbEpisodeId     ImdbEpisodeId     `json:"imdb_id"`
	TheTvdbEpisodeId  TheTvdbEpisodeId  `json:"thetvdb_id"`
	WikidataEpisodeId WikidataEpisodeId `json:"wikidata_id"`
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

	appendTo := []string{}
	if o.wantExternalIds {
		appendTo = append(appendTo, "external_ids")
	}

	params := make(GetParams)
	if len(appendTo) > 0 {
		params["append_to_response"] = strings.Join(appendTo, ",")
	}

	data, err := checkCode(client.Get(ctx, fmt.Sprintf("/tv/%d/season/%d/episode/%d", tvSeriesId, seasonNumber, episodeNumber), params))
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
	// This is a case where this field isn't set in the JSON, but it is trivial for us to set it.
	r.TvSeriesId = tvSeriesId
	return r, nil
}

type GetTvEpisodeOption interface {
	applyToGetTvEpisodeOptions(*getTvEpisodeOptions)
}

type getTvEpisodeOptions struct {
	baseOptions
	wantExternalIds bool
}
