package main

import (
	"note-app/initializers"
	"note-app/src/models"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Note{})
}
