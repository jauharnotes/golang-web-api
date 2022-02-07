package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/blog", blogHandler)
	router.GET("/books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)
	router.POST("/books", booksPostHandler)

	router.Run(":5000")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Jauharuddin",
		"bio":  "sofware enginer & content creator",
	})
}

func blogHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":    "Belajar Golang Web API",
		"tanggal":  "06 Februari 2022",
		"paragraf": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam tristique erat et neque lacinia placerat.",
	})

}

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	terbit := c.Query("terbit")

	c.JSON(http.StatusOK, gin.H{"title": title, "terbit": terbit})
}

type BookInput struct {
	Price int    `json:"price" binding:"required,number"`
	Title string `json:"title" binding:"required"`
}

func booksPostHandler(c *gin.Context) {
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)

	errorMessages := []string{}
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}
