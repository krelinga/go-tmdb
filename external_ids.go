package tmdb

import (
	"context"
	"fmt"

	"github.com/krelinga/go-jsonflex"
)

type ExternalIDs Object

func (e ExternalIDs) ID() (int32, error) {
	return jsonflex.GetField(e, "id", jsonflex.AsInt32())
}

func (e ExternalIDs) IMDBID() (string, error) {
	return jsonflex.GetField(e, "imdb_id", jsonflex.AsString())
}

func (e ExternalIDs) FreebaseMID() (string, error) {
	return jsonflex.GetField(e, "freebase_mid", jsonflex.AsString())
}

func (e ExternalIDs) FreebaseID() (string, error) {
	return jsonflex.GetField(e, "freebase_id", jsonflex.AsString())
}

func (e ExternalIDs) TVDBID() (int32, error) {
	return jsonflex.GetField(e, "tvdb_id", jsonflex.AsInt32())
}

func (e ExternalIDs) TVRageID() (int32, error) {
	return jsonflex.GetField(e, "tvrage_id", jsonflex.AsInt32())
}

func (e ExternalIDs) FacebookID() (string, error) {
	return jsonflex.GetField(e, "facebook_id", jsonflex.AsString())
}

func (e ExternalIDs) InstagramID() (string, error) {
	return jsonflex.GetField(e, "instagram_id", jsonflex.AsString())
}

func (e ExternalIDs) TwitterID() (string, error) {
	return jsonflex.GetField(e, "twitter_id", jsonflex.AsString())
}

func (e ExternalIDs) WikidataID() (string, error) {
	return jsonflex.GetField(e, "wikidata_id", jsonflex.AsString())
}

func GetMovieExternalIDs(ctx context.Context, client Client, movieID int32, opts ...RequestOption) (ExternalIDs, error) {
	return client.Get(ctx, fmt.Sprintf("/3/movie/%d/external_ids", movieID), opts...)
}