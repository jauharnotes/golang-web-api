package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Jauharuddin",
		"bio":  "sofware enginer & content creator",
	})
}

func BlogHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":    "Belajar Golang Web API",
		"tanggal":  "06 Februari 2022",
		"paragraf": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam tristique erat et neque lacinia placerat.",
	})

}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	terbit := c.Query("terbit")

	c.JSON(http.StatusOK, gin.H{"title": title, "terbit": terbit})
}

func BooksPostHandler(c *gin.Context) {
	var bookInput book.BookInput

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
