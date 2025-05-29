package tmdb_test

import (
	"testing"

	"github.com/krelinga/go-tmdb"
	"github.com/stretchr/testify/assert"
)

func TestGetTvSeries(t *testing.T) {
	tv, err := tmdb.GetTvSeries(getClient(t), 1399)
	if err != nil {
		t.Fatalf("GetTVSeries failed: %v", err)
	}
	if tv == nil {
		t.Fatal("GetTVSeries returned nil")
	}

	assert.False(t, tv.Adult, "TV series should not be marked as adult")
	assert.Equal(t, tmdb.BackdropImage("/zZqpAXxVSBtxV9qPBcscfXBcL2w.jpg"), tv.BackdropImage, "Unexpected backdrop image")
}
