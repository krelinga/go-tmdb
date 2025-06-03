package raw

type SearchMovie struct {
	Page       int                  `json:"page"`
	Results    []*SearchMovieResult `json:"results"`
	TotalPages int                  `json:"total_pages"`
}

type SearchMovieResult struct {
	Id int `json:"id"`
}
