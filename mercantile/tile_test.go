package mercantile

import (
	"testing"
)

// TestTile test tile.
func TestTile(t *testing.T) {
	tile := NewTile(2, 3, 4)
	if tile.UL().Lng != -135.0 || tile.UL().Lat != 74.01954331150226 {
		t.Error("ul test failed")
	}
}
