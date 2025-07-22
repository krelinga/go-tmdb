package tmdb

import (
	"testing"

	"github.com/krelinga/go-sets"
	"github.com/krelinga/go-views"
)

func getUniqueValues[K, V comparable](t *testing.T, dict views.Dict[K, V]) *sets.Set[V] {
	t.Helper()
	uniqueValues := sets.New[V]()
	for v := range dict.Values() {
		if !uniqueValues.Add(v) {
			t.Fatalf("Expected unique value for %v, but it was not unique", v)
		}
	}
	return uniqueValues
}

func compareValues[T comparable, K comparable](t *testing.T, expected *sets.Set[T], actual views.Dict[K, T], name string) {
	t.Helper()
	uniqueActual := getUniqueValues(t, actual)
	if expected.Len() != uniqueActual.Len() {
		t.Errorf("Expected %s length %d, got %d", name, expected.Len(), uniqueActual.Len())
		return
	}
	for v := range expected.Values() {
		if !uniqueActual.Has(v) {
			t.Errorf("Expected %s to contain item %v, but it was not found", name, v)
		}
	}
	for v := range uniqueActual.Values() {
		if !expected.Has(v) {
			t.Errorf("Expected %s to not contain item %v, but it was found", name, v)
		}
	}
}

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
				uniquePtrs := sets.New[*Show]()
				for _, id := range c.ids {
					uniqueIds.Add(id)
					show := g.EnsureShow(id)
					uniquePtrs.Add(show)
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
				compareValues(t, uniquePtrs, g.Shows(), "shows")
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
				name:       "no seasons, existing show",
				preShowIds: []ShowId{ShowId(1)},
				keys:       []SeasonKey{},
			},
			{
				name:       "one season, existing show",
				preShowIds: []ShowId{ShowId(1)},
				keys:       []SeasonKey{{ShowId: ShowId(1), SeasonNumber: 1}},
			},
			{
				name:       "two different seasons, existing show",
				preShowIds: []ShowId{ShowId(1)},
				keys:       []SeasonKey{{ShowId: ShowId(1), SeasonNumber: 1}, {ShowId: ShowId(1), SeasonNumber: 2}},
			},
			{
				name:       "two same seasons, existing show",
				preShowIds: []ShowId{ShowId(1)},
				keys:       []SeasonKey{{ShowId: ShowId(1), SeasonNumber: 1}, {ShowId: ShowId(1), SeasonNumber: 1}},
			},
		}
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				g := &Graph{}

				if g.Seasons().Len() != 0 {
					t.Error("Expected empty seasons list")
				}

				uniqueShowKeys := sets.New[ShowId]()
				uniqueShowPtrs := sets.New[*Show]()
				for _, id := range c.preShowIds {
					uniqueShowPtrs.Add(g.EnsureShow(id))
					uniqueShowKeys.Add(id)
				}

				uniqueKeys := sets.New[SeasonKey]()
				uniquePtrs := sets.New[*Season]()
				for _, key := range c.keys {
					uniqueKeys.Add(key)
					uniqueShowKeys.Add(key.ShowId)

					season := g.EnsureSeason(key)
					uniquePtrs.Add(season)
					uniqueShowPtrs.Add(season.Show())
					if season == nil {
						t.Fatalf("Expected to get a non-nil season for key %v", key)
					}
					if season.Key != key {
						t.Errorf("Expected season key %v, got %v", key, season.Key)
					}
					if season.Show() != g.EnsureShow(key.ShowId) {
						t.Errorf("Expected season to have show with ID %d, but it has %d", key.ShowId, season.Show().Key)
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

				compareValues(t, uniquePtrs, g.Seasons(), "seasons")
				compareValues(t, uniqueShowPtrs, g.Shows(), "shows")
			})
		}
	})

	t.Run("Episode", func(t *testing.T) {
		cases := []struct {
			name       string
			preShowIds []ShowId
			preSeason  []SeasonKey
			episodes   []EpisodeKey
		}{
			{
				name:       "no episodes, no shows, no seasons",
				preShowIds: []ShowId{},
				preSeason:  []SeasonKey{},
				episodes:   []EpisodeKey{},
			},
			{
				name:       "one episode, novel show, novel season",
				preShowIds: []ShowId{},
				preSeason:  []SeasonKey{},
				episodes:   []EpisodeKey{{ShowId: ShowId(1), SeasonNumber: 1, EpisodeNumber: 1}},
			},
			{
				name:       "two different episodes, novel show, novel season",
				preShowIds: []ShowId{},
				preSeason:  []SeasonKey{},
				episodes:   []EpisodeKey{{ShowId: ShowId(1), SeasonNumber: 1, EpisodeNumber: 1}, {ShowId: ShowId(1), SeasonNumber: 1, EpisodeNumber: 2}},
			},
			{
				name:       "two same episodes, novel show, novel season",
				preShowIds: []ShowId{},
				preSeason:  []SeasonKey{},
				episodes:   []EpisodeKey{{ShowId: ShowId(1), SeasonNumber: 1, EpisodeNumber: 1}, {ShowId: ShowId(1), SeasonNumber: 1, EpisodeNumber: 1}},
			},
			{
				name:       "no episodes, existing show, no seasons",
				preShowIds: []ShowId{ShowId(1)},
				preSeason:  []SeasonKey{},
				episodes:   []EpisodeKey{},
			},
			{
				name:       "one episode, existing show, novel season",
				preShowIds: []ShowId{ShowId(1)},
				preSeason:  []SeasonKey{},
				episodes:   []EpisodeKey{{ShowId: ShowId(1), SeasonNumber: 1, EpisodeNumber: 1}},
			},
			{
				name:       "two different episodes, existing show, novel season",
				preShowIds: []ShowId{ShowId(1)},
				preSeason:  []SeasonKey{},
				episodes:   []EpisodeKey{{ShowId: ShowId(1), SeasonNumber: 1, EpisodeNumber: 1}, {ShowId: ShowId(1), SeasonNumber: 1, EpisodeNumber: 2}},
			},
			{
				name:       "two same episodes, existing show, novel season",
				preShowIds: []ShowId{ShowId(1)},
				preSeason:  []SeasonKey{},
				episodes:   []EpisodeKey{{ShowId: ShowId(1), SeasonNumber: 1, EpisodeNumber: 1}, {ShowId: ShowId(1), SeasonNumber: 1, EpisodeNumber: 1}},
			},
			{
				name:       "no episodes, existing show, existing season",
				preShowIds: []ShowId{ShowId(1)},
				preSeason:  []SeasonKey{{ShowId: ShowId(1), SeasonNumber: 1}},
				episodes:   []EpisodeKey{},
			},
			{
				name:       "one episode, existing show, existing season",
				preShowIds: []ShowId{ShowId(1)},
				preSeason:  []SeasonKey{{ShowId: ShowId(1), SeasonNumber: 1}},
				episodes:   []EpisodeKey{{ShowId: ShowId(1), SeasonNumber: 1, EpisodeNumber: 1}},
			},
			{
				name:       "two different episodes, existing show, existing season",
				preShowIds: []ShowId{ShowId(1)},
				preSeason:  []SeasonKey{{ShowId: ShowId(1), SeasonNumber: 1}},
				episodes:   []EpisodeKey{{ShowId: ShowId(1), SeasonNumber: 1, EpisodeNumber: 1}, {ShowId: ShowId(1), SeasonNumber: 1, EpisodeNumber: 2}},
			},
			{
				name:       "two same episodes, existing show, existing season",
				preShowIds: []ShowId{ShowId(1)},
				preSeason:  []SeasonKey{{ShowId: ShowId(1), SeasonNumber: 1}},
				episodes:   []EpisodeKey{{ShowId: ShowId(1), SeasonNumber: 1, EpisodeNumber: 1}, {ShowId: ShowId(1), SeasonNumber: 1, EpisodeNumber: 1}},
			},
		}
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				g := &Graph{}

				if g.Episodes().Len() != 0 {
					t.Error("Expected empty episodes list")
				}

				uniqueShowKeys := sets.New[ShowId]()
				uniqueShowPtrs := sets.New[*Show]()
				for _, id := range c.preShowIds {
					uniqueShowPtrs.Add(g.EnsureShow(id))
					uniqueShowKeys.Add(id)
				}

				uniqueSeasonKeys := sets.New[SeasonKey]()
				uniqueSeasonPtrs := sets.New[*Season]()
				for _, key := range c.preSeason {
					season := g.EnsureSeason(key)
					uniqueSeasonPtrs.Add(season)
					uniqueSeasonKeys.Add(key)
					uniqueShowPtrs.Add(season.Show())
				}

				uniqueKeys := sets.New[EpisodeKey]()
				uniquePtrs := sets.New[*Episode]()
				for _, key := range c.episodes {
					uniqueKeys.Add(key)
					episode := g.EnsureEpisode(key)
					uniquePtrs.Add(episode)
					uniqueSeasonPtrs.Add(episode.Season())
					uniqueShowPtrs.Add(episode.Show())
					if episode == nil {
						t.Fatalf("Expected to get a non-nil episode for key %v", key)
					}
					if episode.Key != key {
						t.Errorf("Expected episode key %v, got %v", key, episode.Key)
					}
					if episode.Season() != g.EnsureSeason(SeasonKey{ShowId: key.ShowId, SeasonNumber: key.SeasonNumber}) {
						t.Errorf("Expected episode to have season with key %v, but it has %v", SeasonKey{ShowId: key.ShowId, SeasonNumber: key.SeasonNumber}, episode.Season().Key)
					}
					if episode.Show() != g.EnsureShow(key.ShowId) {
						t.Errorf("Expected episode to have show with ID %d, but it has %d", key.ShowId, episode.Show().Key)
					}
					if found, has := g.Episodes().Get(key); !has {
						t.Errorf("Expected to find episode with key %v, but it was not found", key)
					} else if found != episode {
						t.Errorf("Expected found episode to be the same as created episode, but they are different: %v != %v", found, episode)
					}
					if g.Episodes().Len() != uniqueKeys.Len() {
						t.Errorf("Expected episodes length to be %d, got %d", uniqueKeys.Len(), g.Episodes().Len())
					}
				}
				compareValues(t, uniquePtrs, g.Episodes(), "episodes")
				compareValues(t, uniqueSeasonPtrs, g.Seasons(), "seasons")
				compareValues(t, uniqueShowPtrs, g.Shows(), "shows")
			})
		}
	})

	t.Run("Company", func(t *testing.T) {
		cases := []struct {
			name string
			ids  []CompanyId
		}{
			{
				name: "no companies",
				ids:  []CompanyId{},
			},
			{
				name: "one company",
				ids:  []CompanyId{CompanyId(1)},
			},
			{
				name: "two different companies",
				ids:  []CompanyId{CompanyId(1), CompanyId(2)},
			},
			{
				name: "two same companies",
				ids:  []CompanyId{CompanyId(1), CompanyId(1)},
			},
		}
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				g := &Graph{}

				if g.Companies().Len() != 0 {
					t.Error("Expected empty companies list")
				}

				uniqueIds := sets.New[CompanyId]()
				uniquePtrs := sets.New[*Company]()
				for _, id := range c.ids {
					uniqueIds.Add(id)
					company := g.EnsureCompany(id)
					uniquePtrs.Add(company)
					if company == nil {
						t.Fatalf("Expected to get a non-nil company for ID %d", id)
					}
					if company.Key != id {
						t.Errorf("Expected company key %d, got %d", id, company.Key)
					}
					if found, has := g.Companies().Get(id); !has {
						t.Errorf("Expected to find company with key %d, but it was not found", id)
					} else if found != company {
						t.Errorf("Expected found company to be the same as created company, but they are different: %v != %v", found, company)
					}
					if g.Companies().Len() != uniqueIds.Len() {
						t.Errorf("Expected companies length to be %d, got %d", uniqueIds.Len(), g.Companies().Len())
					}
				}
				compareValues(t, uniquePtrs, g.Companies(), "companies")
			})
		}
	})
}
