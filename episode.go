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

	season    *Season
	lastToAir bool
}

func (e *Episode) Show() *Show {
	return e.season.Show()
}

func (e *Episode) SetLastToAir() {
	e.Show().SetLastEpisodeToAir(e)
}

func (e *Episode) LastToAir() bool {
	return e.lastToAir
}

func (e *Episode) Season() *Season {
	return e.season
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
