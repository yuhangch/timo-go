package api

import (
	"github.com/gin-gonic/gin"
)

// API get gin router for api.
func API() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", hello())
	r.GET("/tiles/:table/:z/:x/:y", tiles())
	return r
}

func hello() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "hello timvt",
		})
	}
}
