package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"errors"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

var books = []book{
	{ID: "1", Title: "Преступление и Наказание", Author: "Фёдор Михайлович Достоевский", Quantity: 15, Price: 2000},
	{ID: "2", Title: "Идиот", Author: "Фёдор Михайлович Достоевский", Quantity: 9, Price: 2700},
	{ID: "3", Title: "Капитанская Дочка", Author: "Александр Сергеевич Пушкин", Quantity: 35, Price: 2200},
	{ID: "4", Title: "Палата №6", Author: "Антон Павлович Чехов", Quantity: 50, Price: 2000},
	{ID: "5", Title: "Война и Мир", Author: "Лев Николаевич Толстой", Quantity: 23, Price: 3000},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBook(c *gin.Context) {
	id := c.Param("id")
	for _, b := range books {
		if b.ID == id {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Книга не найдена"})
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Книга удалена"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Книга не найдена"})
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBook)
	router.POST("/books", createBook)
	router.DELETE("/books/:id", deleteBook)
	router.Run("localhost:8080")
}
