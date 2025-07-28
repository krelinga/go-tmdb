package tmdb

import (
	"context"
	"fmt"
)

type Movie struct {
	data[Object]
}

func NewMovie(o Object) Movie {
	return Movie{data: rootData(o)}
}

func GetMovie(ctx context.Context, client Client, id int32, options ...RequestOption) (Movie, error) {
	o, err := client.Get(ctx, fmt.Sprintf("/3/movie/%d", id), options...)
	if err != nil {
		return Movie{}, err
	}
	return NewMovie(o), nil
}

func (m Movie) ID() Data[int32] {
	return int32LeafData(m, "id")
}

func (m Movie) Title() Data[string] {
	return leafData[string](m, "title")
}

func (m Movie) Keywords() Slice[Keyword] {
	return Slice[Keyword]{
		data: leafData[Array](m, "keywords"),
	}
}

func (m Movie) ExternalIDs() ExternalIDs {
	return ExternalIDs{
		data: leafData[Object](m, "external_ids"),
	}
}

type Keyword struct {
	data[Object]
}

func NewKeyword(o Object) Keyword {
	return Keyword{data: rootData(o)}
}

func (k Keyword) Name() Data[string] {
	return leafData[string](k, "name")
}

func (k Keyword) ID() Data[int32] {
	return int32LeafData(k, "id")
}

type ExternalIDs struct {
	data[Object]
}

func NewExternalIDs(o Object) ExternalIDs {
	return ExternalIDs{data: rootData(o)}
}

func (e ExternalIDs) IMDBID() Data[string] {
	return leafData[string](e, "imdb_id")
}
