package tmdb

import "github.com/krelinga/go-views"

type SeasonId int

type SeasonKey struct {
	ShowId       ShowId
	SeasonNumber int
}

type Season struct {
	Key  SeasonKey
	Data SeasonData

	show     *Show
	episodes []*Episode
}

func (s *Season) Show() *Show {
	return s.show
}

func (s *Season) Episodes() views.Bag[*Episode] {
	return views.BagOfSlice[*Episode]{S: s.episodes}
}

type SeasonData struct {
	Id           *SeasonId
	AirDate      *DateYYYYMMDD
	EpisodeCount *int
	Name         *string
	Overview     *string
	Poster       *Image
}
