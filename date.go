package tmdb

import "time"

type DateYYYYMMDD string

func (d DateYYYYMMDD) GetTime() (time.Time, error) {
	t, err := time.Parse("2006-01-02", string(d))
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

type DateRfc3339 string

func (d DateRfc3339) GetTime() (time.Time, error) {
	t, err := time.Parse(time.RFC3339, string(d))
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

type DateInterface interface {
	GetTime() (time.Time, error)
}
