package routes

import (
	"pokedex-backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.GET(
		"/pokemon",
		controllers.GetPokemons,
	)

	router.GET(
		"/pokemon/:id",
		controllers.GetPokemonByID,
	)

	router.GET(
		"/pokemon/pokedex/:numero",
		controllers.GetPokemonByPokedexID,
	)

	router.POST(
		"/pokemon",
		controllers.CreatePokemon,
	)
	router.PUT(
		"/pokemon/:id",
		controllers.UpdatePokemon,
	)
	router.DELETE(
		"/pokemon/:id",
		controllers.DeletePokemon,
	)
}