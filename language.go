package tmdb

// ISO 639-1 language code.
type LanguageId string

type Language interface {
	Id() LanguageId
	Name() string
	EnglishName() string
}
