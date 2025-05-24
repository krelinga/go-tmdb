package tmdb

import "time"

type Minutes int

func (m Minutes) GetDuration() time.Duration {
	return time.Duration(m) * time.Minute
}