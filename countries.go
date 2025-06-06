package tmdb

import "encoding/json"

type CountryIso3166_1 string

type CountrySum struct {
	CountryIso3166_1 CountryIso3166_1
	EnglishName      string
}

// TODO: this seems unnecessary?
func (c *CountrySum) UnmarshalJSON(data []byte) error {
	var raw struct {
		CountryIso3166_1 CountryIso3166_1 `json:"iso_3166_1"`
		Name             string           `json:"name"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	c.CountryIso3166_1 = raw.CountryIso3166_1
	c.EnglishName = raw.Name

	return nil
}
