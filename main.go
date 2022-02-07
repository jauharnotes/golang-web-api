package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)

	router.GET("/blog", blogHandler)

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
