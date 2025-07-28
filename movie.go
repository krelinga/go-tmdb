package tmdb

import (
	"context"
	"fmt"
)

type Movie struct {
	data[Object]
}

func NewMovie(in Data[Object]) Movie {
	return Movie{data: in}
}

func GetMovie(ctx context.Context, client Client, id int32, options ...RequestOption) (Movie, error) {
	o, err := client.Get(ctx, fmt.Sprintf("/3/movie/%d", id), options...)
	if err != nil {
		return Movie{}, err
	}
	return NewMovie(NewData(o)), nil
}

func (m Movie) ID() Data[int32] {
	return AsInt32(GetField(m, "id"))
}

func (m Movie) Title() Data[string] {
	return AsString(GetField(m, "title"))
}

func (m Movie) Keywords() Slice[Keyword] {
	sliceAny := NewSlice(AsArray(GetField(m, "keywords")))
	sliceObject := ConvertSlice(sliceAny, AsObject)
	return ConvertSlice(sliceObject, NewKeyword)
}

func (m Movie) ExternalIDs() ExternalIDs {
	return NewExternalIDs(AsObject(GetField(m, "external_ids")))
}

func (m Movie) KeywordIDs() Slice[Data[int32]] {
	sliceAny := NewSlice(AsArray(GetField(m, "keyword_ids")))
	return ConvertSlice(sliceAny, AsInt32)
}

type Keyword struct {
	data[Object]
}

func NewKeyword(in Data[Object]) Keyword {
	return Keyword{data: in}
}

func (k Keyword) Name() Data[string] {
	return AsString(GetField(k, "name"))
}

func (k Keyword) ID() Data[int32] {
	return AsInt32(GetField(k, "id"))
}

type ExternalIDs struct {
	data[Object]
}

func NewExternalIDs(in Data[Object]) ExternalIDs {
	return ExternalIDs{data: in}
}

func (e ExternalIDs) IMDBID() Data[string] {
	return AsString(GetField(e, "imdb_id"))
}

type SearchMoviePage = PageOf[Movie]

func NewSearchMoviePage(o Object) SearchMoviePage {
	return newSearchMoviePage(NewData(o))
}

func newSearchMoviePage(in Data[Object]) SearchMoviePage {
	return NewPaged(in, NewMovie)

}

func SearchMovie(ctx context.Context, client Client, query string, options ...RequestOption) (SearchMoviePage, error) {
	o, err := client.Get(ctx, "/3/search/movie", append(options, WithQueryParam("query", query))...)
	if err != nil {
		return SearchMoviePage{}, err
	}
	return NewSearchMoviePage(o), nil
}
