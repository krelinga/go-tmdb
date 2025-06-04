package raw

type GetMovie struct {
	*GetMovieDetails

	Credits     *GetMovieCredits     `json:"credits"`
	Keywords    *GetMovieKeywords    `json:"keywords"`
	ExternalIds *GetMovieExternalIds `json:"external_ids"`
}

type GetMovieDetails struct {
	Adult               *bool    `json:"adult"`
	BackdropPath        string   `json:"backdrop_path"`
	BelongsToCollection string   `json:"belongs_to_collection"`
	Budget              int      `json:"budget"`
	Genres              []*Genre `json:"genres"`
	Homepage            string   `json:"homepage"`
	ImdbId              string   `json:"imdb_id"`
	OriginalLanguage    string   `json:"original_language"`
	OriginalTitle       string   `json:"original_title"`
}

func (g *GetMovie) SetDefaults() {
	if g.Adult == nil {
		g.Adult = new(bool)
		*g.Adult = true
	}
}

type GetMovieCredits struct {
	// TODO: Define the structure for movie credits
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
