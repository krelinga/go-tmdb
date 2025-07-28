package tmdb

type Movie struct {
	data[Object]
}

func NewMovie(o Object) Movie {
	return Movie{data: rootPlan(o)}
}

func (m Movie) ID() Data[int32] {
	return int32LeafPlan(m, "id")
}

func (m Movie) Title() Data[string] {
	return leafPlan[string](m, "title")
}

func (m Movie) Keywords() Slice[Keyword] {
	return Slice[Keyword]{
		data: leafPlan[Array](m, "keywords"),
	}
}

func (m Movie) ExternalIDs() ExternalIDs {
	return ExternalIDs{
		data: leafPlan[Object](m, "external_ids"),
	}
}

type Keyword struct {
	data[Object]
}

func NewKeyword(o Object) Keyword {
	return Keyword{data: rootPlan(o)}
}

func (k Keyword) Name() Data[string] {
	return leafPlan[string](k, "name")
}

func (k Keyword) ID() Data[int32] {
	return int32LeafPlan(k, "id")
}

type ExternalIDs struct {
	data[Object]
}

func NewExternalIDs(o Object) ExternalIDs {
	return ExternalIDs{data: rootPlan(o)}
}

func (e ExternalIDs) IMDBID() Data[string] {
	return leafPlan[string](e, "imdb_id")
}
