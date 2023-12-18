package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=postgres port=9920 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
