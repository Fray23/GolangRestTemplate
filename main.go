package main

import (
	routes "code/api"
	db "code/core/db"
	"log"

	"github.com/gin-gonic/gin"
)

func setupDb() {
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	db.Migrate()
}

func main() {
	setupDb()

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":20800")
}
