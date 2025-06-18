package tmdb

type SeasonId int

type SeasonKey struct {
	ShowId       ShowId
	SeasonNumber int
}

type Season struct {
	Key  SeasonKey
	Data SeasonData
}

type SeasonData struct {
	Id           *SeasonId
	AirDate      *DateYYYYMMDD
	EpisodeCount *int
	Name         *string
	Overview     *string
	Poster       *Image
}
