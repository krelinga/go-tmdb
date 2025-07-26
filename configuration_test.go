package tmdb_test

import (
	"testing"

	"github.com/krelinga/go-tmdb"
)

func TestGetConfiguration(t *testing.T) {
	var raw []byte
	config, err := tmdb.GetConfiguration(getClient(t), tmdb.WithRawReply(&raw))
	if err != nil {
		t.Fatalf("GetConfiguration failed: %v", err)
	}
	if config == nil {
		t.Fatal("GetConfiguration returned nil")
	}
	if len(config.Images.BackdropSizes) == 0 {
		t.Fatal("No backdrop sizes found")
	}
	if len(config.Images.LogoSizes) == 0 {
		t.Fatal("No logo sizes found")
	}
	if len(config.Images.PosterSizes) == 0 {
		t.Fatal("No poster sizes found")
	}
	if len(config.Images.StillSizes) == 0 {
		t.Fatal("No still sizes found")
	}
	t.Log(string(raw))
}
