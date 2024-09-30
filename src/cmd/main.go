package main

import (
	"log"

	routes "code/api"
	db "code/core/db"

	"github.com/gin-gonic/gin"
)

func setupDb() {
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
}

func main() {
	setupDb()

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":20800")
}
