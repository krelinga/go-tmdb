package tmdb

import (
	"context"
	"fmt"

	"github.com/krelinga/go-jsonflex"
)

type Season Object

func (s Season) UnderbarID() (string, error) {
	return jsonflex.GetField(s, "_id", jsonflex.AsString())
}

func (s Season) AirDate() (string, error) {
	return jsonflex.GetField(s, "air_date", jsonflex.AsString())
}

func (s Season) Episodes() ([]Episode, error) {
	return jsonflex.GetField(s, "episodes", jsonflex.AsArray(jsonflex.AsObject[Episode]()))
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

func (s Season) ShowID() (int32, error) {
	return jsonflex.GetField(s, "show_id", jsonflex.AsInt32())
}

func GetSeason(ctx context.Context, client Client, showID, seasonNumber int32, opts ...RequestOption) (Season, error) {
	return client.GetObject(ctx, fmt.Sprintf("/3/tv/%d/season/%d", showID, seasonNumber), opts...)
}
