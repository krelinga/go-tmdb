package tmdb

type LanguageIso639_1 string

type Language struct {
	EnglishName string           `json:"english_name"`
	Iso639_1    LanguageIso639_1 `json:"iso_639_1"`
	Name        string           `json:"name"`
}
