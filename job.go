package tmdb

import "github.com/krelinga/go-jsonflex"

type Job Object

func (j Job) CreditID() (string, error) {
	return jsonflex.GetField(j, "credit_id", jsonflex.AsString())
}

func (j Job) Job() (string, error) {
	return jsonflex.GetField(j, "job", jsonflex.AsString())
}

func (j Job) EpisodeCount() (int32, error) {
	return jsonflex.GetField(j, "episode_count", jsonflex.AsInt32())
}
