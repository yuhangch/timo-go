package api

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yuhangch/timvt/db"
)

func tiles() gin.HandlerFunc {
	return func(c *gin.Context) {
		table := c.Param("table")
		x, y, z, err := xyz(c.Param("x"), c.Param("y"), c.Param("z"))
		if err != nil {

			c.JSON(500, gin.H{
				"message": fmt.Sprintf("%v", err),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("tile is x:%d y:%d z%d from table %s", x, y, z, table),
		})
	}
}

// handleTiles to handle tiles request.
func handleTiles(x, y, z int, table string) {
	conn := db.Conn
	fmt.Println(conn)
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
