package tmdb

import "testing"

func TestGraph(t *testing.T) {
	t.Run("Show", func(t *testing.T) {
		g := &Graph{}

		if g.Shows().Len() != 0 {
			t.Error("Expected empty shows list")
		}

		s1 := g.EnsureShow(ShowId(1))
		if s1 == nil {
			t.Fatal("Expected to create a new show")
		}
		if s1.Key != ShowId(1) {
			t.Errorf("Expected show key 1, got %d", s1.Key)
		}

		s1Again := g.EnsureShow(ShowId(1))
		if s1Again != s1 {
			t.Fatal("Expected to retrieve the same show instance")
		}

		s2 := g.EnsureShow(ShowId(2))
		if s2 == s1 {
			t.Fatal("Expected different show instances for different keys")
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
