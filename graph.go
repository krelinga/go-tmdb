package tmdb

import "github.com/krelinga/go-views"

type Graph struct {
	companies map[CompanyId]*Company
	credits map[CreditId]*Credit
	episodes map[EpisodeKey]*Episode
	genres map[GenreId]*Genre
	keywords map[KeywordId]*Keyword
	movies map[MovieId]*Movie
	networks map[NetworkId]*Network
	people map[PersonId]*Person
	seasons map[SeasonKey]*Season
	tvShows map[TvId]*Tv

	Companies views.Dict[CompanyId, *Company]
	Credits views.Dict[CreditId, *Credit]
	Episodes views.Dict[EpisodeKey, *Episode]
	Genres views.Dict[GenreId, *Genre]
	Keywords views.Dict[KeywordId, *Keyword]
	Movies views.Dict[MovieId, *Movie]
	Networks views.Dict[NetworkId, *Network]
	People views.Dict[PersonId, *Person]
	Seasons views.Dict[SeasonKey, *Season]
	TvShows views.Dict[TvId, *Tv]
}

func (g *Graph) init() {
	g.Companies = views.DictOfMap[CompanyId, *Company]{M: g.companies}
	g.Credits = views.DictOfMap[CreditId, *Credit]{M: g.credits}
	g.Episodes = views.DictOfMap[EpisodeKey, *Episode]{M: g.episodes}
	g.Genres = views.DictOfMap[GenreId, *Genre]{M: g.genres}
	g.Keywords = views.DictOfMap[KeywordId, *Keyword]{M: g.keywords}
	g.Movies = views.DictOfMap[MovieId, *Movie]{M: g.movies}
	g.Networks = views.DictOfMap[NetworkId, *Network]{M: g.networks}
	g.People = views.DictOfMap[PersonId, *Person]{M: g.people}
	g.Seasons = views.DictOfMap[SeasonKey, *Season]{M: g.seasons}
	g.TvShows = views.DictOfMap[TvId, *Tv]{M: g.tvShows}
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
