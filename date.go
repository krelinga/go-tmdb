package tmdb

import (
	"fmt"
	"strconv"
	"strings"
)

type Date interface {
	Raw() string

	// Returns any error that came while parsing the raw date.
	Err() error

	// These methods will panic if Err() returns a non-nil error.
	Year() int
	Month() int
	Day() int
}

func newDateYYYYMMDD(raw string) Date {
	parts := strings.Split(raw, "-")
	if len(parts) != 3 {
		return &errDate{raw: raw, err: fmt.Errorf("invalid date format, expected YYYY-MM-DD but got %s", raw)}
	}
	year, err := strconv.Atoi(parts[0])
	if err != nil || year < 1000 || year > 9999 {
		return &errDate{raw: raw, err: fmt.Errorf("invalid year in date %s: %w", raw, err)}
	}
	month, err := strconv.Atoi(parts[1])
	if err != nil || month < 1 || month > 12 {
		return &errDate{raw: raw, err: fmt.Errorf("invalid month in date %s: %w", raw, err)}
	}
	day, err := strconv.Atoi(parts[2])
	if err != nil || day < 1 || day > 31 {
		return &errDate{raw: raw, err: fmt.Errorf("invalid day in date %s: %w", raw, err)}
	}
	return &okDate{
		raw:   raw,
		year:  year,
		month: month,
		day:   day,
	}
}

type errDate struct {
	raw string
	err error
}

func (d *errDate) Raw() string {
	return d.raw
}

func (d *errDate) Err() error {
	return d.err
}

func (d *errDate) Year() int {
	panic("cannot call Year() on a Date with an error: " + d.err.Error())
}

func (d *errDate) Month() int {
	panic("cannot call Month() on a Date with an error: " + d.err.Error())
}

func (d *errDate) Day() int {
	panic("cannot call Day() on a Date with an error: " + d.err.Error())
}

type okDate struct {
	raw   string
	year  int
	month int
	day   int
}

func (d *okDate) Raw() string {
	return d.raw
}

func (d *okDate) Err() error {
	return nil
}

func (d *okDate) Year() int {
	return d.year
}

func (d *okDate) Month() int {
	return d.month
}

func (d *okDate) Day() int {
	return d.day
}