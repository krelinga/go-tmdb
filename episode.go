package tmdb

import (
	"context"
	"fmt"

	"github.com/krelinga/go-jsonflex"
)

type Episode Object

func (e Episode) ID() (int32, error) {
	return jsonflex.GetField(e, "id", jsonflex.AsInt32())
}

func (e Episode) Name() (string, error) {
	return jsonflex.GetField(e, "name", jsonflex.AsString())
}

func (e Episode) Overview() (string, error) {
	return jsonflex.GetField(e, "overview", jsonflex.AsString())
}

func (e Episode) VoteAverage() (float64, error) {
	return jsonflex.GetField(e, "vote_average", jsonflex.AsFloat64())
}

func (e Episode) VoteCount() (int32, error) {
	return jsonflex.GetField(e, "vote_count", jsonflex.AsInt32())
}

func (e Episode) AirDate() (string, error) {
	return jsonflex.GetField(e, "air_date", jsonflex.AsString())
}

func (e Episode) EpisodeNumber() (int32, error) {
	return jsonflex.GetField(e, "episode_number", jsonflex.AsInt32())
}

func (e Episode) EpisodeType() (string, error) {
	return jsonflex.GetField(e, "episode_type", jsonflex.AsString())
}

func (e Episode) ProductionCode() (string, error) {
	return jsonflex.GetField(e, "production_code", jsonflex.AsString())
}

func (e Episode) Runtime() (int32, error) {
	return jsonflex.GetField(e, "runtime", jsonflex.AsInt32())
}

func (e Episode) SeasonNumber() (int32, error) {
	return jsonflex.GetField(e, "season_number", jsonflex.AsInt32())
}

func (e Episode) ShowID() (int32, error) {
	return jsonflex.GetField(e, "show_id", jsonflex.AsInt32())
}

func (e Episode) StillPath() (string, error) {
	return jsonflex.GetField(e, "still_path", jsonflex.AsString())
}

func (e Episode) Crew() ([]Credit, error) {
	return jsonflex.GetField(e, "crew", jsonflex.AsArray(jsonflex.AsObject[Credit]()))
}

func (e Episode) GuestStars() ([]Credit, error) {
	return jsonflex.GetField(e, "guest_stars", jsonflex.AsArray(jsonflex.AsObject[Credit]()))
}

func (e Episode) Credits() (Credits, error) {
	return jsonflex.GetField(e, "credits", jsonflex.AsObject[Credits]())
}

func (e Episode) ExternalIDs() (ExternalIDs, error) {
	return jsonflex.GetField(e, "external_ids", jsonflex.AsObject[ExternalIDs]())
}

func GetEpisode(ctx context.Context, client Client, showID int32, seasonNumber int32, episodeNumber int32, opts ...RequestOption) (Episode, error) {
	return client.Get(ctx, fmt.Sprintf("/3/tv/%d/season/%d/episode/%d", showID, seasonNumber, episodeNumber), opts...)
}
