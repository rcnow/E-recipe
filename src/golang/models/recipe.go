package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Recipe struct {
	ID         int    `json:"id"`
	RecipeName string `json:"recipe_name"`
	BottleID   string `json:"bottle_id"`
	BottleSize int    `json:"bottle_size"`
	Pg         int    `json:"pg"`
	Vg         int    `json:"vg"`
	Other      int    `json:"other"`
	Nicotine   int    `json:"nicotine"`
	Date       string `json:"date"`
	SteepTime  int    `json:"steep_time"`
	Note       string `json:"note"`
}
type Flavors struct {
	RecipeID      int    `json:"recipe_id"`
	FlavorName    string `json:"flavor_name"`
	FlavorPercent int    `json:"flavor_percent"`
}
type RecipeCollection struct {
	Recipe  []Recipe  `json:"recipe"`
	Flavors []Flavors `json:"flavors"`
}

func GetRecipeList(db *sql.DB) RecipeCollection {
	sql := "SELECT id,recipename FROM recipe"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := RecipeCollection{}
	for rows.Next() {
		recipe := Recipe{}
		err2 := rows.Scan(&recipe.ID, &recipe.RecipeName)
		if err2 != nil {
			log.Fatal(err2)
		}
		result.Recipe = append(result.Recipe, recipe)
	}
	return result
}

//GetRecipe get recipe and flavors by id
func GetRecipe(db *sql.DB, id int) RecipeCollection {
	sql := "SELECT * FROM recipe LEFT JOIN flavor ON recipe.id = flavor.recipeid WHERE id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	result := RecipeCollection{}
	for rows.Next() {
		recipe := Recipe{}
		flavors := Flavors{}
		err2 := rows.Scan(&recipe.ID, &recipe.RecipeName, &recipe.BottleID, &recipe.BottleSize, &recipe.Pg,
			&recipe.Vg, &recipe.Other, &recipe.Nicotine, &recipe.Date, &recipe.SteepTime, &recipe.Note, &flavors.FlavorName, &flavors.FlavorPercent, &flavors.RecipeID)
		if err2 != nil {
			log.Fatal(err2)
		}
		result.Recipe = append(result.Recipe, recipe)
		result.Flavors = append(result.Flavors, flavors)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

//PostRecipe add new recipe
func PostRecipe(db *sql.DB, recipename string, bottleid string, bottlesize int, pg int, vg int, other int, nicotine int, date string, steeptime int, note string) (int64, error) {
	sql := "INSERT INTO recipe(recipename,bottleid,bottlesize,pg,vg,other,nicotine,date,steeptime,note) VALUES(?,?,?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(recipename, bottleid, bottlesize, pg, vg, other, nicotine, date, steeptime, note)
	if err2 != nil {
		log.Fatal(err2)
	}
	return result.LastInsertId()
}

//PostFlavor add new flavor together with recipe
func PostFlavor(db *sql.DB, flavor_name string, flavor_percent int, recipeid int) (int64, error) {
	sql := "INSERT INTO flavor(flavor_name, flavor_percent,recipeid) VALUES(?,?,?);"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(flavor_name, flavor_percent, recipeid)
	if err2 != nil {
		log.Fatal(err2)
	}
	return result.LastInsertId()
}
func UpdateRecipe(db *sql.DB, id int, note string) (int64, error) {
	sql := "update recipe set note=? where id=?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	result, err2 := stmt.Exec(note, id)
	if err2 != nil {
		panic(err2)
	}
	return result.RowsAffected()
}

//DeleteRecipe delete recipe together with flavors by id
func DeleteRecipe(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM recipe WHERE ID = ? "
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	result, err2 := stmt.Exec(id)
	if err2 != nil {
		panic(err2)
	}
	return result.RowsAffected()
}
