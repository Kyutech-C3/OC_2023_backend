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
	config.LoadConfig()
	db.InitDB()
	r := gin.Default()
	r.Use(cors.Default())
	routers.InitRouter(r)
	if err := r.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
