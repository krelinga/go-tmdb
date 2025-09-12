package tmdb

import (
	"context"
	"fmt"

	"github.com/krelinga/go-jsonflex"
)

type Collection Object

func (c Collection) ID() (int32, error) {
	return jsonflex.GetField(c, "id", jsonflex.AsInt32())
}

func (c Collection) Name() (string, error) {
	return jsonflex.GetField(c, "name", jsonflex.AsString())
}

func (c Collection) Overview() (string, error) {
	return jsonflex.GetField(c, "overview", jsonflex.AsString())
}

func (c Collection) PosterPath() (string, error) {
	return jsonflex.GetField(c, "poster_path", jsonflex.AsString())
}

func (c Collection) BackdropPath() (string, error) {
	return jsonflex.GetField(c, "backdrop_path", jsonflex.AsString())
}

func GetCollection(ctx context.Context, client Client, collectionID int32, options ...RequestOption) (Collection, error) {
	return client.GetObject(ctx, fmt.Sprintf("/3/collection/%d", collectionID), options...)
}
