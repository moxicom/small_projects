package elastic

import (
	"github.com/olivere/elastic/v7"
)

func makeProductSearchQuery(searchString string) *elastic.MatchQuery {
	return elastic.NewMatchQuery("name", "*"+searchString+"*").Analyzer("product_name_analyzer").Fuzziness("auto").Boost(2.0)
}