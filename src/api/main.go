package main

import (
	"client"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"error": "route not Found",
		})
	})

	usecase = client.NewClient()

	app.POST("/verify", Verify)

	app.GET("/ping", Ping)

	if err := app.Run(":3000"); err != nil {
		log.Println(err)
	}
}
