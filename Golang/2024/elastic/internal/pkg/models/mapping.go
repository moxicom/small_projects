package models

const (
	CategoryMapping = `{
		"settings": {
			"analysis": {
				"filter": {
					"russian_stop": {
						"type": "stop",
						"stopwords": "_russian_"
					},
					"russian_stemmer": {
						"type": "stemmer",
						"language": "russian"
					}
				},
				"analyzer": {
					"ru": {
						"type": "custom",
						"tokenizer": "standard",
						"filter": [
							"lowercase",
							"russian_stop",
							"russian_stemmer"
						]
					}
				}
			}
		},
		"mappings": {
			"properties": {
				"id": { "type": "integer" },
				"name": { 
					"type": "text",
					"analyzer": "ru"
				}
			}
		}
	}`

	ProductMapping = `{
		"settings": {
			"analysis": {
				"filter": {
					"russian_stop": {
						"type": "stop",
						"stopwords": "_russian_"
					},
					"russian_stemmer": {
						"type": "stemmer",
						"language": "russian"
					}
				},
				"analyzer": {
					"ru": {
						"type": "custom",
						"tokenizer": "standard",
						"filter": [
							"lowercase",
							"russian_stop",
							"russian_stemmer"
						]
					}
				}
			}
		},
		"mappings": {
			"properties": {
				"id": { "type": "integer" },
				"category_id": { "type": "integer" },
				"name": { 
					"type": "text",
					"analyzer": "ru"
				},
				"price": { "type": "float" },
				"properties": {
					"type": "nested",
					"properties": {
						"title": { 
							"type": "text",
							"analyzer": "ru"
						},
						"desc": { 
							"type": "text",
							"analyzer": "ru"
						}
					}
				}
			}
		}
	}`
)
