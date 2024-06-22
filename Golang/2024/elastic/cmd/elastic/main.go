package main

import (
	"context"
	"elastic/internal/pkg/storage/elastic"
	"elastic/internal/service"
	"fmt"
)

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

	// err = service.InsertProduct(context.TODO(), models.Product{
	// 	ID:    3,
	// 	Name:  "Products",
	// 	Price: 12,
	// 	Properties: []models.Property{
	// 		{
	// 			Title: "Описание",
	// 			Desc:  "Отличный огнетуш",
	// 		},
	// 		{
	// 			 Title: "Size",
	// 			 Desc: "123x12",
	// 		},
	// 	},
	// })

	// _, err = service.SearchProduct(context.TODO(), "qe")

	//service.UpdateProduct(context.TODO(), models.Product{
	//	ID:    1,
	//	Price: 122,
	//	Properties: []models.Proprty{
	//		models.Proprty{
	//			Title: "Descr",
	//			Desc:  "fixed",
	//		},
	//	},
	//	Name: "Fixed name",
	//})

	// service.DeleteProduct(context.TODO(), 1)

	products, err := service.SearchProduct(context.Background(), "пеный")
	
	// err = service.UpdateProduct(context.Background(),
	// models.Product{
	// 		ID:    3,
	// 		Name:  "Products Updated V2",
	// 		Price: 12,
	// 		Properties: []models.Property{
	// 			{
	// 				Title: "Описание",
	// 				Desc:  "Прекрасный пенный огнетушитель",
	// 			},
	// 			{
	// 				 Title: "Size",
	// 				 Desc: "123x12",
	// 			},
	// 		},
	// 	},
	// )
	if err != nil {
		panic(err)
	}

	for _, prod := range products {
		fmt.Println(prod)
	}

	// Todo Init API layer

	// Todo run app
}
