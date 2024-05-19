package models

type Product struct {
	ID          string  `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string  `bson:"name" json:"name"`
	Description string  `bson:"description" json:"description"`
	Price       float64 `bson:"price" json:"price"`
	CategoryID  string  `bson:"category_id" json:"category_id"`
}

type Properties struct {
	PropName        string `bson:"prop_name,omitempty" json:"prop_name"`
	PropDescription string `bson:"prop_description" json:"prop_description"`
}
