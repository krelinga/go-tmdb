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

		s1_again := g.EnsureShow(ShowId(1))
		if s1_again != s1 {
			t.Fatal("Expected to retrieve the same show instance")
		}

		s2 := g.EnsureShow(ShowId(2))
		if s2 == s1 {
			t.Fatal("Expected different show instances for different keys")
		}
	})
}
