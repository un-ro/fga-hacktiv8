package controllers

import (
	"fmt"
	"net/http"
	"web-swagger/database"
	"web-swagger/models"

	"github.com/gin-gonic/gin"
)

// GetAllCars godoc
// @Summary Get details
// @Description Get details of all cars
// @Tags cars
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Car
// @Router /cars [get]
func GetAllCars(c *gin.Context) {
	var db = database.GetDB()
	var cars []models.Car

	err := db.Find(&cars).Error

	if err != nil {
		fmt.Println("Error getting cars: ", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": cars})
}

// GetOneCars godoc
// @Summary Get details for a given Id
// @Description Get details of car corresponding to the input Id
// @Tags cars
// @Accept  json
// @Produce  json
// @Param Id path int true "ID of the car"
// @Success 200 {object} models.Car
// @Router /cars/{Id} [get]
func GetOneCars(c *gin.Context) {
	var db = database.GetDB()
	var car models.Car

	err := db.First(&car, "Id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": car})
}

// CreateCars godoc
// @Summary Create a new car details
// @Description Create a new car details
// @Tags cars
// @Accept  json
// @Produce  json
// @Param models.Car body models.Car true "Car details"
// @Success 200 {object} models.Car
// @Router /cars [post]
func CreateCars(c *gin.Context) {
	var db = database.GetDB()
	var input models.Car

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car := models.Car{Merk: input.Merk, Harga: input.Harga, Typecars: input.Typecars, Pemilik: input.Pemilik}
	db.Create(&car)

	c.JSON(http.StatusOK, gin.H{"data": car})
}

// UpdateCars godoc
// @Summary Update a car details
// @Description Update a car details
// @Tags cars
// @Accept  json
// @Produce  json
// @Param Id path int true "ID of the car"
// @Success 200 {object} models.Car
// @Router /cars/{Id} [patch]
func UpdateCars(c *gin.Context) {
	var db = database.GetDB()
	var car models.Car

	// Check data is exist or not
	err := db.First(&car, "Id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.Car
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update data
	db.Model(&car).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": car})
}

// DeleteCars godoc
// @Summary Delete a car details
// @Description Delete a car details
// @Tags cars
// @Accept  json
// @Produce  json
// @Param Id path int true "ID of the car"
// @Success 204 "No Content"
// @Router /cars/{Id} [delete]
func DeleteCars(c *gin.Context) {
	var db = database.GetDB()
	var carDelete models.Car

	err := db.First(&carDelete, "Id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&carDelete)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
