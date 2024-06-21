package main

import (
	"context"
	"elastic/internal/pkg/storage/elastic"
	"elastic/internal/service"
)

// docker run --name kib01 --net somenetwork -p 5601:5601 kibana:8.14.1
// docker run -d --name elasticsearch --net somenetwork -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:8.14.1

func main() {
	// Todo init all storages Mxc*o5237LyzBqadjGkY

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

	// Todo Init service layer
	service := service.New(e)

	// err = service.InsertProduct(context.TODO(), models.Product{
	// 	ID:    1,
	// 	Name:  "Product name",
	// 	Price: 12,
	// 	Properties: []models.Property{
	// 		{
	// 			Title: "Desc",
	// 			Desc:  "Good zxczxczx",
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

	_, err = service.SearchProduct(context.Background(), "qwe")

	if err != nil {
		panic(err)
	}

	// Todo Init API layer

	// Todo run app
}
