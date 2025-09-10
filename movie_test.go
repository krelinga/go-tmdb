package tmdb_test

import (
	"context"
	"fmt"
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
	fightClub, err := tmdb.GetMovie(context.Background(), client, 550, tmdb.WithAppendToResponse("credits", "release_dates", "external_ids", "keywords", "images"))
	if err != nil {
		t.Fatalf("failed to get movie: %v", err)
	}
	checkField(t, false, fightClub, tmdb.Movie.Adult)
	checkField(t, int32(63000000), fightClub, tmdb.Movie.Budget)
	checkField(t, "Drama", fightClub, tmdb.Movie.Genres, index(0), tmdb.Genre.Name)
	checkField(t, "http://www.foxmovies.com/movies/fight-club", fightClub, tmdb.Movie.Homepage)
	checkField(t, int32(550), fightClub, tmdb.Movie.ID)
	checkField(t, "tt0137523", fightClub, tmdb.Movie.IMDBID)
	checkField(t, "US", fightClub, tmdb.Movie.OriginCountry, index(0))
	checkField(t, "en", fightClub, tmdb.Movie.OriginalLanguage)
	checkField(t, "Fight Club", fightClub, tmdb.Movie.OriginalTitle)
	checkField(t, "A ticking-time-bomb insomniac and a slippery soap salesman channel primal male aggression into a shocking new form of therapy. Their concept catches on, with underground \"fight clubs\" forming in every town, until an eccentric gets in the way and ignites an out-of-control spiral toward oblivion.", fightClub, tmdb.Movie.Overview)
	checkField(t, float64(24.1745), fightClub, tmdb.Movie.Popularity)
	checkField(t, "/jSziioSwPVrOy9Yow3XhWIBDjq1.jpg", fightClub, tmdb.Movie.PosterPath)
	checkField(t, int32(711), fightClub, tmdb.Movie.ProductionCompanies, index(0), tmdb.Company.ID)
	checkField(t, int32(18), fightClub, tmdb.Movie.Genres, index(0), tmdb.Genre.ID)
	checkField(t, "/tEiIH5QesdheJmDAqQwvtN60727.png", fightClub, tmdb.Movie.ProductionCompanies, index(0), tmdb.Company.LogoPath)
	checkField(t, "Fox 2000 Pictures", fightClub, tmdb.Movie.ProductionCompanies, index(0), tmdb.Company.Name)
	checkField(t, "US", fightClub, tmdb.Movie.ProductionCompanies, index(0), tmdb.Company.OriginCountry)
	if pc, err := fightClub.ProductionCompanies(); err != nil || len(pc) != 5 {
		t.Errorf("expected no error and 5 production companies, got %v and %d", err, len(pc))
	}
	checkField(t, "DE", fightClub, tmdb.Movie.ProductionCountries, index(0), tmdb.Country.ISO3166_1)
	checkField(t, "Germany", fightClub, tmdb.Movie.ProductionCountries, index(0), tmdb.Country.Name)
	if pc, err := fightClub.ProductionCountries(); err != nil || len(pc) != 2 {
		t.Errorf("expected no error and 2 production countries, got %v and %d", err, len(pc))
	}
	checkField(t, "1999-10-15", fightClub, tmdb.Movie.ReleaseDate)
	checkField(t, int32(100853753), fightClub, tmdb.Movie.Revenue)
	checkField(t, int32(139), fightClub, tmdb.Movie.Runtime)
	checkField(t, "English", fightClub, tmdb.Movie.SpokenLanguages, index(0), tmdb.Language.EnglishName)
	checkField(t, "en", fightClub, tmdb.Movie.SpokenLanguages, index(0), tmdb.Language.ISO639_1)
	checkField(t, "English", fightClub, tmdb.Movie.SpokenLanguages, index(0), tmdb.Language.Name)
	checkField(t, "Released", fightClub, tmdb.Movie.Status)
	checkField(t, "Mischief. Mayhem. Soap.", fightClub, tmdb.Movie.Tagline)
	checkField(t, "Fight Club", fightClub, tmdb.Movie.Title)
	checkField(t, false, fightClub, tmdb.Movie.Video)
	checkField(t, float64(8.438), fightClub, tmdb.Movie.VoteAverage)
	checkField(t, int32(30717), fightClub, tmdb.Movie.VoteCount)

	// Credits appended to response.
	checkField(t, false, fightClub, tmdb.Movie.Credits, tmdb.Credits.Cast, index(0), tmdb.Credit.Adult)
	checkField(t, int32(2), fightClub, tmdb.Movie.Credits, tmdb.Credits.Cast, index(0), tmdb.Credit.Gender)
	checkField(t, int32(819), fightClub, tmdb.Movie.Credits, tmdb.Credits.Cast, index(0), tmdb.Credit.ID)
	checkField(t, "Acting", fightClub, tmdb.Movie.Credits, tmdb.Credits.Cast, index(0), tmdb.Credit.KnownForDepartment)
	checkField(t, "Edward Norton", fightClub, tmdb.Movie.Credits, tmdb.Credits.Cast, index(0), tmdb.Credit.Name)
	checkField(t, "Edward Norton", fightClub, tmdb.Movie.Credits, tmdb.Credits.Cast, index(0), tmdb.Credit.OriginalName)
	checkField(t, float64(3.9679), fightClub, tmdb.Movie.Credits, tmdb.Credits.Cast, index(0), tmdb.Credit.Popularity)
	checkField(t, "/8nytsqL59SFJTVYVrN72k6qkGgJ.jpg", fightClub, tmdb.Movie.Credits, tmdb.Credits.Cast, index(0), tmdb.Credit.ProfilePath)
	checkField(t, int32(4), fightClub, tmdb.Movie.Credits, tmdb.Credits.Cast, index(0), tmdb.Credit.CastID)
	checkField(t, "Narrator", fightClub, tmdb.Movie.Credits, tmdb.Credits.Cast, index(0), tmdb.Credit.Character)
	checkField(t, "52fe4250c3a36847f80149f3", fightClub, tmdb.Movie.Credits, tmdb.Credits.Cast, index(0), tmdb.Credit.CreditID)
	checkField(t, int32(0), fightClub, tmdb.Movie.Credits, tmdb.Credits.Cast, index(0), tmdb.Credit.Order)
	if credits, err := fightClub.Credits(); err != nil {
		t.Fatalf("failed to get credits: %v", err)
	} else if cast, err := credits.Cast(); err != nil || len(cast) != 75 {
		t.Errorf("expected no error and 75 cast members, got %v and %d", err, len(cast))
	}
	checkField(t, false, fightClub, tmdb.Movie.Credits, tmdb.Credits.Crew, index(0), tmdb.Credit.Adult)
	checkField(t, int32(2), fightClub, tmdb.Movie.Credits, tmdb.Credits.Crew, index(0), tmdb.Credit.Gender)
	checkField(t, int32(7764), fightClub, tmdb.Movie.Credits, tmdb.Credits.Crew, index(0), tmdb.Credit.ID)
	checkField(t, "Sound", fightClub, tmdb.Movie.Credits, tmdb.Credits.Crew, index(0), tmdb.Credit.KnownForDepartment)
	checkField(t, "Richard Hymns", fightClub, tmdb.Movie.Credits, tmdb.Credits.Crew, index(0), tmdb.Credit.Name)
	checkField(t, "Richard Hymns", fightClub, tmdb.Movie.Credits, tmdb.Credits.Crew, index(0), tmdb.Credit.OriginalName)
	checkField(t, float64(0.3053), fightClub, tmdb.Movie.Credits, tmdb.Credits.Crew, index(0), tmdb.Credit.Popularity)
	checkField(t, "52fe4250c3a36847f8014a41", fightClub, tmdb.Movie.Credits, tmdb.Credits.Crew, index(0), tmdb.Credit.CreditID)
	checkField(t, "Sound", fightClub, tmdb.Movie.Credits, tmdb.Credits.Crew, index(0), tmdb.Credit.Department)
	checkField(t, "Sound Editor", fightClub, tmdb.Movie.Credits, tmdb.Credits.Crew, index(0), tmdb.Credit.Job)
	if credits, err := fightClub.Credits(); err != nil {
		t.Fatalf("failed to get credits: %v", err)
	} else if crew, err := credits.Crew(); err != nil || len(crew) != 188 {
		t.Errorf("expected no error and 188 crew members, got %v and %d", err, len(crew))
	}
	if releaseDates, err := fightClub.ReleaseDates(); err != nil {
		t.Fatalf("failed to get release dates: %v", err)
	} else if results, err := releaseDates.Results(); err != nil || len(results) == 0 {
		t.Errorf("expected no error and some release dates, got %v and %d", err, len(results))
	} else if usReleases, err := findReleseDates(results, "US"); err != nil {
		t.Error(err)
	} else {
		checkField(t, "US", usReleases, tmdb.CountryReleaseDates.ISO3166_1)
		if date, err := findReleaseDate(usReleases, "1999-10-15T00:00:00.000Z"); err != nil {
			t.Error(err)
		} else {
			checkField(t, "R", date, tmdb.ReleaseDate.Certification)
			checkField(t, "1999-10-15T00:00:00.000Z", date, tmdb.ReleaseDate.ReleaseDate)
			checkField(t, int32(3), date, tmdb.ReleaseDate.Type)
		}
		if date, err := findReleaseDate(usReleases, "2000-04-25T00:00:00.000Z"); err != nil {
			t.Error(err)
		} else {
			checkField(t, "R", date, tmdb.ReleaseDate.Certification)
			checkField(t, "2000-04-25T00:00:00.000Z", date, tmdb.ReleaseDate.ReleaseDate)
			checkField(t, "VHS", date, tmdb.ReleaseDate.Note)
			checkField(t, int32(5), date, tmdb.ReleaseDate.Type)
		}
	}
	checkField(t, "tt0137523", fightClub, tmdb.Movie.ExternalIDs, tmdb.ExternalIDs.IMDBID)
	checkField(t, "Q190050", fightClub, tmdb.Movie.ExternalIDs, tmdb.ExternalIDs.WikidataID)
	checkField(t, "FightClub", fightClub, tmdb.Movie.ExternalIDs, tmdb.ExternalIDs.FacebookID)
	if keywords, err := fightClub.Keywords(); err != nil {
		t.Fatalf("failed to get keywords: %v", err)
	} else if keywordList, err := keywords.Keywords(); err != nil {
		t.Fatalf("failed to get keywords: %v", err)
	} else {
		if kw, err := findKeyword(keywordList, "insomnia"); err != nil {
			t.Error(err)
		} else {
			checkField(t, int32(4142), kw, tmdb.Keyword.ID)
			checkField(t, "insomnia", kw, tmdb.Keyword.Name)
		}
		if kw, err := findKeyword(keywordList, "support group"); err != nil {
			t.Error(err)
		} else {
			checkField(t, int32(825), kw, tmdb.Keyword.ID)
			checkField(t, "support group", kw, tmdb.Keyword.Name)
		}
	}
	// Images appended to response.
	if images, err := fightClub.Images(); err != nil {
		t.Fatalf("failed to get images: %v", err)
	} else {
		if backdrops, err := images.Backdrops(); err != nil || len(backdrops) == 0 {
			t.Errorf("expected no error and some backdrops, got %v and %d", err, len(backdrops))
		} else if backdrop, err := findImage(backdrops, "/b9HyPoxwxjxkWEUL5ErZdhApQe2.jpg"); err != nil {
			t.Error(err)
		} else {
			checkField(t, float64(1.778), backdrop, tmdb.Image.AspectRatio)
			checkField(t, int32(1080), backdrop, tmdb.Image.Height)
			checkField(t, "en", backdrop, tmdb.Image.ISO639_1)
			checkField(t, "/b9HyPoxwxjxkWEUL5ErZdhApQe2.jpg", backdrop, tmdb.Image.FilePath)
			checkField(t, float64(3.334), backdrop, tmdb.Image.VoteAverage)
			checkField(t, int32(1), backdrop, tmdb.Image.VoteCount)
			checkField(t, int32(1920), backdrop, tmdb.Image.Width)
		}

		if logos, err := images.Logos(); err != nil || len(logos) == 0 {
			t.Errorf("expected no error and some logos, got %v and %d", err, len(logos))
		} else if logo, err := findImage(logos, "/7Uqhv24pGJs4Ns31NoOPWFJGWNG.png"); err != nil {
			t.Error(err)
		} else {
			checkField(t, float64(4.638), logo, tmdb.Image.AspectRatio)
			checkField(t, int32(389), logo, tmdb.Image.Height)
			checkField(t, "en", logo, tmdb.Image.ISO639_1)
			checkField(t, "/7Uqhv24pGJs4Ns31NoOPWFJGWNG.png", logo, tmdb.Image.FilePath)
			checkField(t, float64(8.034), logo, tmdb.Image.VoteAverage)
			checkField(t, int32(5), logo, tmdb.Image.VoteCount)
			checkField(t, int32(1804), logo, tmdb.Image.Width)
		}

		if posters, err := images.Posters(); err != nil || len(posters) == 0 {
			t.Errorf("expected no error and some posters, got %v and %d", err, len(posters))
		} else if poster, err := findImage(posters, "/r3pPehX4ik8NLYPpbDRAh0YRtMb.jpg"); err != nil {
			t.Error(err)
		} else {
			checkField(t, float64(0.667), poster, tmdb.Image.AspectRatio)
			checkField(t, int32(900), poster, tmdb.Image.Height)
			checkField(t, "pt", poster, tmdb.Image.ISO639_1)
			checkField(t, "/r3pPehX4ik8NLYPpbDRAh0YRtMb.jpg", poster, tmdb.Image.FilePath)
			checkField(t, float64(3.984), poster, tmdb.Image.VoteAverage)
			checkField(t, int32(30), poster, tmdb.Image.VoteCount)
			checkField(t, int32(600), poster, tmdb.Image.Width)
		}
	}
}

