package elastic

import (
	"context"
	"elastic/internal/models"
	"encoding/json"
	"fmt"
)


func (e *Elastic) InsertProduct(ctx context.Context, id string, jsonBody []byte) error {
	index := e.Indexes.Product
	if _, err := e.client.IndexExists(index).Do(ctx); err != nil {
		return fmt.Errorf("error checking index while inserting product %s", err.Error())
	}

	_, err := e.client.Index().Index(index).Id(id).BodyString(string(jsonBody)).Do(ctx)
	if err != nil {
		return fmt.Errorf("error inserting product %s", err.Error())
	}

	return nil
}

func (e *Elastic) UpdateProduct(ctx context.Context, prodID string, product models.Product) error {
	_, err := e.client.Update().
		Index(e.Indexes.Product).
		Id(prodID).
		Doc(product).
		Do(ctx)
	if err != nil {
		return fmt.Errorf("error during product update: %s", err.Error())
	}
	return nil
}

func (e *Elastic) DeleteProduct(ctx context.Context, prodID string) error {
	_, err := e.client.Delete().
		Index(e.Indexes.Product).
		Id(prodID).
		Do(ctx)
	if err != nil {
		return fmt.Errorf("error during product delete: %s", err.Error())
	}
	return nil
}

func (e *Elastic) FindManyProducts(ctx context.Context, searchString string) ([]models.Product, error) {
	limit := 30

	q := makeProductSearchQuery(searchString)

	searchResult, err := e.client.Search().
		Index(e.Indexes.Product).
		Query(q).
		Size(limit).
		Do(context.TODO())
	if err != nil {
		return []models.Product{}, err
	}

	if searchResult.Hits.TotalHits.Value != 0 {
		products := []models.Product{}
		for _, hit := range searchResult.Hits.Hits {
			var product models.Product
			err := json.Unmarshal(hit.Source, &product)
			if err != nil {
				return []models.Product{}, err
			}
			product.Name += fmt.Sprintf("%f", *hit.Score)
			products = append(products, product)
		}
		return products, nil
	}
	return []models.Product{}, nil 
}
