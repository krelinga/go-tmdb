package tmdb

import "github.com/krelinga/go-jsonflex"

type Season Object

func (s Season) AirDate() (string, error) {
	return jsonflex.GetField(s, "air_date", jsonflex.AsString())
}

func (s Season) EpisodeCount() (int32, error) {
	return jsonflex.GetField(s, "episode_count", jsonflex.AsInt32())
}

func (s Season) ID() (int32, error) {
	return jsonflex.GetField(s, "id", jsonflex.AsInt32())
}

func (s Season) Name() (string, error) {
	return jsonflex.GetField(s, "name", jsonflex.AsString())
}

func (s Season) Overview() (string, error) {
	return jsonflex.GetField(s, "overview", jsonflex.AsString())
}

func (s Season) PosterPath() (string, error) {
	return jsonflex.GetField(s, "poster_path", jsonflex.AsString())
}

func (s Season) SeasonNumber() (int32, error) {
	return jsonflex.GetField(s, "season_number", jsonflex.AsInt32())
}

func (s Season) VoteAverage() (float64, error) {
	return jsonflex.GetField(s, "vote_average", jsonflex.AsFloat64())
}