func findReleseDates(in []tmdb.CountryReleaseDates, want string) (tmdb.CountryReleaseDates, error) {
	for _, r := range in {
		if iso, err := r.ISO3166_1(); err != nil {
			return tmdb.CountryReleaseDates{}, fmt.Errorf("failed to get ISO3166_1: %w", err)
		} else if iso == want {
			return r, nil
		}
	}
	return tmdb.CountryReleaseDates{}, fmt.Errorf("no release dates found for ISO 3166-1 code: %s", want)
}

func findReleaseDate(in tmdb.CountryReleaseDates, want string) (tmdb.ReleaseDate, error) {
	rd, err := in.ReleaseDates()
	if err != nil {
		return tmdb.ReleaseDate{}, fmt.Errorf("failed to get release dates: %w", err)
	}
	for _, r := range rd {
		if date, err := r.ReleaseDate(); err != nil {
			return tmdb.ReleaseDate{}, fmt.Errorf("failed to get release date: %w", err)
		} else if date == want {
			return r, nil
		}
	}
	return tmdb.ReleaseDate{}, fmt.Errorf("no release date found for date: %s", want)
}

func findKeyword(in []tmdb.Keyword, want string) (tmdb.Keyword, error) {
	for _, k := range in {
		if name, err := k.Name(); err != nil {
			return tmdb.Keyword{}, fmt.Errorf("failed to get keyword name: %w", err)
		} else if name == want {
			return k, nil
		}
	}
	return tmdb.Keyword{}, fmt.Errorf("no keyword found with name: %s", want)
}

func findImage(in []tmdb.Image, want string) (tmdb.Image, error) {
	for _, img := range in {
		if path, err := img.FilePath(); err != nil {
			return tmdb.Image{}, fmt.Errorf("failed to get image file path: %w", err)
		} else if path == want {
			return img, nil
		}
	}
	return tmdb.Image{}, fmt.Errorf("no image found with file path: %s", want)
}