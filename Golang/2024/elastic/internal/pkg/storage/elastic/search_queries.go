package elastic

import (
	"github.com/olivere/elastic/v7"
)

func makeProductSearchQuery(searchString string) *elastic.BoolQuery {
	return elastic.NewBoolQuery().Should(
		elastic.NewMultiMatchQuery("*" + searchString + "*", "name", "properties.title", "properties.desc").
			Fuzziness("AUTO"),
		elastic.NewNestedQuery("properties",
			elastic.NewQueryStringQuery("*"+ searchString + "*").
				Field("properties.desc").
				Field("properties.title").
				Fuzziness("AUTO"),
		).ScoreMode("avg"),
		elastic.NewNestedQuery("properties",
			elastic.NewQueryStringQuery(searchString).
				Field("properties.desc").
				Field("properties.title").
				Fuzziness("AUTO"),
		).ScoreMode("avg"),
	)
}