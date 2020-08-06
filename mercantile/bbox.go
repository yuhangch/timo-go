package mercantile

// Bbox represent tile bbox.
type Bbox struct {
	Left, Bottom, Top, Right float64
}

// LngLatBbox represent lng lat bbox.
type LngLatBbox struct {
	West, South, East, North float64
}

func (box Bbox) List() (float64, float64, float64, float64) {
	return box.Left, box.Bottom, box.Top, box.Right
}

func (box LngLatBbox) List() (float64, float64, float64, float64) {
	return box.West, box.South, box.East, box.North
}
