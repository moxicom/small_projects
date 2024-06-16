package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func (e *Elastic) IndexCategory(ctx context.Context, id string, category interface{}) error {
	data, err := json.Marshal(category)
	if err != nil {
		return fmt.Errorf("%s : %w", ErrMarshal, err)
	}

	req := esapi.IndexRequest{
		Index:      e.Indexes.Category,
		DocumentID: id,
		Body:       bytes.NewReader(data),
		Refresh:    "true",
	}

	res, err := req.Do(ctx, e.client)
	if err != nil {
		return fmt.Errorf("cannot index category: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("cannot index category: %s", res.Status())
	}
	return nil
}
