package tmdb

type CountryId string

type Country interface {
	Id() CountryId
	Name() string
}

type country struct {
	id   CountryId
	name string
}

func (c country) Id() CountryId {
	return c.id
}

func (c country) Name() string {
	return c.name
}
