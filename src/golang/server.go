package main

import (
	"database/sql"
	"log"

	"./handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db := initDB("db/recipe.db")

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Static("/static", "../../dist/static")
	e.File("/", "../../dist/index.html")
	e.GET("/recipe", handlers.GetRecipeList(db))
	e.GET("/recipeview/:id", handlers.GetRecipe(db))
	e.POST("/recipe/new", handlers.PostRecipe(db))
	e.POST("/flavor/new", handlers.PostFlavor(db))
	e.PUT("/recipe/update/:id", handlers.UpdateRecipe(db))
	e.DELETE("recipe/delete/:id", handlers.DeleteRecipe(db))

	e.Logger.Fatal(e.Start(":8000"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}
	if db == nil {
		log.Fatal("db nil")
	}
	return db
}
