package elastic

import (
	"context"
	"elastic/internal/pkg/models"
	"fmt"
	"reflect"

	"github.com/olivere/elastic/v7"
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

// func (e *Elastic) UpdateProduct(ctx context.Context, prod models.Product) error {
// 	update := map[string]interface{}{
// 		"doc": prod,
// 	}

// 	data, err := json.Marshal(update)
// 	if err != nil {
// 		return fmt.Errorf("%s : %w", ErrMarshal, err)
// 	}

// 	id := strconv.Itoa(prod.ID)

// 	req := esapi.UpdateRequest{
// 		Index:      e.Indexes.Product,
// 		DocumentID: id,
// 		Body:       bytes.NewReader(data),
// 	}

// 	res, err := req.Do(ctx, e.client)
// 	if err != nil {
// 		return fmt.Errorf("cannot update document: %w", err)
// 	}
// 	defer res.Body.Close()

// 	if res.IsError() {
// 		return fmt.Errorf("error in updating document response: %s", res.String())
// 	}

// 	return nil
// }

// func (e *Elastic) DeleteProduct(ctx context.Context, prodID int) error {
// 	id := strconv.Itoa(prodID)

// 	req := esapi.DeleteRequest{
// 		Index:      e.Indexes.Product,
// 		DocumentID: id,
// 		Refresh:    "true",
// 	}

// 	res, err := req.Do(ctx, e.client)
// 	if err != nil {
// 		return fmt.Errorf("cannot delete document: %w", err)
// 	}
// 	defer res.Body.Close()

// 	if res.IsError() {
// 		return fmt.Errorf("error in deleting document response: %s", res.String())
// 	}

// 	return nil
// }

func (e *Elastic) FindManyProducts(ctx context.Context, searchStr string) ([]models.Product, error) {
	limit := 10
	esQueryTest := []string{"*"}
	for _, query := range esQueryTest {
		fmt.Println("query test:", query)
		searchResult, err := e.client.Search().
			Index(e.Indexes.Product).
			Query(elastic.NewQueryStringQuery(query)).
			Size(limit).
			Sort("last_attend", false).
			Do(context.TODO())
		if err != nil {
			return []models.Product{}, err
		}

		var ttyp map[string]interface{}
		for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
			if t, ok := item.(map[string]interface{}); ok {
				fmt.Println(t)
			}
		}
	}

	return []models.Product{}, nil 
}
