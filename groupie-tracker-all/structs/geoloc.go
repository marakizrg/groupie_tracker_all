package structs

type GeocodeResponse struct {
	Results []struct {
		Geometry struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"geometry"`
	} `json:"results"`
}

type LocationWithCoordinates struct {
	Name string // "Berlin, Germany"
	Lat  float64
	Lng  float64
}
