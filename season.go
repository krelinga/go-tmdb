package tmdb

type SeasonId int

type Season struct {
	Id SeasonId
	SeasonNumber int
	TvId TvId

	AirDate *DateYYYYMMDD
	EpisodeCount *int
	Name *string
	Overview *string
	Poster *Image
}