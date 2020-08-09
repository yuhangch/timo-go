package mercantile

import (
	"testing"
)

var (
	tile Tile = NewTile(2, 3, 4)
)

// TestTile test tile.
func TestTile(t *testing.T) {
	if tile.UL().Lng != -135.0 || tile.UL().Lat != 74.01954331150226 {
		t.Error()
	}
}

// Test .
func TestQuadKey(t *testing.T) {
	t2 := NewTile(2, 3, 4)
	if t2.QuadKey() != "0032" {
		t.Error()
	}
}

// Test tile.XYBounds.
func TestXYBounds(t2 *testing.T) {
	tile2 := NewTile(10, 10, 10)
	l, b, r, t := tile2.XYBounds().List()
	if l != -19646150.75796914 {
		t2.Error()
	}
	if b != 19607014.99948713 {
		t2.Error()
	}
	if r != -19607014.99948713 {
		t2.Error()
	}
	if t != 19646150.75796914 {
		t2.Error()
	}

}
