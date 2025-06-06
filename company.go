package tmdb

type CompanyId int

type CompanyKey struct {
	Id CompanyId
}

type Company struct {
	CompanyKey

	Logo *Image
	Name *string
	OriginCountry *string
}
