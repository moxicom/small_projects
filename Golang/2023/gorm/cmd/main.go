package main

import (
	"fmt"
	"gorm-lrn/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=314159 dbname=gorm_test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// MIGRATION
	db.AutoMigrate(&models.User{}, &models.CreditCard{})

	// CRUD
	// CREATE
	user := models.User{Name: "Alexey"}
	result := db.Create(&user)

	fmt.Println("ID:", user.ID)
	fmt.Println(result.Error)
	fmt.Println("Affected", result.RowsAffected) // Returns inserted records count

	fmt.Println("MULTIPLE")

	// Multiple records
	users := []*models.User{
		{Name: "Multiple1"},
		{Name: "Multiple2"},
	}

	result = db.Create(users)

	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)

	// READ
	fmt.Println("\nREAD")
	db.Where("id = ?", 5).Find(&user)
	// name = ? AND age >= ?
	// name IN ?, []string{}
	// .Or().Find()
	// ...
	fmt.Println(user)

	result = db.Where(&models.User{Name: "Alexey"}).First(&user)
	//.Find(&users)
	// if result != nil {
	// 	fmt.Println("No rows")
	// } else {
	// 	fmt.Println(user.ID)
	// }

	// UPDATE
	result = db.Where(&models.User{Name: "Alexey"}).First(&user)
	if result.Error == nil {
		user.Name = "ahhahaha, edited"
		db.Save(&user)
	}

	// DELETE
	result = db.Where(&models.User{Name: "Alexey"}).First(&user)
	db.Delete(&user)

}
