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
	"github.com/stretchr/testify/assert"
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

func checkBackdropImage(t *testing.T, backdropImage tmdb.BackdropImage, config *tmdb.Configuration) bool {
	t.Helper()
	size := config.Images.BackdropSizes[0]

	secureUrl, ok := backdropImage.GetSecureUrl(config, size)
	if !assert.True(t, ok, "BackdropImage %q GetSecureUrl() should support size %q", backdropImage, size) {
		return false
	}
	if !assert.True(t, strings.HasSuffix(secureUrl, string(backdropImage)), "BackdropImage %q GetSecureUrl() should end with %q", secureUrl, backdropImage) {
		return false
	}
	if !assert.True(t, strings.HasPrefix(secureUrl, config.Images.SecureBaseUrl), "BackdropImage %q GetSecureUrl() should start with %q", secureUrl, config.Images.SecureBaseUrl) {
		return false
	}

	insecureUrl, ok := backdropImage.GetUrl(config, size)
	if !assert.True(t, ok, "BackdropImage %q GetUrl() should support size %q", backdropImage, size) {
		return false
	}
	if !assert.True(t, strings.HasSuffix(secureUrl, string(backdropImage)), "BackdropImage %q GetUrl() should end with %q", insecureUrl, backdropImage) {
		return false
	}
	if !assert.True(t, strings.HasPrefix(secureUrl, config.Images.SecureBaseUrl), "BackdropImage %q GetUrl() should start with %q", insecureUrl, config.Images.SecureBaseUrl) {
		return false
	}

	return true
}
