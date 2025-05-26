package tmdb_test

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/krelinga/go-tmdb"
)

var replayFlag = flag.String("replay", "read", "Which mode to run the replay client in.  Options are: 'direct', 'read', 'write', or 'replace'")

var clientOnce = sync.OnceValues(func() (tmdb.Client, error) {
	var direct, write, replace bool
	switch *replayFlag {
	case "direct":
		direct = true
	case "read":
		// Nothing to do.
	case "write":
		write = true
	case "replace":
		replace = true
	default:
		return nil, errors.New("Invalid replay mode. Options are: 'direct', 'read', 'write', or 'replace'")
	}

	const dataDir = "testdata"

	if replace {
		// Remove all the existing files.
		files, err := os.ReadDir(dataDir)
		if err != nil && !os.IsNotExist(err) {
			return nil, fmt.Errorf("Failed to read test data directory: %w", err)
		}
		for _, file := range files {
			if file.IsDir() || !strings.HasSuffix(file.Name(), ".json") {
				continue
			}
			err := os.Remove(dataDir + "/" + file.Name())
			if err != nil {
				return nil, fmt.Errorf("Failed to remove file %s: %w", file.Name(), err)
			}
		}
	}

	var upstream tmdb.Client
	if direct || write || replace {
		key, ok := os.LookupEnv("TMDB_API_KEY")
		if !ok {
			return nil, errors.New("environment variable TMDB_API_KEY not set")
		}
		upstream = tmdb.NewClient(key)
	}
	var client tmdb.Client
	if direct {
		client = upstream
	} else {
		var err error
		client, err = tmdb.NewReplayClient(upstream, dataDir)
		if err != nil {
			return nil, fmt.Errorf("Failed to create replay client: %w", err)
		}
	}
	return client, nil
})

func getClient(t *testing.T) tmdb.Client {
	t.Helper()
	client, err := clientOnce()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	} else {
		t.Logf("Using client: %s", client)
	}
	return client
}
