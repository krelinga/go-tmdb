package tmdb

import (
	"testing"

	"github.com/krelinga/go-sets"
)

func TestGraph(t *testing.T) {
	t.Run("Show", func(t *testing.T) {
		cases := []struct {
			name string
			ids  []ShowId
		}{
			{
				name: "no shows",
				ids:  []ShowId{},
			},
			{
				name: "one show",
				ids:  []ShowId{ShowId(1)},
			},
			{
				name: "two different shows",
				ids:  []ShowId{ShowId(1), ShowId(2)},
			},
			{
				name: "two same shows",
				ids:  []ShowId{ShowId(1), ShowId(1)},
			},
		}
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				g := &Graph{}

				if g.Shows().Len() != 0 {
					t.Error("Expected empty shows list")
				}

				uniqueIds := sets.New[ShowId]()
				for _, id := range c.ids {
					uniqueIds.Add(id)
					show := g.EnsureShow(id)
					if show == nil {
						t.Fatalf("Expected to get a non-nil show for ID %d", id)
					}
					if show.Key != id {
						t.Errorf("Expected show key %d, got %d", id, show.Key)
					}
					if found, has := g.Shows().Get(id); !has {
						t.Errorf("Expected to find show with key %d, but it was not found", id)
					} else if found != show {
						t.Errorf("Expected found show to be the same as created show, but they are different: %v != %v", found, show)
					}
					if g.Shows().Len() != uniqueIds.Len() {
						t.Errorf("Expected shows length to be %d, got %d", uniqueIds.Len(), g.Shows().Len())
					}
				}
			})
		}
	})

	t.Run("Season", func(t *testing.T) {
		g := &Graph{}

		if g.Seasons().Len() != 0 {
			t.Error("Expected empty seasons list")
		}

		season1Key := SeasonKey{ShowId: ShowId(1), SeasonNumber: 1}
		season1 := g.EnsureSeason(season1Key)
		if season1 == nil {
			t.Fatal("Expected to create a new season")
		}
		if season1.Key.ShowId != ShowId(1) || season1.Key.SeasonNumber != 1 {
			t.Errorf("Expected season key {ShowId: 1, SeasonNumber: 1}, got {%d, %d}", season1.Key.ShowId, season1.Key.SeasonNumber)
		}
		if found, has := g.Seasons().Get(season1Key); !has {
			t.Errorf("Expected to find season with key %v, but it was not found", season1Key)
		} else if found != season1 {
			t.Errorf("Expected found season to be the same as created season, but they are different")
		}

		season1Again := g.EnsureSeason(season1Key)
		if season1Again != season1 {
			t.Fatal("Expected to retrieve the same season instance")
		}

		season2Key := SeasonKey{ShowId: ShowId(2), SeasonNumber: 1}
		season2 := g.EnsureSeason(season2Key)
		if season2 == season1 {
			t.Fatal("Expected different season instances for different keys")
		}
		if found, has := g.Seasons().Get(season2Key); !has {
			t.Errorf("Expected to find season with key %v, but it was not found", season2Key)
		} else if found != season2 {
			t.Errorf("Expected found season to be the same as created season, but they are different")
		}
	})
}
