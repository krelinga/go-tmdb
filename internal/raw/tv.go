package raw

type Tv struct {
	Adult *bool `json:"adult,omitempty"`
	BackdropPath string `json:"backdrop_path"`
	CreatedBy []*TvCreatedBy `json:"created_by"`
	EpisodeRunTime []int `json:"episode_run_time"`
	FirstAirDate string `json:"first_air_date"`
	Genres []Genre `json:"genres"`
	Homepage string `json:"homepage"`
	Id int `json:"id"`
	InProduction *bool `json:"in_production"`
	Languages []string `json:"languages"`
	LastAirDate string `json:"last_air_date"`
	LastEpisodeToAir *TvEpisode `json:"last_episode_to_air"`
	Name string `json:"name"`
	NextEpisodeToAir string `json:"next_episode_to_air"`
	Networks []*TvNetwork `json:"networks"`
	NumberOfEpisodes int `json:"number_of_episodes"`
	NumberOfSeasons int `json:"number_of_seasons"`
	OriginCountry []string `json:"origin_country"`
	OriginalLanguage string `json:"original_language"`
	OriginalName string `json:"original_name"`
	Overview string `json:"overview"`
	Popularity float64 `json:"popularity"`
	PosterPath string `json:"poster_path"`
	ProductionCompanies []*TvCompany `json:"production_companies"`
	ProductionCountries []*TvCountry `json:"production_countries"`
	Seasons []*TvSeason `json:"seasons"`
	SpokenLanguages []*TvSpokenLanguage `json:"spoken_languages"`
	Status string `json:"status"`
	Tagline string `json:"tagline"`
	Type string `json:"type"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount int `json:"vote_count"`
}

func (t *Tv) SetDefaults() {
	if t.Adult == nil {
		t.Adult = new(bool)
		*t.Adult = true
	}
	if t.InProduction == nil {
		t.InProduction = new(bool)
		*t.InProduction = true
	}
}

type TvCreatedBy struct {
	Id int `json:"id"`
	CreditId string `json:"credit_id"`
	Name string `json:"name"`
	Gender int `json:"gender"`
	ProfilePath string `json:"profile_path"`
}

type TvEpisode struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Overview string `json:"overview"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount int `json:"vote_count"`
	AirDate string `json:"air_date"`
	EpisodeNumber int `json:"episode_number"`
	ProductionCode string `json:"production_code"`
	Runtime int `json:"runtime"`
	SeasonNumber int `json:"season_number"`
	ShowId int `json:"show_id"`
	StillPath string `json:"still_path"`
}

type TvNetwork struct {
	Id int `json:"id"`
	LogoPath string `json:"logo_path"`
	Name string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

type TvCompany struct {
	Id int `json:"id"`
	LogoPath string `json:"logo_path"`
	Name string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

type TvCountry struct {
	Iso3166_1 string `json:"iso_3166_1"`
	Name string `json:"name"`
}

type TvSeason struct {
	AirDate string `json:"air_date"`
	EpisodeCount int `json:"episode_count"`
	Id int `json:"id"`
	Name string `json:"name"`
	Overview string `json:"overview"`
	PosterPath string `json:"poster_path"`
	SeasonNumber int `json:"season_number"`
	VoteAverage float64 `json:"vote_average"`
}

type TvSpokenLanguage struct {
	EnglishName string `json:"english_name"`
	Iso639_1 string `json:"iso_639_1"`
	Name string `json:"name"`
}