package elastic

import (
	"context"
	"elastic/internal/models"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
)

type Indexes struct {
	Category string
	Product  string
}

type Elastic struct {
	client  *elastic.Client
	Indexes Indexes
}

func New(ctx context.Context, URL string, indexes Indexes) (*Elastic, error) {
	client, err := elastic.NewClient(
		elastic.SetBasicAuth(
			"elastic",
			"changeme",
		),
		elastic.SetURL(URL),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating Elasticsearch client: %w", err)
	}

	_, _, err = client.Ping(URL).Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting Elasticsearch info: %w", err)
	}

	e := &Elastic{
		client:  client,
		Indexes: indexes,
	}

	err = e.initIndexes(ctx, indexes)
	if err != nil {
		return nil, fmt.Errorf("error initializing indexes: %w", err)
	}

	return e, nil
}

func (e *Elastic) createIndexIfNotExists(ctx context.Context, index string, mapping string) error {
	log.Printf("Checking if index %s exists...\n", index)
	exists, err := e.client.IndexExists(index).Do(ctx)
	if err != nil {
		return fmt.Errorf("error checking if index %s exists: %s", index, err.Error())
	}

	if exists {
		log.Printf("Index %s already exists\n", index)
		return nil
	}

	log.Printf("Index %s does not exist. Creating index...\n", index)
	_, err = e.client.CreateIndex(index).BodyString(mapping).Do(ctx)
	if err != nil {
		return fmt.Errorf("error creating index %s: %s", index, err.Error())
	}

	log.Printf("Index %s created successfully\n", index)
	return nil
}

// InitIndexes Manually init every index
func (e *Elastic) initIndexes(ctx context.Context, indexes Indexes) error {
	err := e.createIndexIfNotExists(ctx, indexes.Product, models.ProductMapping)
	if err != nil {
		return err
	}

	err = e.createIndexIfNotExists(ctx, indexes.Category, models.CategoryMapping)
	if err != nil {
		return err
	}

	return nil
}
