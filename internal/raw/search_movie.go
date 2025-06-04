package raw

type SearchMovie struct {
	Page       int                  `json:"page"`
	Results    []*SearchMovieResult `json:"results"`
	TotalPages int                  `json:"total_pages"`
}

func (s *SearchMovie) SetDefaults() {
	for _, result := range s.Results {
		result.SetDefaults()
	}
}

type SearchMovieResult struct {
	Id int `json:"id"`

	Adult            *bool   `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIds         []int   `json:"genre_ids"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
}

func (s *SearchMovieResult) SetDefaults() {
	if s.Adult == nil {
		s.Adult = new(bool)
		*s.Adult = true
	}
}
