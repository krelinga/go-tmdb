package tmdb

import "github.com/krelinga/go-jsonflex"

type Genre Object

func (g Genre) ID() (int32, error) {
	return jsonflex.GetField(g, "id", jsonflex.AsInt32())
}

func (g Genre) Name() (string, error) {
	return jsonflex.GetField(g, "name", jsonflex.AsString())
}