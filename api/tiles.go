package api

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/yuhangch/timo/mercantile"
)

const (
	// 	sqlTpl string = `
	// 	WITH
	// 	bounds AS (
	// 		SELECT
	// 			ST_Segmentize(
	// 				ST_MakeEnvelope(
	// 					{{printf "%.4f" .Bbox.Left}},
	// 					{{printf "%.4f" .Bbox.Bottom}},
	// 					{{printf "%.4f" .Bbox.Right}},
	// 					{{printf "%.4f" .Bbox.Top}},
	// 					3857
	// 				),
	// 				{{printf "%.4f" .SegSize}}
	// 			) AS geom
	// 	),
	// 	mvtgeom AS (
	// 		SELECT ST_AsMVTGeom(
	// 			ST_Transform(t.{{.GeomColName}}, 4326),
	// 			bounds.geom,
	// 			{{.TileResolution}},
	// 			{{.TileBuffer}}
	// 		) AS geom, {{.Cols}}
	// 		FROM {{.TableName}} t, bounds
	// 		WHERE ST_Intersects(
	// 			ST_Transform(t.{{.GeomColName}}, 4326), ST_Transform(bounds.geom, 4326)
	// 		) LIMIT 10
	// 	)
	// 	SELECT ST_AsMVT(mvtgeom.*) FROM mvtgeom
	// `

	sqlTpl string = `
with geoms as(
	SELECT ST_AsMVTGeom(ST_Transform({{.GeomColName}}, 3857),ST_TileEnvelope({{.Tile.Z}},{{.Tile.X}},{{.Tile.Y}})) as geom {{.Cols}}
		from {{.TableName}}
		where  ST_Intersects(ST_Transform({{.GeomColName}}, 3857), ST_TileEnvelope({{.Tile.Z}},{{.Tile.X}},{{.Tile.Y}})) 
		{{.Filter}}
		
	)
	SELECT ST_AsMVT(geoms.*) FROM geoms
`
)

func tiles() gin.HandlerFunc {
	return func(c *gin.Context) {

		data, err := handleTiles(c)
		if err != nil {
			c.JSON(500, gin.H{
				"msg": err,
			})
			return
		}

		w := c.Writer
		w.Header().Add("Content-Type", "application/vnd.mapbox-vector-tile")
		w.Write(data)
	}
}

// TilesContext represent parameters in tile request.
type TilesContext struct {
	TableName, GeomColName, Cols, Filter string
	mercantile.Tile
}

// NewTilesCtx create a new context for tile request.
func NewTilesCtx(tableName, geom, cols, filters string, tile mercantile.Tile) *TilesContext {
	// fmt.Println(tableName, geom, cols, filters)
	return &TilesContext{
		TableName:   tableName,
		GeomColName: geom,
		Cols:        cols,
		Tile:        tile,
		Filter:      filters,
	}

}

// handleTiles to handle tiles request.
func handleTiles(c *gin.Context) ([]byte, error) {

	var data []byte
	ctx, err := handleTilesParams(c)
	if err != nil {
		return []byte{}, err
	}
	query := sql(ctx)
	// fmt.Println(query)
	err = conn.QueryRow(context.Background(), query).Scan(&data)

	if err != nil {
		return []byte{}, err
	}
	if len(data) == 0 {
		return data, fmt.Errorf("empty")
	}
	return data, nil

}

// handleTilesQueryParams handle query parameters.
func handleTilesParams(c *gin.Context) (*TilesContext, error) {
	table := c.Param("table")
	if table == "" {
		return nil, fmt.Errorf("require table name")
	}
	x, y, z, err := xyz(c.Param("x"), c.Param("y"), c.Param("z"))
	if err != nil {
		return nil, fmt.Errorf("parse x,y,z error")
	}
	columns := c.Query("columns")
	if len(columns) > 0 {
		clsNames := strings.Split(columns, ",")

		if len(clsNames) > 1 {
			for i, v := range clsNames {
				clsNames[i] = fmt.Sprintf(`"%s"`, v)
			}
			columns = strings.Join(clsNames, ", ")
		}
		columns = fmt.Sprintf(", %s", columns)

	}
	filters := c.Query("filters")
	if len(filters) > 0 {
		fts := strings.Split(filters, ",")

		filters = strings.Join(fts, " , and")
		filters = fmt.Sprintf(" and %s", filters)
	}
	geom := c.Query("geom")
	if geom == "" {
		geom = "geom"
	}
	tile := mercantile.NewTile(x, y, z)
	return NewTilesCtx(table, geom, columns, filters, tile), nil

}

// sql get sql from sql pattern and context.
func sql(ctx *TilesContext) string {
	tmpl, err := template.New("sql").Parse(sqlTpl)
	if err != nil {
		fmt.Println(err)
	}
	var tpl bytes.Buffer
	err = tmpl.ExecuteTemplate(&tpl, "sql", ctx)
	if err != nil {
		fmt.Println(err)
	}
	return tpl.String()
}

func xyz(xstr, ystr, zstr string) (int, int, int, error) {
	x, err := strconv.Atoi(xstr)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("%v", err)
	}
	ok := strings.HasSuffix(ystr, ".pbf")
	if !ok {
		return 0, 0, 0, fmt.Errorf("%v", err)
	}
	ystr = ystr[0 : len(ystr)-4]

	y, err := strconv.Atoi(ystr)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("%v", err)
	}

	z, err := strconv.Atoi(zstr)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("%v", err)
	}
	return x, y, z, nil
}
