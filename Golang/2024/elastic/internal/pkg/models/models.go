package models

type Product struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Price      float32   `json:"price"`
	Properties []Proprty `json:"properties"`
}

type Proprty struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
