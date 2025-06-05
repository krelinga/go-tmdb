package tmdb

type CompanyId int

type Company struct {
	Id CompanyId

	Logo *Image
	Name *string
	OriginCountry *string
}
