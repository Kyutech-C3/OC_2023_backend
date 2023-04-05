package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))

	hub := newHub()
	go hub.run()

	// r.GET("/", func(c *gin.Context) {
	// 	http.ServeFile(c.Writer, c.Request, "index.html")
	// })

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.GET("/ws", func(c *gin.Context) {
		serveWs(hub, c.Writer, c.Request)
	})
	if err := r.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
