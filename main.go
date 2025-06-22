package main

import (
	"hospital-system/config"
	"hospital-system/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//log.Fatal("Error loading .env file")
	//}

	config.ConnectDatabase()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
