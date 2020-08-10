package api

import (
	"fmt"
	"strings"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

// TileJSON represent tilejson only include min/max zoom ,url,name.
type TileJSON struct {
	Minzoom int      `json:"minzoom"`
	Maxzoom int      `json:"maxzoom"`
	Name    string   `json:"name"`
	Tiles   []string `json:"tiles"`
}

// tilejson get tilejson handle func .
func tilejson() gin.HandlerFunc {
	return func(c *gin.Context) {

		data, err := handleTilejson(c)
		if err != nil {
			c.JSON(500, gin.H{
				"msg": err,
			})
			return
		}
		c.JSON(200, data)

	}
}

// handleTilejson handle tilejson request.
func handleTilejson(c *gin.Context) (*TileJSON, error) {
	table := c.Param("table")
	if table == "" {
		return nil, fmt.Errorf("require table name")
	}
	if !strings.HasSuffix(table, ".json") {
		return nil, fmt.Errorf("can't parse table")
	}
	table = table[0 : len(table)-5]
	url := location.Get(c)
	URL := c.Request.URL.String()
	host := url.Host
	URL = strings.Replace(URL, "tilejson/", "tiles/", 1)
	URL = strings.Replace(URL, ".json", "/{z}/{x}/{y}.pbf", 1)
	URL = fmt.Sprintf("http://%s%s", host, URL)
	tilejson := NewTileJSON(table, URL)
	return &tilejson, nil

}

// NewTileJSON create a new tilejson.
func NewTileJSON(name, url string) TileJSON {
	return TileJSON{
		Name:    name,
		Tiles:   []string{url},
		Minzoom: 0,
		Maxzoom: 24,
	}
}
