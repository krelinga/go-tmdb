package tmdb

import "github.com/krelinga/go-jsonflex"

type Role Object

func (r Role) CreditID() (string, error) {
	return jsonflex.GetField(r, "credit_id", jsonflex.AsString())
}

func (r Role) Character() (string, error) {
	return jsonflex.GetField(r, "character", jsonflex.AsString())
}

func (r Role) EpisodeCount() (int32, error) {
	return jsonflex.GetField(r, "episode_count", jsonflex.AsInt32())
}
