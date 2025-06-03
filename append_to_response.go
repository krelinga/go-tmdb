package tmdb

import (
	"slices"
	"strings"
)

type part interface {
	Endpoint() string
}

func appendToResponse[T part](parts []T) string {
	stringParts := make([]string, len(parts))
	for _, part := range parts {
		stringParts = append(stringParts, part.Endpoint())
	}
	slices.Sort(stringParts)
	return strings.Join(stringParts, ",")
}