package tmdb

type TvNetworkId int

type TvNetwork struct {
	TvNetworkId   TvNetworkId      `json:"id"`
	LogoImage     LogoImage        `json:"logo_path"`
	Name          string           `json:"name"`
	OriginCountry CountryIso3166_1 `json:"origin_country"`
}
