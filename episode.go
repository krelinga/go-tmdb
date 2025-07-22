package tmdb

import "time"

type EpisodeId int

type EpisodeKey struct {
	SeasonNumber  int
	EpisodeNumber int
	ShowId        ShowId
}

type Episode struct {
	Key  EpisodeKey
	Data EpisodeData

	show      *Show
	lastToAir bool
}

func (e *Episode) Show() *Show {
	return e.show
}

func (e *Episode) SetLastToAir() {
	e.show.SetLastEpisodeToAir(e)
}

func (e *Episode) LastToAir() bool {
	return e.lastToAir
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
