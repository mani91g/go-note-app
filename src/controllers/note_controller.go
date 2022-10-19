package controllers

import (
	"fmt"
	"net/http"
	"note-app/initializers"
	"note-app/src/models"

	"github.com/gin-gonic/gin"
)

func GetAllNotes(c *gin.Context) {
	var notes []models.Note
	initializers.DB.Find(&notes)
	c.IndentedJSON(200, notes)
}

func GeNoteById(c *gin.Context) {
	id := c.Param("id")

	var note models.Note
	initializers.DB.First(&note, id)

	if note.Id != "" {
		c.JSON(http.StatusAccepted, note)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Note not found"})
}

func AddNote(c *gin.Context) {
	fmt.Println("Add note called")
	var newNote models.Note

	if err := c.BindJSON(&newNote); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	result := initializers.DB.Create(&newNote)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
	}

	c.IndentedJSON(http.StatusCreated, newNote)
}

func UpdateNote(c *gin.Context) {
	var request models.Note
	c.BindJSON(&request)

	var note models.Note
	id := c.Param("id")
	initializers.DB.First(&note, id)

	initializers.DB.Model(&note).Updates(request)
	c.IndentedJSON(http.StatusNoContent, note)
}

func DeleteNote(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Note{}, id)
	c.Status(http.StatusOK)
}
