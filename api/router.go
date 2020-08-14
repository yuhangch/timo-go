package api

import (
	"context"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

var conn *pgxpool.Pool

// API init router and pg pool
func API(url string) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:1234", "https://yuhangch.github.io"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))
	r.Use(location.Default())
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		panic("unable parse database config")
	}
	conn, err = pgxpool.ConnectConfig(context.Background(), cfg)

	r.GET("/hello", hello())
	r.GET("/", hello())
	r.GET("/tilejson/:table", tilejson())
	r.GET("/tiles/:table/:z/:x/:y", tiles())

	return r
}

func hello() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Hello ,This is timo",
		})
	}
}
