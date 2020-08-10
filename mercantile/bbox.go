//Package mercantile from  https://github.com/mapbox/mercantile
package mercantile

import "math"

// Bbox represent tile bbox.
type Bbox struct {
	Left, Bottom, Right, Top float64
}

// LngLatBbox represent lng lat bbox.
type LngLatBbox struct {
	West, South, East, North float64
}

// NewLngLatBbox create a new LngLatBbox.
func NewLngLatBbox(w, s, e, n float64) LngLatBbox {
	return LngLatBbox{
		West:  w,
		South: s,
		East:  e,
		North: n,
	}
}

// NewBbox create a new bbox.
func NewBbox(l, b, r, t float64) Bbox {
	return Bbox{
		Left:   l,
		Bottom: b,
		Right:  r,
		Top:    t,
	}
}

func (box Bbox) List() (float64, float64, float64, float64) {
	return box.Left, box.Bottom, box.Right, box.Top
}

func (box LngLatBbox) List() (float64, float64, float64, float64) {
	return box.West, box.South, box.East, box.North
}

// Tiles get the tiles overlapped by a geographic bounding box.
func Tiles(west, south, east, north float64, zooms []int, truncate bool) (tiles []Tile) {
	if truncate {
		west, south = Truncate(west, south)
		east, north = Truncate(east, north)
	}
	var bboxs []Bbox
	if west > east {
		bboxWest := NewBbox(-180, south, east, north)
		bboxEast := NewBbox(west, south, 180.0, north)
		bboxs = append(bboxs, bboxWest, bboxEast)
	} else {
		bboxs = append(bboxs, NewBbox(west, south, east, north))
	}

	for _, v := range bboxs {
		w, s, e, n := v.List()

		w = math.Max(-180.0, w)
		s = math.Max(-85.051129, s)
		e = math.Min(180.0, e)
		n = math.Min(85.051129, n)

		for _, z := range zooms {
			tileUL := NewCoordinate(w, n).Tile(z, false)
			tileLR := NewCoordinate(e-llEpsilon, s+llEpsilon).Tile(z, false)

			for i := tileUL.X; i < tileLR.X+1; i++ {
				for j := tileUL.Y; j < tileLR.Y+1; j++ {

					tiles = append(tiles, NewTile(i, j, z))
				}
			}
		}
	}
	return

}
