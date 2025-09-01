package tmdb

import "github.com/krelinga/go-jsonflex"

type Company Object

func (c Company) ID() (int32, error) {
	return jsonflex.GetField(c, "id", jsonflex.AsInt32())
}

func (c Company) Name() (string, error) {
	return jsonflex.GetField(c, "name", jsonflex.AsString())
}

func (c Company) LogoPath() (string, error) {
	return jsonflex.GetField(c, "logo_path", jsonflex.AsString())
}

func (c Company) OriginCountry() (string, error) {
	return jsonflex.GetField(c, "origin_country", jsonflex.AsString())
}
