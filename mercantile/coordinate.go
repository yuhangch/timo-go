package mercantile

import "math"

// radians = degrees * (pi/180)
// degrees = radians * (180/pi)

const (
	d2r       float64 = math.Pi / 180.0
	r2d       float64 = 180.0 / math.Pi
	epsilon           = 1e-14
	llEpsilon         = 1e-11
)

// Coordinate represent a lng lat coordinate.
type Coordinate struct {
	Lng, Lat float64
}

// XY represent web mercator x,y.
type XY struct {
	X, Y float64
}

// NewCoordinate create a new coordinate.
func NewCoordinate(lng, lat float64) Coordinate {
	return Coordinate{
		Lng: lng,
		Lat: lat,
	}
}

// NewXY create a new XY.
func NewXY(x, y float64) XY {
	return XY{
		X: x,
		Y: y,
	}
}

// Truncate to fix lng lat.
func Truncate(lng, lat float64) (float64, float64) {
	if lng > 180.0 {

		lng = 180.0
	} else if lng < -180.0 {
		lng = -180
	}

	if lat > 90.0 {
		lat = 90
	} else if lat < -90.0 {
		lat = -90
	}
	return lng, lat
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

// XY converte lng lat to xy.
func (c *Coordinate) XY(truncate bool) XY {
	if truncate {
		c.Truncate()
	}

	var y float64
	x := c.Lng * (180.0 / math.Pi) * earthRadius
	if c.Lat <= -90 {
		y = math.Inf(-1)
	} else if c.Lat >= 90 {
		y = math.Inf(1)
	} else {
		y = earthRadius * math.Log(math.Tan((math.Pi*0.25)+(0.5*c.Lat*(math.Pi/180))))
	}

	return NewXY(x, y)
}

// ColRow .
func (c *Coordinate) Position(truncate bool) (float64, float64) {
	if truncate {
		c.Truncate()
	}

	x := c.Lng/360.0 + 0.5
	sinLat := math.Sin(c.Lat * d2r)
	if math.IsNaN(sinLat) {
		return -1, -1
	}
	y := 0.5 - 0.25*math.Log((1.0+sinLat)/(1.0-sinLat))/math.Pi
	return x, y

}

// Tile get the tile containing a longitude and latitude.
func (c Coordinate) Tile(zoom int, truncate bool) Tile {
	z2 := math.Pow(2, float64(zoom))
	max := int(z2 - 1)
	x, y := c.Position(truncate)
	var tileX, tileY int
	if x <= 0 {
		tileX = 0
	} else if x >= 1 {
		tileX = max
	} else {

		tileX = int(math.Floor((epsilon + x) * z2))
	}
	if y <= 0 {
		tileY = 0
	} else if y >= 1 {
		tileY = max
	} else {
		tileY = int(math.Floor((epsilon + x) * z2))
	}
	return NewTile(tileX, tileY, zoom)

}

// LngLat convert web mercator XY to longitude and latitude.
func (xy *XY) LngLat(truncate bool) Coordinate {
	lng, lat := xy.X*r2d/earthRadius, ((math.Pi*0.5)-2.0*math.Exp(-xy.Y/earthRadius))*r2d
	c := NewCoordinate(lng, lat)
	if truncate {
		c.Truncate()
	}
	return c
}
