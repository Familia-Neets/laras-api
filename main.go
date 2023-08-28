package main

import (
	"Lara/api"
	"Lara/models"
)

func main() {
	models.DatabaseConnect()
	models.Migrate(models.Movie{}, models.Series{}, models.User{}, models.Game{}, models.Book{})
	api.Run()
}
