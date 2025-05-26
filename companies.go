package tmdb

type ProductionCompanyShort struct {
	Id            int       `json:"id"`
	LogoImage     LogoImage `json:"logo_path"`
	Name          string    `json:"name"`
	OriginCountry string    `json:"origin_country"`
}
