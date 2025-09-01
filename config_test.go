package tmdb_test

import (
	"context"
	"slices"
	"testing"

	"github.com/krelinga/go-tmdb"
)

func TestGetConfigDetails(t *testing.T) {
	client := testClientOptions{useApiReadAccessToken: true}.newClient(t)
	config, err := tmdb.GetConfigDetails(context.Background(), client)
	if err != nil {
		t.Fatalf("failed to get config details: %v", err)
	}

	checkChangeKey(t, config, "adult")
	checkChangeKey(t, config, "air_date")
	checkChangeKey(t, config, "also_known_as")
	checkField(t, "http://image.tmdb.org/t/p/", config, tmdb.ConfigDetails.Images, tmdb.ConfigImages.BaseURL)
	checkField(t, "https://image.tmdb.org/t/p/", config, tmdb.ConfigDetails.Images, tmdb.ConfigImages.SecureBaseURL)
	checkImageSize(t, config, tmdb.ConfigImages.BackdropSizes, "original")
	checkImageSize(t, config, tmdb.ConfigImages.LogoSizes, "original")
	checkImageSize(t, config, tmdb.ConfigImages.PosterSizes, "original")
	checkImageSize(t, config, tmdb.ConfigImages.ProfileSizes, "original")
	checkImageSize(t, config, tmdb.ConfigImages.StillSizes, "original")
}

func checkChangeKey(t *testing.T, config tmdb.ConfigDetails, key string) {
	t.Helper()
	changeKeys, err := config.ChangeKeys()
	if err != nil {
		t.Fatalf("failed to get change keys: %v", err)
	}
	if slices.Contains(changeKeys, key) {
		return
	}
	t.Errorf("expected change key %q not found in %v", key, changeKeys)
}

func checkImageSize(t *testing.T, config tmdb.ConfigDetails, sizeType func(tmdb.ConfigImages) ([]string, error), wantSize string) {
	t.Helper()
	if images, err := config.Images(); err != nil {
		t.Fatalf("failed to get images: %v", err)
	} else if sizes, err := sizeType(images); err != nil {
		t.Fatalf("failed to get image sizes: %v", err)
	} else if !slices.Contains(sizes, wantSize) {
		t.Errorf("expected size %q not found in %v", wantSize, sizes)
	}
}

func TestGetConfigCountries(t *testing.T) {
	client := testClientOptions{useApiReadAccessToken: true}.newClient(t)
	countries, err := tmdb.GetConfigCountries(context.Background(), client)
	if err != nil {
		t.Fatalf("failed to get config countries: %v", err)
	}
	if len(countries) == 0 {
		t.Fatal("expected countries, got none")
	} else if usa, err := findCountry(countries, "US"); err != nil {
		t.Fatal(err)
	} else {
		checkField(t, "US", usa, tmdb.Country.ISO3166_1)
		checkField(t, "United States of America", usa, tmdb.Country.EnglishName)
		checkField(t, "United States", usa, tmdb.Country.NativeName)
	}
}

func TestGetConfigJobs(t *testing.T) {
	client := testClientOptions{useApiReadAccessToken: true}.newClient(t)
	jobs, err := tmdb.GetConfigJobs(context.Background(), client)
	if err != nil {
		t.Fatalf("failed to get config jobs: %v", err)
	}
	if len(jobs) == 0 {
		t.Fatal("expected jobs, got none")
	}
	checkConfigJob(t, jobs, "Actors", "Actor")
	checkConfigJob(t, jobs, "Actors", "Stunt Double")
	checkConfigJob(t, jobs, "Actors", "Voice")
	checkConfigJob(t, jobs, "Writing", "Writer")
}

func checkConfigJob(t *testing.T, jobs []tmdb.ConfigJobs, department, job string) {
	t.Helper()
	for _, j := range jobs {
		if dep, err := j.Department(); err != nil {
			t.Fatalf("failed to get department: %v", err)
		} else if dep == department {
			if jobs, err := j.Jobs(); err != nil {
				t.Fatalf("failed to get jobs: %v", err)
			} else if slices.Contains(jobs, job) {
				return
			}
		}
	}
	t.Errorf("expected job %q in department %q not found", job, department)
}

func TestGetConfigLanguages(t *testing.T) {
	client := testClientOptions{useApiReadAccessToken: true}.newClient(t)
	languages, err := tmdb.GetConfigLanguages(context.Background(), client)
	if err != nil {
		t.Fatalf("failed to get config languages: %v", err)
	}
	if len(languages) == 0 {
		t.Fatal("expected languages, got none")
	} else if english, err := findLanguage(languages, "en"); err != nil {
		t.Fatal(err)
	} else {
		checkField(t, "en", english, tmdb.Language.ISO639_1)
		checkField(t, "English", english, tmdb.Language.EnglishName)
		checkField(t, "English", english, tmdb.Language.Name)
	}
}
