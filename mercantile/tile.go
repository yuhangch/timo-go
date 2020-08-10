//Package mercantile from  https://github.com/mapbox/mercantile
package mercantile

import (
	"math"
	"strconv"
	"strings"
)

const (
	earthRadius float64 = 6378137.0
	ce          float64 = 2 * math.Pi * earthRadius
)

// Tile present a tile.
type Tile struct {
	X, Y, Z int
}

// NewTile to make a new tile .
func NewTile(x, y, z int) Tile {
	return Tile{
		X: x,
		Y: y,
		Z: z,
	}
}

//  implementation.
func (t Tile) XYZ() (int, int, int) {
	return t.X, t.Y, t.Z
}

func (t Tile) UL() Coordinate {

	x, y, z := t.XYZ()
	zp := math.Pow(2, float64(z))
	lngDeg := float64(x)/zp*360.0 - 180
	latRad := math.Atan(math.Sinh(math.Pi * (1 - 2*float64(y)/zp)))
	latDeg := latRad * (180.0 / math.Pi)

	return NewCoordinate(lngDeg, latDeg)
}

//  implementation.
func (t Tile) Bounds() LngLatBbox {
	x, y, z := t.XYZ()
	tLR := NewTile(x+1, y+1, z)

	a := t.UL()
	b := tLR.UL()

	return NewLngLatBbox(a.Lng, b.Lat, b.Lng, a.Lat)

}

// XYBounds get the web mercator bounding box of a tile.
func (t *Tile) XYBounds() Bbox {
	x, y, z := t.XYZ()
	tileSize := ce / math.Pow(2, float64(z))
	left := float64(x)*tileSize - (ce / 2)
	right := left + tileSize
	top := (ce / 2) - float64(y)*tileSize
	bottom := top - tileSize
	return NewBbox(left, bottom, right, top)
}

// QuadKey get the quadkey of a tile.
func (t *Tile) QuadKey() string {
	x, y, z := t.XYZ()
	var qk []string
	for i := z; i > 0; i-- {
		d := 0
		mask := 1 << (i - 1)
		// fmt.Println(mask)
		if !((x & mask) == 0) {
			d++
		}

		if !((y & mask) == 0) {
			d += 2
		}
		qk = append(qk, strconv.Itoa(d))
	}
	return strings.Join(qk, "")

}

// TODO:
// QuadKeyToTile get tile from quadkey.
func QuadKeyToTile() {

}
