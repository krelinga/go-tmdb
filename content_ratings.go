package tmdb

import "github.com/krelinga/go-jsonflex"

type ContentRatings Object

func (cr ContentRatings) Results() ([]ContentRating, error) {
	return jsonflex.GetField(cr, "results", jsonflex.AsArray(jsonflex.AsObject[ContentRating]()))
}

func (cr ContentRatings) ID() (int32, error) {
	return jsonflex.GetField(cr, "id", jsonflex.AsInt32())
}

type ContentRating Object

func (cr ContentRating) ISO3166_1() (string, error) {
	return jsonflex.GetField(cr, "iso_3166_1", jsonflex.AsString())
}

func (cr ContentRating) Rating() (string, error) {
	return jsonflex.GetField(cr, "rating", jsonflex.AsString())
}
