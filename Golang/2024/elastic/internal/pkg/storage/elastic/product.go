package elastic

import (
	"bytes"
	"context"
	"elastic/internal/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strconv"
	"strings"
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

func (e *Elastic) FindManyProducts(ctx context.Context, searchStr string) ([]models.Product, error) {
	query := fmt.Sprintf(`{"query": {"match": {"name": "%s"}}}`, searchStr)

	//data, err := json.Marshal(query)
	//if err != nil {
	//	return []models.Product{}, fmt.Errorf("%s : %w", ErrMarshal, err)
	//}

	req := esapi.SearchRequest{
		Index: []string{e.Indexes.Product + "*"},
		Body:  strings.NewReader(query),
	}

	res, err := req.Do(ctx, e.client)

	if err != nil {
		return []models.Product{}, fmt.Errorf("cannot search document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return []models.Product{}, fmt.Errorf("error in searching document response: %s", res.String())
	}

	fmt.Println(res)

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	var products []models.Product
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		// Преобразуем источник документа в JSON
		source, err := json.Marshal(hit.(map[string]interface{})["_source"])
		if err != nil {
			log.Fatalf("Error marshalling hit source: %s", err)
		}

		// Декодируем JSON в структуру Article
		var product models.Product
		if err := json.Unmarshal(source, &product); err != nil {
			log.Fatalf("Error unmarshalling hit source: %s", err)
		}

		// Добавляем документ в список статей
		products = append(products, product)
	}

	fmt.Printf("Found %d products\n", len(products))
	fmt.Println(products)
	return []models.Product{}, nil
}
