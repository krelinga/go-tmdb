package tmdb

type NetworkId int

type Network struct {
	Key  NetworkId
	Data NetworkData

	shows []*Show
}

func (n *Network) AddShow(s *Show) {
	s.AddNetwork(n)
}

type NetworkData struct {
	Logo          *Image
	Name          *string
	OriginCountry *string
}
