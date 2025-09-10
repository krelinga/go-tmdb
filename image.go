package tmdb

import "github.com/krelinga/go-jsonflex"

type Image jsonflex.Object

func (i Image) AspectRatio() (float64, error) {
	return jsonflex.GetField(i, "aspect_ratio", jsonflex.AsFloat64())
}

func (i Image) Height() (int32, error) {
	return jsonflex.GetField(i, "height", jsonflex.AsInt32())
}

func (i Image) ISO639_1() (string, error) {
	return jsonflex.GetField(i, "iso_639_1", jsonflex.AsString())
}

func (i Image) FilePath() (string, error) {
	return jsonflex.GetField(i, "file_path", jsonflex.AsString())
}

func (i Image) VoteAverage() (float64, error) {
	return jsonflex.GetField(i, "vote_average", jsonflex.AsFloat64())
}

func (i Image) VoteCount() (int32, error) {
	return jsonflex.GetField(i, "vote_count", jsonflex.AsInt32())
}

func (i Image) Width() (int32, error) {
	return jsonflex.GetField(i, "width", jsonflex.AsInt32())
}

type Images jsonflex.Object

func (i Images) Backdrops() ([]Image, error) {
	return jsonflex.GetField(i, "backdrops", jsonflex.AsArray(jsonflex.AsObject[Image]()))
}

func (i Images) Logos() ([]Image, error) {
	return jsonflex.GetField(i, "logos", jsonflex.AsArray(jsonflex.AsObject[Image]()))
}

func (i Images) Posters() ([]Image, error) {
	return jsonflex.GetField(i, "posters", jsonflex.AsArray(jsonflex.AsObject[Image]()))
}