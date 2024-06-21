package models

const (
	CategoryMapping = `{
		"mappings": {
			"properties": {
				"id": { "type": "integer" },
				"name": { "type": "text" }
			}
		}
	}`

	ProductMapping = `{
		"mappings": {
			"properties": {
				"id": { "type": "integer" },
				"category_id": { "type": "integer" },
				"name": { "type": "text" },
				"price": { "type": "float" },
				"properties": {
					"type": "nested",
					"properties": {
						"title": { "type": "text" },
						"desc": { "type": "text" }
					}
				}
			}
		}
	}`
)