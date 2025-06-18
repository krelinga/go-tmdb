package tmdb

type CompanyId int

type Company struct {
	Key CompanyId

	Logo *Image
	Name *string
	OriginCountry *string
}
