package tmdb

import "github.com/krelinga/go-views"

type GenreId int

type Genre struct {
	Key  GenreId
	Data GenreData

	movies []*Movie
}

func (g *Genre) AddMovie(m *Movie) {
	m.AddGenre(g)
}

func (g *Genre) Movies() views.Bag[*Movie] {
	return views.BagOfSlice[*Movie]{S: g.movies}
}

type GenreData struct {
	Name *string
}
