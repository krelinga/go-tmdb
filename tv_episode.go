package tmdb

type TvEpisodeId int
type TvEpisodeNumber int

type TvEpisode struct {
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
	// TODO: Crew
	// TODO: GuestStars
}
