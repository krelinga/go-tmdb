package tmdb

import "github.com/krelinga/go-jsonflex"

type Language Object

func (l Language) EnglishName() (string, error) {
	return jsonflex.GetField(l, "english_name", jsonflex.AsString())
}

func (l Language) ISO639_1() (string, error) {
	return jsonflex.GetField(l, "iso_639_1", jsonflex.AsString())
}

func (l Language) Name() (string, error) {
	return jsonflex.GetField(l, "name", jsonflex.AsString())
}