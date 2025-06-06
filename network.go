package tmdb

type NetworkId int

type NetworkKey struct {
	Id NetworkId
}

type Network struct {
	NetworkKey

	Logo *Image
	Name *string
	OriginCountry *string
}