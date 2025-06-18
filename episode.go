package tmdb

import "time"

type EpisodeId int

type EpisodeKey struct {
	SeasonNumber  int
	EpisodeNumber int
	TvId          TvId
}

type Episode struct {
	Key  EpisodeKey
	Data EpisodeData
}

type EpisodeData struct {
	Id             *EpisodeId
	Name           *string
	Overview       *string
	VoteAverage    *float64
	VoteCount      *int
	AirDate        *DateYYYYMMDD
	ProductionCode *string
	Runtime        *time.Duration
	Still          *Image
}
