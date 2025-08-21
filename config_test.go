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
