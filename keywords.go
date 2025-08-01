package tmdb

import "github.com/krelinga/go-jsonflex"

type Keywords Object

func (k Keywords) ID() (int32, error) {
	return jsonflex.GetField(k, "id", jsonflex.AsInt32())
}

func (k Keywords) Keywords() ([]Keyword, error) {
	return jsonflex.GetField(k, "keywords", jsonflex.AsArray(jsonflex.AsObject[Keyword]()))
}

type Keyword Object

func (k Keyword) ID() (int32, error) {
	return jsonflex.GetField(k, "id", jsonflex.AsInt32())
}

func (k Keyword) Name() (string, error) {
	return jsonflex.GetField(k, "name", jsonflex.AsString())
}
