package tmdb

import "github.com/krelinga/go-views"

type CompanyId int

type Company struct {
	Key  CompanyId
	Data CompanyData

	shows []*Show
}

func (c *Company) AddShow(s *Show) {
	s.AddProductionCompany(c)
}

func (c *Company) Shows() views.Bag[*Show] {
	return views.BagOfSlice[*Show]{S: c.shows}
}

type CompanyData struct {
	Logo          *Image
	Name          *string
	OriginCountry *string
}
