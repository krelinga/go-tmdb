package raw

type SearchTv struct {
	Page         int               `json:"page"`
	Results      []*SearchTvResult `json:"results"`
	TotalPages   int               `json:"total_pages"`
	TotalResults int               `json:"total_results"`
}

func (s *SearchTv) SetDefaults() {
	for _, result := range s.Results {
		result.SetDefaults()
	}
}

type SearchTvResult struct {
	Id int `json:"id"`

	Adult            *bool    `json:"adult"`
	BackdropPath     string   `json:"backdrop_path"`
	GenreIds         []int    `json:"genre_ids"`
	OriginCountry    []string `json:"origin_country"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       string   `json:"poster_path"`
	FirstAirDate     string   `json:"first_air_date"`
	Name             string   `json:"name"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
}

func (s *SearchTvResult) SetDefaults() {
	if s.Adult == nil {
		s.Adult = new(bool)
		*s.Adult = true
	}
}