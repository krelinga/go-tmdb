package tmdb

import "github.com/krelinga/go-views"

type Graph struct {
	companies map[CompanyId]*Company
	credits   map[CreditId]*Credit
	episodes  map[EpisodeKey]*Episode
	genres    map[GenreId]*Genre
	keywords  map[KeywordId]*Keyword
	movies    map[MovieId]*Movie
	networks  map[NetworkId]*Network
	people    map[PersonId]*Person
	seasons   map[SeasonKey]*Season
	tvShows   map[ShowId]*Show
}

func ensureHelper[K comparable, V any](m *map[K]*V, key K, newfn func() *V) *V {
	if *m == nil {
		*m = make(map[K]*V)
	}
	if value, ok := (*m)[key]; ok {
		return value
	}
	value := newfn()
	(*m)[key] = value
	return value
}

func (g *Graph) EnsureCompany(id CompanyId) *Company {
	return ensureHelper(&g.companies, id, func() *Company { return &Company{Key: id} })
}

func (g *Graph) Companies() views.Dict[CompanyId, *Company] {
	return views.DictOfMap[CompanyId, *Company]{M: g.companies}
}

func (g *Graph) EnsureCredit(id CreditId) *Credit {
	return ensureHelper(&g.credits, id, func() *Credit { return &Credit{Key: id} })
}

func (g *Graph) Credits() views.Dict[CreditId, *Credit] {
	return views.DictOfMap[CreditId, *Credit]{M: g.credits}
}

func (g *Graph) EnsureEpisode(key EpisodeKey) *Episode {
	return ensureHelper(&g.episodes, key, func() *Episode { return &Episode{Key: key} })
}

func (g *Graph) Episodes() views.Dict[EpisodeKey, *Episode] {
	return views.DictOfMap[EpisodeKey, *Episode]{M: g.episodes}
}

func (g *Graph) EnsureGenre(id GenreId) *Genre {
	return ensureHelper(&g.genres, id, func() *Genre { return &Genre{Key: id} })
}

func (g *Graph) Genres() views.Dict[GenreId, *Genre] {
	return views.DictOfMap[GenreId, *Genre]{M: g.genres}
}

func (g *Graph) EnsureKeyword(id KeywordId) *Keyword {
	return ensureHelper(&g.keywords, id, func() *Keyword { return &Keyword{Key: id} })
}

func (g *Graph) Keywords() views.Dict[KeywordId, *Keyword] {
	return views.DictOfMap[KeywordId, *Keyword]{M: g.keywords}
}

func (g *Graph) EnsureMovie(id MovieId) *Movie {
	return ensureHelper(&g.movies, id, func() *Movie { return &Movie{Key: id} })
}

func (g *Graph) Movies() views.Dict[MovieId, *Movie] {
	return views.DictOfMap[MovieId, *Movie]{M: g.movies}
}

func (g *Graph) EnsureNetwork(id NetworkId) *Network {
	return ensureHelper(&g.networks, id, func() *Network { return &Network{Key: id} })
}

func (g *Graph) Networks() views.Dict[NetworkId, *Network] {
	return views.DictOfMap[NetworkId, *Network]{M: g.networks}
}

func (g *Graph) EnsurePerson(id PersonId) *Person {
	return ensureHelper(&g.people, id, func() *Person { return &Person{Key: id} })
}

func (g *Graph) People() views.Dict[PersonId, *Person] {
	return views.DictOfMap[PersonId, *Person]{M: g.people}
}

func (g *Graph) EnsureSeason(key SeasonKey) *Season {
	return ensureHelper(&g.seasons, key, func() *Season { return &Season{Key: key} })
}

func (g *Graph) Seasons() views.Dict[SeasonKey, *Season] {
	return views.DictOfMap[SeasonKey, *Season]{M: g.seasons}
}

func (g *Graph) EnsureShow(id ShowId) *Show {
	return ensureHelper(&g.tvShows, id, func() *Show { return &Show{Key: id} })
}

func (g *Graph) Shows() views.Dict[ShowId, *Show] {
	return views.DictOfMap[ShowId, *Show]{M: g.tvShows}
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
