package tmdb_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/krelinga/go-tmdb"
)

type index int

func checkField[ObjType ~tmdb.Object, ValueType comparable](t *testing.T, want ValueType, obj ObjType, calls ...any) {
	t.Helper()
	v := reflect.ValueOf(obj)
	for i, call := range calls {
		callV := reflect.ValueOf(call)
		if callV.Type() == reflect.TypeFor[index]() {
			// Handle this as an array index.
			if v.Kind() != reflect.Slice {
				t.Fatalf("expected a slice, got %s at function %d", v.Kind(), i)
			}
			wantIndex := int(callV.Int())
			if wantIndex < 0 || wantIndex >= v.Len() {
				t.Fatalf("index %d out of bounds for slice of length %d at function %d", wantIndex, v.Len(), i)
			}
			v = v.Index(wantIndex)
		} else {
			// Handle this as a method call.
			if callV.Kind() != reflect.Func {
				t.Fatalf("expected a function, got %s at function %d", callV.Kind(), i)
			}
			gotValues := callV.Call([]reflect.Value{v})
			if len(gotValues) != 2 {
				t.Fatalf("expected function %d to return 2 values, got %d", i, len(gotValues))
			}
			if !gotValues[1].IsNil() {
				err, ok := gotValues[1].Interface().(error)
				if !ok {
					t.Fatalf("expected second return value of function %d to be an error, got %T", i, gotValues[1].Interface())
				}
				if err != nil {
					t.Fatalf("function %d returned an error: %v", i, err)
				}
			}
			v = gotValues[0]
		}
	}
	got, ok := v.Interface().(ValueType)
	if !ok {
		t.Fatalf("expected field to be of type %T, got %T", want, v.Interface())
	}
	if got != want {
		t.Errorf("field: got %v, want %v", got, want)
	}
}

func TestGetMovie(t *testing.T) {
	client := testClientOptions{useApiReadAccessToken: true}.newClient(t)
	fightClub, err := tmdb.GetMovie(context.Background(), client, 550)
	if err != nil {
		t.Fatalf("failed to get movie: %v", err)
	}
	checkField(t, false, fightClub, tmdb.Movie.Adult)
	checkField(t, int32(63000000), fightClub, tmdb.Movie.Budget)
	checkField(t, "http://www.foxmovies.com/movies/fight-club", fightClub, tmdb.Movie.Homepage)
	checkField(t, int32(550), fightClub, tmdb.Movie.ID)
	checkField(t, "tt0137523", fightClub, tmdb.Movie.IMDBID)
	checkField(t, "US", fightClub, tmdb.Movie.OriginCountry, index(0))
	checkField(t, "en", fightClub, tmdb.Movie.OriginalLanguage)
	checkField(t, "Fight Club", fightClub, tmdb.Movie.OriginalTitle)
	checkField(t, "A ticking-time-bomb insomniac and a slippery soap salesman channel primal male aggression into a shocking new form of therapy. Their concept catches on, with underground \"fight clubs\" forming in every town, until an eccentric gets in the way and ignites an out-of-control spiral toward oblivion.", fightClub, tmdb.Movie.Overview)
	checkField(t, float64(23.1752), fightClub, tmdb.Movie.Popularity)
	checkField(t, "/jSziioSwPVrOy9Yow3XhWIBDjq1.jpg", fightClub, tmdb.Movie.PosterPath)
	checkField(t, int32(711), fightClub, tmdb.Movie.ProductionCompanies, index(0), tmdb.Company.ID)
	checkField(t, int32(18), fightClub, tmdb.Movie.Genres, index(0), tmdb.Genre.ID)
	checkField(t, "/tEiIH5QesdheJmDAqQwvtN60727.png", fightClub, tmdb.Movie.ProductionCompanies, index(0), tmdb.Company.LogoPath)
	checkField(t, "Fox 2000 Pictures", fightClub, tmdb.Movie.ProductionCompanies, index(0), tmdb.Company.Name)
	checkField(t, "US", fightClub, tmdb.Movie.ProductionCompanies, index(0), tmdb.Company.OriginCountry)
	if pc, err := fightClub.ProductionCompanies(); err != nil || len(pc) != 5 {
		t.Errorf("expected no error and 5 production companies, got %v and %d", err, len(pc))
	}
	checkField(t, "Drama", fightClub, tmdb.Movie.Genres, index(0), tmdb.Genre.Name)
	// TODO: write more tests for other fields.
}
