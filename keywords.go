package tmdb

type KeywordId int

type Keyword struct {
	KeywordId KeywordId `json:"id"`
	Name      string    `json:"name"`
}
