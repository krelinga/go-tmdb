package tmdb

// ISO 639-1 language code.
type LanguageId string

// TODO: do these really need to be pointers?
type Language struct {
	Code *string
	Name *string
	EnglishName *string
}
