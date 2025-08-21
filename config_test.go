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
