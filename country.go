package tmdb

import "github.com/krelinga/go-jsonflex"

type Country Object

func (c Country) ISO3166_1() (string, error) {
	return jsonflex.GetField(c, "iso_3166_1", jsonflex.AsString())
}

func (c Country) Name() (string, error) {
	return jsonflex.GetField(c, "name", jsonflex.AsString())
}

func (c Country) EnglishName() (string, error) {
	return jsonflex.GetField(c, "english_name", jsonflex.AsString())
}

func (c Country) NativeName() (string, error) {
	return jsonflex.GetField(c, "native_name", jsonflex.AsString())
}
