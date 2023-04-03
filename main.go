package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func GinMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

func main() {
	r := gin.New()
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		s.Join("bcast")
		return nil
	})

	server.OnEvent("/", "msg", func(s socketio.Conn, msg string) {
		server.BroadcastToRoom("", "bcast", "reply", msg)
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	// r.Use(cors.New(config))

	r.Use(GinMiddleware("http://localhost:5500"))
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{
	// 		"http://localhost:5500",
	// 		"http://localhost:3000",
	// 	},
	// 	AllowMethods: []string{
	// 		"POST",
	// 		"GET",
	// 		"PATCH",
	// 		"DELETE",
	// 		"OPTIONS",
	// 	},
	// 	AllowHeaders: []string{
	// 		"Access-Control-Allow-Credentials",
	// 		"Access-Control-Allow-Headers",
	// 		"Content-Type",
	// 		"Content-Length",
	// 		"Accept-Encoding",
	// 	},
	// 	AllowWebSockets: true,
	// }))

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))
	// r.StaticFS("/public", http.Dir("./assets"))
	if err := r.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
