package main

import (
	"Lara/api"
	"Lara/models"
	"Lara/models/contents"
	"Lara/models/users"
)

func main() {
	models.DatabaseConnect()
	models.Migrate(contents.Movie{}, contents.Series{}, users.User{}, contents.Game{}, contents.Book{})
	api.Run()
}
