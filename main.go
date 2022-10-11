package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Note struct {
	Id string	`json:"id"`
	Name string `json:"title"`
	Author string `json:"author"`
}

var Notes = []Note{
	{
		Id: "1",
		Name: "Test note 1",
		Author: "Mani",
	},
	{
		Id: "2",
		Name: "Test note 2",
		Author: "Mani",
	},
}

func getAllNotes (c *gin.Context) {
	c.IndentedJSON(200, Notes)
}

func geNoteById (c *gin.Context) {
	id := c.Param("id")
	// id, _ := strconv.Atoi(i)

	for _, note := range Notes {
		if id == note.Id {
			fmt.Println(note)
			c.IndentedJSON(http.StatusOK, note)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Note not found"})
}

func addNote (c *gin.Context) {
	fmt.Println("Add note called")
	var newNote Note

	if err := c.BindJSON(&newNote); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	Notes = append(Notes, newNote)
	fmt.Println(Notes)
	c.IndentedJSON(http.StatusCreated, newNote)
}

func returnCompressedNoteAsAttachment (c *gin.Context) {
	file, _ := os.Open("")
	read := bufio.NewReader(file)
	data, _ := ioutil.ReadAll(read)
	c.Writer.Write(data)
}


func returnNoteAsAttachment (c *gin.Context) {
	c.File("~/downloads/test-5mb.bin")
}




func main() {
	fmt.Println("Hello world")

	r := gin.Default()
	r.POST("/original", returnNoteAsAttachment)
	r.POST("/compressed", returnCompressedNoteAsAttachment)
	r.GET("/", getAllNotes)
	r.GET("/:id", geNoteById)
	r.POST("/", addNote)
	r.Run()
}