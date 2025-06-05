package tmdb

import (
	"time"
)

// A Date represnted in the format YYYY-MM-DD.
type DateYYYYMMDD string

func (d DateYYYYMMDD) AsTime() (time.Time, error) {
	return time.Parse("2006-01-02", string(d))
}
