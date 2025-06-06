package tmdb

type SeasonId int

type SeasonKey struct {
	TvId TvId
	SeasonNumber int
}

type Season struct {
	SeasonKey

	Id *SeasonId
	AirDate *DateYYYYMMDD
	EpisodeCount *int
	Name *string
	Overview *string
	Poster *Image
}