package tmdb_test

import (
	"os"
	"testing"

	"github.com/krelinga/go-tmdb"
)

var globalClient tmdb.Client

func TestMain(m *testing.M) {
	key, ok := os.LookupEnv("TMDB_API_KEY")
	if !ok {
		os.Stderr.WriteString("environment variable TMDB_API_KEY not set\n")
		os.Exit(1)
	}
	globalClient = tmdb.NewClient(key)
	os.Exit(m.Run())
}