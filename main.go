package main

import (
	"log"
	"oc-2023/config"
	"oc-2023/db"
	"oc-2023/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	cc := cors.DefaultConfig()
	cc.AllowAllOrigins = true
	cc.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	r.Use(cors.New(cc))
	config.LoadConfig()
	db.InitDB()
	routers.Router(r)

	if err := r.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
