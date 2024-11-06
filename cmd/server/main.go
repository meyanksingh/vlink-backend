package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	database "github.com/meyanksingh/vlink-backend/internal/db"
	routes "github.com/meyanksingh/vlink-backend/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error with loading ENv File")
	}

	database.ConnectDB()
	port := os.Getenv("PORT")

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)

	router.Run(":" + port)

}
