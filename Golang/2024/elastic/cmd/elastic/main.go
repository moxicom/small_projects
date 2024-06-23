package main

import (
	"context"
	"elastic/internal/handlers"
	"elastic/internal/models"
	"elastic/internal/server"
	"elastic/internal/service"
	"elastic/internal/storage/elastic"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// using https://github.com/nickyat/elasticsearch-analysis-morphology/releases

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	e, err := elastic.New(
		context.Background(),
		"http://localhost:9200",
		elastic.Indexes{
			Product:  "product",
			Category: "category",
		},
	)
	if err != nil {
		panic(err)
	}

	service := service.New(e)

	// TEST INPUT DATA
	jsonFile, err := os.Open("products.json")
	check(err)
	defer jsonFile.Close()

	byteVal, _ := io.ReadAll(jsonFile)

	var data models.InputData
	err = json.Unmarshal(byteVal, &data)
	if err != nil {
		fmt.Println("Ошибка при распаковке JSON:", err)
		return
	}
	prodID := 0
	for categoryID, val := range data.Data {
		for _, product := range val.Category.Products {
			prod := models.Product {
				ID: prodID,
				CategoryID: categoryID,
				Name: product.ProdName,
			}
			prodID++
			service.InsertProduct(context.TODO(), prod)
		}
	}

	//

	handlers := handlers.New(service)

	server := server.NewServer()

	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		panic(err)
	}
}
