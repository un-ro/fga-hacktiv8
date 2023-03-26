package main

import (
	"errors"
	"fmt"
	"web-gorm/database"
	"web-gorm/models"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()

	// createUser("nero@gmail.com")
	// getUserById(2)
	// updateUserById(2, "neronero@gmail.com")

	// createProduct(1, "Nvidia", "RTX 3050")
	getUsersWithProducts()
}

func createUser(email string) {
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user", err)
		return
	}

	fmt.Println("User created", User)
}

func getUserById(id uint) {
	db := database.GetDB()
	user := models.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User not found")
			return
		}
		fmt.Println("Error getting user", err)
	}

	fmt.Println("User found", user)
}

func updateUserById(id uint, email string) {
	db := database.GetDB()
	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error

	if err != nil {
		fmt.Println("Error updating user", err)
		return
	}

	fmt.Println("User updated", user)
}

func createProduct(userId uint, brand string, name string) {
	db := database.GetDB()

	Product := models.Product{
		UserID: userId,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error creating product", err.Error())
		return
	}

	fmt.Println("Product created:", Product)
}

func getUsersWithProducts() {
	db := database.GetDB()

	users := models.User{}
	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting users", err.Error())
		return
	}

	fmt.Println("User Datas with Products:")
	fmt.Println(users)
}

func deleteProductById(id uint) {
	db := database.GetDB()
	product := models.Product{}

	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error deleting product", err.Error())
		return
	}

	fmt.Println("Product deleted")
}
