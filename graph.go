package tmdb

import "github.com/krelinga/go-views"

type Graph struct {
	companies map[CompanyKey]*Company
	credits map[CreditKey]*Credit
	episodes map[EpisodeKey]*Episode
	genres map[GenreKey]*Genre
	keywords map[KeywordKey]*Keyword
	movies map[MovieKey]*Movie
	networks map[NetworkKey]*Network
	people map[PersonKey]*Person
	seasons map[SeasonKey]*Season
	tvShows map[TvKey]*Tv

	Companies views.Dict[CompanyKey, *Company]
	Credits views.Dict[CreditKey, *Credit]
	Episodes views.Dict[EpisodeKey, *Episode]
	Genres views.Dict[GenreKey, *Genre]
	Keywords views.Dict[KeywordKey, *Keyword]
	Movies views.Dict[MovieKey, *Movie]
	Networks views.Dict[NetworkKey, *Network]
	People views.Dict[PersonKey, *Person]
	Seasons views.Dict[SeasonKey, *Season]
	TvShows views.Dict[TvKey, *Tv]
}

func (g *Graph) init() {
	g.Companies = views.DictOfMap[CompanyKey, *Company]{M: g.companies}
	g.Credits = views.DictOfMap[CreditKey, *Credit]{M: g.credits}
	g.Episodes = views.DictOfMap[EpisodeKey, *Episode]{M: g.episodes}
	g.Genres = views.DictOfMap[GenreKey, *Genre]{M: g.genres}
	g.Keywords = views.DictOfMap[KeywordKey, *Keyword]{M: g.keywords}
	g.Movies = views.DictOfMap[MovieKey, *Movie]{M: g.movies}
	g.Networks = views.DictOfMap[NetworkKey, *Network]{M: g.networks}
	g.People = views.DictOfMap[PersonKey, *Person]{M: g.people}
	g.Seasons = views.DictOfMap[SeasonKey, *Season]{M: g.seasons}
	g.TvShows = views.DictOfMap[TvKey, *Tv]{M: g.tvShows}
}

// I've thought a lot about this and here's how this is going to go:
// For each type of entity there will be a "data" struct that holds the keys along with a lot of optional fields for random data.
// - This has all exported fields and no linkage to other types of entities.
// - This also has a Update() method that will take a pointer to another "data" struct and update the fields of this one.
// There will also be a corresponding "node" struct that provides all the linkage pointers.
// - this will be all methods, no exported fields. (or maybe this will embed the "data" struct?)
// - one getter method for each type of link that returns a bag view.
// - one Add() method for each type of edge, which goes and updates the node at the other end of the edge as well.
// Graph will have several families of methods to update the graph:
// - Upsert() which will take a pointer to the "data" struct and return a pointer to the "node" struct.
// - Rpc() which will use the API to fetch the data and return a pointer to the "node" struct.  Call Upsert() internally.
// - Merge(), which will take pointers to other graph structures to be merged into this one.

// Why not merge "data" and "node" structs?
// - This is nicer for testing & creation syntax.  The "data" struct can be created with a simple literal, 
//   and the "node" struct can be created with a few terse AddEdge() calls.

// for keys, let's keep it simple:
// - each "data" structure has a Key field that is of some type that uniquely identifies the entity.
// - most "Key" fields will be backed by a typed string or int, but some will be backed by a struct named FooKey
// - That way we can still require the nicer syntax around explicit type conversions, but we can avoid a bunch of single-field embedded structs.
