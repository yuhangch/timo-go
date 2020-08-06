package mercantile

import "math"

// Tile present a tile.
type Tile struct {
	X, Y, Z float64
}

// NewTile to make a new tile .
func NewTile(x, y, z float64) Tile {
	return Tile{
		X: x,
		Y: y,
		Z: z,
	}
}

//  implementation.
func (t Tile) XYZ() (float64, float64, float64) {
	return t.X, t.Y, t.Z
}

func (t Tile) UL() Coordinate {

	x, y, z := t.XYZ()
	zp := math.Pow(2, z)
	lngDeg := x/zp*360.0 - 180
	latRad := math.Atan(math.Sinh(math.Pi * (1 - 2*y/zp)))
	latDeg := latRad * (180.0 / math.Pi)

	return NewCoordinate(lngDeg, latDeg)
}

//  implementation.
func (t Tile) Bounds() {

}
