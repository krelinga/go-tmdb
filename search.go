package tmdb

import (
	"context"

	"github.com/krelinga/go-jsonflex"
)

type SearchResults[T ~Object] Object

func (s SearchResults[T]) Page() (int32, error) {
	return jsonflex.GetField(s, "page", jsonflex.AsInt32())
}

func (s SearchResults[T]) Results() ([]T, error) {
	return jsonflex.GetField(s, "results", jsonflex.AsArray(jsonflex.AsObject[T]()))
}

func (s SearchResults[T]) TotalResults() (int32, error) {
	return jsonflex.GetField(s, "total_results", jsonflex.AsInt32())
}

func (s SearchResults[T]) TotalPages() (int32, error) {
	return jsonflex.GetField(s, "total_pages", jsonflex.AsInt32())
}

func SearchMovie(ctx context.Context, client Client, query string, opts ...RequestOption) (SearchResults[Movie], error) {
	opts = append([]RequestOption{WithQueryParam("query", query)}, opts...)
	return client.GetObject(ctx, "/3/search/movie", opts...)
}

func SearchTv(ctx context.Context, client Client, query string, opts ...RequestOption) (SearchResults[Show], error) {
	opts = append([]RequestOption{WithQueryParam("query", query)}, opts...)
	return client.GetObject(ctx, "/3/search/tv", opts...)
}
