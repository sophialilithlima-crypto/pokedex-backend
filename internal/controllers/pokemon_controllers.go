package controllers

import (
	"net/http"

	"pokedex-backend/database"
	"pokedex-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func GetPokemons(c *gin.Context) {

	var pokemons []models.Pokemon

	database.DB.Find(&pokemons)

	c.JSON(
		http.StatusOK,
		pokemons,
	)

}

func CreatePokemon(c *gin.Context) {

	var pokemon models.Pokemon

	if err := c.ShouldBindJSON(&pokemon); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	database.DB.Create(&pokemon)

	c.JSON(
		http.StatusCreated,
		pokemon,
	)

}

func GetPokemonByID(c *gin.Context) {

	id := c.Param("id")

	var pokemon models.Pokemon

	result := database.DB.First(&pokemon, id)

	if result.Error != nil {

		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "Pokémon não encontrado",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		pokemon,
	)

}

func GetPokemonByPokedexID(c *gin.Context) {

	pokedexID := c.Param("numero")

	var pokemon models.Pokemon

	result := database.DB.
		Where("pokedex_id = ?", pokedexID).
		First(&pokemon)

	if result.Error != nil {

		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "Pokémon não encontrado",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		pokemon,
	)

}

func UpdatePokemon(c *gin.Context) {

	id := c.Param("id")

	var pokemon models.Pokemon

	if err := database.DB.First(&pokemon, id).Error; err != nil {

		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "Pokémon não encontrado",
			},
		)

		return

	}

	var dados models.Pokemon

	if err := c.ShouldBindJSON(&dados); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return

	}

	pokemon.PokedexID = dados.PokedexID
	pokemon.Name = dados.Name
	pokemon.Type1 = dados.Type1
	pokemon.Type2 = dados.Type2
	pokemon.Height = dados.Height
	pokemon.Weight = dados.Weight
	pokemon.Description = dados.Description

	database.DB.Save(&pokemon)

	c.JSON(
		http.StatusOK,
		pokemon,
	)

}

func DeletePokemon(c *gin.Context) {

	id := c.Param("id")

	var pokemon models.Pokemon

	if err := database.DB.First(&pokemon, id).Error; err != nil {

		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "Pokémon não encontrado",
			},
		)

		return

	}

	database.DB.Delete(&pokemon)

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Pokémon removido com sucesso",
		},
	)

}