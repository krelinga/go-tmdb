package tmdb

type NetworkId int

type Network struct {
	Key  NetworkId
	Data NetworkData
}

type NetworkData struct {
	Logo          *Image
	Name          *string
	OriginCountry *string
}
