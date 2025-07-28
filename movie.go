package tmdb

type Movie struct {
	plan[Object]
}

func NewMovie(o Object) Movie {
	return Movie{plan: rootPlan(o)}
}

func (m Movie) ID() Plan[int32] {
	return int32LeafPlan(m, "id")
}

func (m Movie) Title() Plan[string] {
	return leafPlan[string](m, "title")
}

func (m Movie) Keywords() Slice[Keyword] {
	return Slice[Keyword]{
		plan: leafPlan[Array](m, "keywords"),
	}
}

func (m Movie) ExternalIDs() ExternalIDs {
	return ExternalIDs{
		plan: leafPlan[Object](m, "external_ids"),
	}
}

type Keyword struct {
	plan[Object]
}

func NewKeyword(o Object) Keyword {
	return Keyword{plan: rootPlan(o)}
}

func (k Keyword) Name() Plan[string] {
	return leafPlan[string](k, "name")
}

func (k Keyword) ID() Plan[int32] {
	return int32LeafPlan(k, "id")
}

type ExternalIDs struct {
	plan[Object]
}

func NewExternalIDs(o Object) ExternalIDs {
	return ExternalIDs{plan: rootPlan(o)}
}

func (e ExternalIDs) IMDBID() Plan[string] {
	return leafPlan[string](e, "imdb_id")
}
