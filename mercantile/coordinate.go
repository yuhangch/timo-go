package mercantile

// Coordinate represent a lng lat coordinate.
type Coordinate struct {
	Lng, Lat float64
}

// NewCoordinate create a new coordinate.
func NewCoordinate(lng, lat float64) Coordinate {
	return Coordinate{
		Lng: lng,
		Lat: lat,
	}
}

//  implementation.
func (c *Coordinate) Truncate() {
	if c.Lng > 180.0 {

		c.Lng = 180.0
	} else if c.Lng < -180.0 {
		c.Lng = -180
	}

	if c.Lat > 90.0 {
		c.Lat = 90
	} else if c.Lat < -90.0 {
		c.Lat = -90
	}
}
