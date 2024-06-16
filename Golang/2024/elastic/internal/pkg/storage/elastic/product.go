package elastic

import (
	"bytes"
	"context"
	"elastic/internal/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"strconv"
)

func (e *Elastic) InsertProduct(ctx context.Context, prod models.Product) error {
	// Marshal the product to JSON
	data, err := json.Marshal(prod)
	if err != nil {
		return fmt.Errorf("%s : %w", ErrMarshal, err)
	}

	// Convert product ID to string
	id := strconv.Itoa(prod.ID)

	// Create the Elasticsearch request
	req := esapi.CreateRequest{
		Index:      e.Indexes.Product,
		DocumentID: id,
		Body:       bytes.NewReader(data),
	}

	// Execute the request
	res, err := req.Do(ctx, e.client)
	if err != nil {
		return fmt.Errorf("cannot index document: %w", err)
	}
	defer res.Body.Close()

	// Check if the response contains an error
	if res.IsError() {
		var errMsg map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&errMsg); err != nil {
			return fmt.Errorf("error decoding error response: %w", err)
		}
		return fmt.Errorf("error in indexing document response: %s", errMsg)
	}

	//log.Printf("Document indexed successfully with ID: %s", id)
	//log.Printf("Response: %s", res.String())
	return nil
}

func (e *Elastic) UpdateProduct(ctx context.Context, prod models.Product) error {
	update := map[string]interface{}{
		"doc": prod,
	}

	data, err := json.Marshal(update)
	if err != nil {
		return fmt.Errorf("%s : %w", ErrMarshal, err)
	}

	id := strconv.Itoa(prod.ID)

	req := esapi.UpdateRequest{
		Index:      e.Indexes.Product,
		DocumentID: id,
		Body:       bytes.NewReader(data),
	}

	res, err := req.Do(ctx, e.client)
	if err != nil {
		return fmt.Errorf("cannot update document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error in updating document response: %s", res.String())
	}

	return nil
}

func (e *Elastic) DeleteProduct(ctx context.Context, prodID int) error {
	id := strconv.Itoa(prodID)

	req := esapi.DeleteRequest{
		Index:      e.Indexes.Product,
		DocumentID: id,
		Refresh:    "true",
	}

	res, err := req.Do(ctx, e.client)
	if err != nil {
		return fmt.Errorf("cannot delete document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error in deleting document response: %s", res.String())
	}

	return nil
}

func (e *Elastic) FindManyProducts(ctx context.Context, prod models.Product) ([]models.Product, error) {
	return []models.Product{}, nil
}
