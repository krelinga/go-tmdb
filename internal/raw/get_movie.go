package raw

type GetMovie struct {
	*GetMovieDetails

	Credits     *GetMovieCredits     `json:"credits"`
	Keywords    *GetMovieKeywords    `json:"keywords"`
	ExternalIds *GetMovieExternalIds `json:"external_ids"`
}

type GetMovieDetails struct {
	Adult               *bool                        `json:"adult"`
	BackdropPath        string                       `json:"backdrop_path"`
	BelongsToCollection string                       `json:"belongs_to_collection"`
	Budget              int                          `json:"budget"`
	Genres              []*Genre                     `json:"genres"`
	Homepage            string                       `json:"homepage"`
	ImdbId              string                       `json:"imdb_id"`
	OriginalLanguage    string                       `json:"original_language"`
	OriginalTitle       string                       `json:"original_title"`
	Overview            string                       `json:"overview"`
	Popularity          float64                      `json:"popularity"`
	PosterPath          string                       `json:"poster_path"`
	ProductionCompanies []*GetMovieProductionCompany `json:"production_companies"`
	ProductionCountries []*GetMovieProductionCountry `json:"production_countries"`
	ReleaseDate         string                       `json:"release_date"`
	Revenue             int                          `json:"revenue"`
	Runtime             int                          `json:"runtime"`
	SpokenLanguages     []*GetMovieSpokenLanguage    `json:"spoken_languages"`
	Status              string                       `json:"status"`
	Tagline             string                       `json:"tagline"`
	Title               string                       `json:"title"`
	Video               bool                         `json:"video"`
	VoteAverage         float64                      `json:"vote_average"`
	VoteCount           int                          `json:"vote_count"`
}

func (g *GetMovie) SetDefaults() {
	if g.Adult == nil {
		g.Adult = new(bool)
		*g.Adult = true
	}
}

type GetMovieProductionCompany struct {
	Id            int    `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

type GetMovieProductionCountry struct {
	Iso3166_1 string `json:"iso_3166_1"`
	Name      string `json:"name"`
}

type GetMovieSpokenLanguage struct {
	Iso639_1    string `json:"iso_639_1"`
	Name        string `json:"name"`
	EnglishName string `json:"english_name"`
}

type GetMovieCredits struct {
	Cast []*GetMovieCreditsCast `json:"cast"`
	Crew []*GetMovieCreditsCrew `json:"crew"`
}

type GetMovieCreditsCast struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	Id                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
	CastId             int     `json:"cast_id"`
	Character          string  `json:"character"`
	CreditId           string  `json:"credit_id"`
	Order              int     `json:"order"`
}

type GetMovieCreditsCrew struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	Id                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
	CreditId           string  `json:"credit_id"`
	Department         string  `json:"department"`
	Job                string  `json:"job"`
}

type GetMovieKeywords struct {
	Keywords []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"keywords"`
}

type GetMovieExternalIds struct {
	WikidataId string `json:"wikidata_id"`
}
