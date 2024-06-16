package main

import (
	"context"
	"elastic/internal/pkg/storage/elastic"
	"elastic/internal/service"
)

func main() {
	// Todo init all storages Mxc*o5237LyzBqadjGkY

	e, err := elastic.New(
		[]string{"asdasd"},
		elastic.Indexes{
			Product:  "document",
			Category: "category",
		},
	)
	if err != nil {
		panic(err)
	}

	// Todo Init service layer
	service := service.New(e)
	//err = service.InsertProduct(context.TODO(), models.Product{
	//	ID:    1,
	//	Name:  "Product name",
	//	Price: 12,
	//	Properties: []models.Proprty{
	//		models.Proprty{
	//			Title: "Desc",
	//			Desc:  "Good zxczxczx",
	//		},
	//	},
	//})
	//err = service.UpdateProduct(context.TODO(), models.Product{
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
	err = service.DeleteProduct(context.TODO(), 1)
	if err != nil {
		panic(err)
	}
	// Todo Init API layer

	// Todo run app

}
