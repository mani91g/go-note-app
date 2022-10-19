package main

import (
	"note-app/initializers"
	"note-app/src/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.GET("/note", controllers.GetAllNotes)
	r.GET("/note/:id", controllers.GeNoteById)
	r.PUT("/note:id", controllers.UpdateNote)
	r.POST("/note", controllers.AddNote)
	r.DELETE("/note/:id", controllers.DeleteNote)

	r.Run()
}
