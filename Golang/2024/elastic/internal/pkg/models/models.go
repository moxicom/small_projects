package models

type Category struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	CategoryID int 		 `json:"category_id"`
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Price      float32   `json:"price"`
	Properties []Property `json:"properties"`
}

type Property struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
