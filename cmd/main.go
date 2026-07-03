package main

import (
	"pokedex-backend/database"
	"pokedex-backend/internal/models"
	"pokedex-backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := database.Connect()

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&models.Pokemon{},
	)

	if err != nil {
		panic(err)
	}

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")
}