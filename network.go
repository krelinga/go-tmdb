package tmdb

type NetworkId int

type Network struct {
	Id   NetworkId

	Logo *Image
	Name *string
	OriginCountry *string
}