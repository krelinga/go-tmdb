package tmdb

import (
	"slices"
	"strings"
)

type column interface {
	Endpoint() string
}

func appendToResponse[T column](columns []T) string {
	stringCols := make([]string, len(columns))
	for _, col := range columns {
		stringCols = append(stringCols, col.Endpoint())
	}
	slices.Sort(stringCols)
	return strings.Join(stringCols, ",")
}