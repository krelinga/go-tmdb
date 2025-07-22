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
		cases := []struct {
			name       string
			preShowIds []ShowId
			keys       []SeasonKey
		}{
			{
				name: "no seasons, no shows",
				keys: []SeasonKey{},
			},
			{
				name: "one season, novel show",
				keys: []SeasonKey{{ShowId: ShowId(1), SeasonNumber: 1}},
			},
			{
				name: "two different seasons, novel show",
				keys: []SeasonKey{{ShowId: ShowId(1), SeasonNumber: 1}, {ShowId: ShowId(1), SeasonNumber: 2}},
			},
			{
				name: "two same seasons, novel show",
				keys: []SeasonKey{{ShowId: ShowId(1), SeasonNumber: 1}, {ShowId: ShowId(1), SeasonNumber: 1}},
			},
			{
				name: "no seasons, existing show",
				preShowIds: []ShowId{ShowId(1)},
				keys: []SeasonKey{},
			},
			{
				name: "one season, existing show",
				preShowIds: []ShowId{ShowId(1)},
				keys: []SeasonKey{{ShowId: ShowId(1), SeasonNumber: 1}},
			},
			{
				name: "two different seasons, existing show",
				preShowIds: []ShowId{ShowId(1)},
				keys: []SeasonKey{{ShowId: ShowId(1), SeasonNumber: 1}, {ShowId: ShowId(1), SeasonNumber: 2}},
			},
			{
				name: "two same seasons, existing show",
				preShowIds: []ShowId{ShowId(1)},
				keys: []SeasonKey{{ShowId: ShowId(1), SeasonNumber: 1}, {ShowId: ShowId(1), SeasonNumber: 1}},
			},
		}
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				g := &Graph{}

				if g.Seasons().Len() != 0 {
					t.Error("Expected empty seasons list")
				}

				uniqueShowKeys := sets.New[ShowId]()
				for _, id := range c.preShowIds {
					g.EnsureShow(id)
					uniqueShowKeys.Add(id)
				}

				uniqueKeys := sets.New[SeasonKey]()
				for _, key := range c.keys {
					uniqueKeys.Add(key)
					uniqueShowKeys.Add(key.ShowId)

					season := g.EnsureSeason(key)
					if season == nil {
						t.Fatalf("Expected to get a non-nil season for key %v", key)
					}
					if season.Key != key {
						t.Errorf("Expected season key %v, got %v", key, season.Key)
					}
					if found, has := g.Seasons().Get(key); !has {
						t.Errorf("Expected to find season with key %v, but it was not found", key)
					} else if found != season {
						t.Errorf("Expected found season to be the same as created season, but they are different: %v != %v", found, season)
					}
					if g.Seasons().Len() != uniqueKeys.Len() {
						t.Errorf("Expected seasons length to be %d, got %d", uniqueKeys.Len(), g.Seasons().Len())
					}
					if g.Shows().Len() != uniqueShowKeys.Len() {
						t.Errorf("Expected shows length to be %d, got %d", uniqueShowKeys.Len(), g.Shows().Len())
					}
				}
			})
		}
	})
}
