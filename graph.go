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