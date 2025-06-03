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

	Adult *bool `json:"adult"`
}

func (s *SearchMovieResult) SetDefaults() {
	if s.Adult == nil {
		s.Adult = new(bool)
		*s.Adult = true
	}
}
