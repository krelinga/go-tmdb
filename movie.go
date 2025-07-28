package tmdb

import (
	"context"
	"fmt"
)

type Movie struct {
	data[Object]
}

func NewMovie(o Object) Movie {
	return newMovie(NewData(o))
}

func newMovie(in Data[Object]) Movie {
	return Movie{data: in}
}

func GetMovie(ctx context.Context, client Client, id int32, options ...RequestOption) (Movie, error) {
	o, err := client.Get(ctx, fmt.Sprintf("/3/movie/%d", id), options...)
	if err != nil {
		return Movie{}, err
	}
	return NewMovie(o), nil
}

func (m Movie) ID() Data[int32] {
	return fieldData(m, "id", asInt32)
}

func (m Movie) Title() Data[string] {
	return fieldData(m, "title", asString)
}

func (m Movie) Keywords() Slice[Keyword] {
	return newSlice(fieldData(m, "keywords", asArray), asTypedObject(NewKeyword))
}

func (m Movie) ExternalIDs() ExternalIDs {
	return newExternalIDs(fieldData(m, "external_ids", asObject))
}

func (m Movie) KeywordIDs() Slice[int32] {
	return newSlice(fieldData(m, "keyword_ids", asArray), asInt32)
}

type Keyword struct {
	data[Object]
}

func NewKeyword(o Object) Keyword {
	return Keyword{data: NewData(o)}
}

func (k Keyword) Name() Data[string] {
	return fieldData(k, "name", asString)
}

func (k Keyword) ID() Data[int32] {
	return fieldData(k, "id", asInt32)
}

type ExternalIDs struct {
	data[Object]
}

func NewExternalIDs(o Object) ExternalIDs {
	return ExternalIDs{data: NewData(o)}
}

func newExternalIDs(in Data[Object]) ExternalIDs {
	return ExternalIDs{data: in}
}

func (e ExternalIDs) IMDBID() Data[string] {
	return fieldData(e, "imdb_id", asString)
}

type SearchMoviePage = PageOf[Movie]

func NewSearchMoviePage(o Object) SearchMoviePage {
	return newSearchMoviePage(NewData(o))
}

func newSearchMoviePage(in Data[Object]) SearchMoviePage {
	return newPaged(in, asTypedObject(NewMovie))
}

func SearchMovie(ctx context.Context, client Client, query string, options ...RequestOption) (SearchMoviePage, error) {
	o, err := client.Get(ctx, "/3/search/movie", append(options, WithQueryParam("query", query))...)
	if err != nil {
		return SearchMoviePage{}, err
	}
	return NewSearchMoviePage(o), nil
}
