package tmdb

import (
	"context"

	"github.com/krelinga/go-jsonflex"
)

type Genre Object

func (g Genre) ID() (int32, error) {
	return jsonflex.GetField(g, "id", jsonflex.AsInt32())
}

func (g Genre) Name() (string, error) {
	return jsonflex.GetField(g, "name", jsonflex.AsString())
}

type Genres Object

func (g Genres) Genres() ([]Genre, error) {
	return jsonflex.GetField(g, "genres", jsonflex.AsArray(jsonflex.AsObject[Genre]()))
}

func GetMovieGenres(ctx context.Context, client Client, opts ...RequestOption) (Genres, error) {
	return client.GetObject(ctx, "/3/genre/movie/list", opts...)
}

func GetTvGenres(ctx context.Context, client Client, opts ...RequestOption) (Genres, error) {
	return client.GetObject(ctx, "/3/genre/tv/list", opts...)
}
