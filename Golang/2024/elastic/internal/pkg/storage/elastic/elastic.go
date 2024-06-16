package elastic

import (
	"crypto/tls"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"net/http"
)

const aliasPostfix = "_alias"

type Indexes struct {
	Category string
	Product  string
}

type Elastic struct {
	client  *elasticsearch.Client
	Indexes Indexes
}

func New(addresses []string, indexes Indexes) (*Elastic, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Username: "elastic",
		Password: "Mxc*o5237LyzBqadjGkY",
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	res, err := client.Info()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error: %s", res.String())
	}

	e := &Elastic{
		client:  client,
		Indexes: indexes,
	}

	err = e.InitIndexes(indexes)
	if err != nil {
		return nil, fmt.Errorf("error initializing indexes: %s", err)
	}

	return e, nil
}

func (e *Elastic) createIndex(index string) error {
	alias := index + aliasPostfix

	res, err := e.client.Indices.Exists([]string{index})
	if err != nil {
		return fmt.Errorf("cannot check index existence: %w", err)
	}
	if res.StatusCode == 200 {
		return nil
	}
	if res.StatusCode != 404 {
		return fmt.Errorf("error in index existence response: %s", res.String())
	}

	res, err = e.client.Indices.Create(index)
	if err != nil {
		return fmt.Errorf("cannot create index: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("error in index creation response: %s", res.String())
	}

	res, err = e.client.Indices.PutAlias([]string{index}, alias)
	if err != nil {
		return fmt.Errorf("cannot create index alias: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("error in index alias creation response: %s", res.String())
	}

	return nil
}

// document represents a single document in Get API response body.
type document struct {
	Source interface{} `json:"_source"`
}

// InitIndexes Manually init every index
func (e *Elastic) InitIndexes(indexes Indexes) error {
	err := e.createIndex(indexes.Product)
	if err != nil {
		return err
	}

	err = e.createIndex(indexes.Category)
	if err != nil {
		return err
	}

	return nil
}
