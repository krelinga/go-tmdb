package tmdb_test

import (
	"flag"
	"os"
	"strings"
	"testing"

	"github.com/krelinga/go-tmdb"
)

var globalClient tmdb.Client
var replayFlag = flag.String("replay", "read", "Which mode to run the replay client in.  Options are: 'direct', 'read', 'write', or 'replace'")

func TestMain(m *testing.M) {
	flag.Parse()

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
		os.Stderr.WriteString("Invalid replay mode. Options are: 'direct', 'read', 'write', or 'replace'\n")
		os.Exit(1)
	}

	const dataDir = "testdata"

	if replace {
		// Remove all the existing files.
		files, err := os.ReadDir(dataDir)
		if err != nil && !os.IsNotExist(err) {
			os.Stderr.WriteString("Failed to read test data directory: " + err.Error() + "\n")
			os.Exit(1)
		}
		for _, file := range files {
			if file.IsDir() || !strings.HasSuffix(file.Name(), ".json") {
				continue
			}
			err := os.Remove(dataDir + "/" + file.Name())
			if err != nil {
				os.Stderr.WriteString("Failed to remove file " + file.Name() + ": " + err.Error() + "\n")
				os.Exit(1)
			}
		}
	}

	var upstream tmdb.Client
	if direct || write || replace {
		key, ok := os.LookupEnv("TMDB_API_KEY")
		if !ok {
			os.Stderr.WriteString("environment variable TMDB_API_KEY not set\n")
			os.Exit(1)
		}
		upstream = tmdb.NewClient(key)
	}

	if direct {
		globalClient = upstream
	} else {
		var err error
		globalClient, err = tmdb.NewReplayClient(upstream, dataDir)
		if err != nil {
			os.Stderr.WriteString("Failed to create replay client: " + err.Error() + "\n")
			os.Exit(1)
		}
	}

	os.Exit(m.Run())
}
