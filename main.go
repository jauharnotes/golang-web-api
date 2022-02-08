package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:pas@tcp(127.0.0.1:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db Connection Error")
	}

	db.AutoMigrate(&book.Book{})
	// CRUD

	// Create

	// book := book.Book{}
	// book.Title = "sapies"
	// book.Price = 80000
	// book.Rating = 5
	// book.Description = "ini adalah buku sejarah recomendasi yg bagus"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("=======================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("=======================")
	// }

	// =====
	// Read
	// =====
	// var books []book.Book

	// err = db.Debug().Where("rating", 5).Find(&books).Error
	// if err != nil {
	// 	fmt.Println("=======================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("=======================")
	// }

	// for _, b := range books {
	// 	fmt.Println("Title : ", b.Title)
	// 	fmt.Printf("Title object %v ", b)
	// }

	// =====
	// Update
	// =====
	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("=======================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("=======================")
	// }

	// book.Title = "Tipping Point (Revised edition)"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("=======================")
	// 	fmt.Println("Error updating book record")
	// 	fmt.Println("=======================")
	// }

	// =======
	// Delete
	// =======
	var book book.Book

	err = db.Debug().Last(&book).Error
	if err != nil {
		fmt.Println("=======================")
		fmt.Println("Error finding book record")
		fmt.Println("=======================")
	}

	book.Title = "Tipping Point (Revised edition)"
	err = db.Delete(&book).Error
	if err != nil {
		fmt.Println("=======================")
		fmt.Println("Error deleting book record")
		fmt.Println("=======================")
	}

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/blog", handler.BlogHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.BooksPostHandler)

	router.Run(":5000")
}
