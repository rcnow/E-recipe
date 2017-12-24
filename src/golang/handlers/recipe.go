package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"../models"

	"github.com/labstack/echo"
)

type H map[string]interface{}

//GetRecipeList receiving full recipe list
func GetRecipeList(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetRecipeList(db))
	}
}

//GetRecipe receiving one recipe with id
func GetRecipe(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, models.GetRecipe(db, id))
	}
}

//PostRecipe add new recipe
func PostRecipe(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var recipe models.Recipe
		c.Bind(&recipe)
		id, err := models.PostRecipe(db, recipe.RecipeName, recipe.BottleID, recipe.BottleSize, recipe.Pg, recipe.Vg, recipe.Other, recipe.Nicotine,
			recipe.Date, recipe.SteepTime, recipe.Note)
		if err == nil {
			return c.JSON(http.StatusCreated,
				H{
					"New recipe created id - ": id,
				})
		}
		return err
	}
}

//PostFlavor add new flavor with recipe id
func PostFlavor(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var flavors []models.Flavors
		var ids []int64
		c.Bind(&flavors)
		for _, val := range flavors {
			id, err := models.PostFlavor(db, val.FlavorName, val.FlavorPercent, val.RecipeID)
			if err != nil {
				return err
			}
			ids = append(ids, id)
		}
		return c.JSON(http.StatusCreated, H{
			"New flavor created id - ": ids,
		})
	}
}
func UpdateRecipe(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var recipe models.Recipe
		c.Bind(&recipe)
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := models.UpdateRecipe(db, id, recipe.Note)
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"update": id,
			})
		}
		return err
	}
}

//DeleteRecipe delete one recipe by id
func DeleteRecipe(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := models.DeleteRecipe(db, id)
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		}
		return err
	}
}
