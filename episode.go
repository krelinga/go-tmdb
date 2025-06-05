package tmdb

import "time"

type EpisodeId int

type Episode struct {
	Id            EpisodeId
	SeasonNumber  int
	EpisodeNumber int
	TvId          TvId

	Name           *string
	Overview       *string
	VoteAverage    *float64
	VoteCount      *int
	AirDate        *DateYYYYMMDD
	ProductionCode *string
	Runtime        *time.Duration
	Still          *Image
}
