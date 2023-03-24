package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID      string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"desc"`
}

var BookDatas = []Book{}

// Create
func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	newBook.BookID = fmt.Sprintf("c%d", len(BookDatas)+1)
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"book": newBook,
	})
}

// Update
func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var updateBook Book

	if err := ctx.ShouldBindJSON(&updateBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		if book.BookID == bookID {
			condition = true
			BookDatas[i] = updateBook
			BookDatas[i].BookID = bookID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  http.StatusNotFound,
			"error_message": fmt.Sprintf("Book with ID %s not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with ID %s has been updated", bookID),
	})
}

// Delete
func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false

	for i, book := range BookDatas {
		if book.BookID == bookID {
			condition = true
			BookDatas = append(BookDatas[:i], BookDatas[i+1:]...)
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  http.StatusNotFound,
			"error_message": fmt.Sprintf("Book with ID %s not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with ID %s has been deleted", bookID),
	})
}

// Get One Book
func GetBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var book Book

	for _, book = range BookDatas {
		if book.BookID == bookID {
			condition = true
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  http.StatusNotFound,
			"error_message": fmt.Sprintf("Book with ID %s not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

// Get All Books
func GetBooks(ctx *gin.Context) {

	if len(BookDatas) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  http.StatusNotFound,
			"error_message": "Empty Book Data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"books": BookDatas,
	})
}
