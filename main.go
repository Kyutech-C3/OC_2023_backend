package main

import (
	"log"
	"oc-2023/config"
	"oc-2023/db"
	"oc-2023/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	config.LoadConfig()
	db.InitDB()
	routers.Router(r)

	if err := r.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
