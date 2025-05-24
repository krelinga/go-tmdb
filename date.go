package tmdb

import "time"

type Date string

func (d Date) GetTime() (time.Time, error) {
	t, err := time.Parse("2006-01-02", string(d))
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}