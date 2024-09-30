package main

import (
	"fmt"
	"log"

	routes "code/api"
	db "code/core/db"
	"code/core/setting"

	"github.com/gin-gonic/gin"
)

func setupDb() {
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
}

func main() {
	setting.Setup()
	setupDb()

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(fmt.Sprintf(":%s", setting.AppSetting.Port))
}
