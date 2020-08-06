package mercantile

import (
	"fmt"
	"testing"
)

// Test .
func TestCoordTruncate(t *testing.T) {
	c := NewCoordinate(181.0, 55)
	c.Truncate()
	fmt.Println(c)
}
