package tmdb

type NetworkId int

type Network struct {
	Key NetworkId

	Logo *Image
	Name *string
	OriginCountry *string
}