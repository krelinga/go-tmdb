package tmdb

type CompanyId int

type Company struct {
	Key CompanyId
	Data CompanyData
}

type CompanyData struct {
	Logo *Image
	Name *string
	OriginCountry *string
}