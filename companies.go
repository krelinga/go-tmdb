package tmdb

type ProductionCompanyId int

type ProductionCompanyShort struct {
	Id            ProductionCompanyId `json:"id"`
	LogoImage     LogoImage           `json:"logo_path"`
	Name          string              `json:"name"`
	OriginCountry string              `json:"origin_country"`
}
