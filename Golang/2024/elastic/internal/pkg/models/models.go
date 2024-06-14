package models

type Product struct {
	ID         int
	Name       string
	Price      float32
	Properties []Proprty
}

type Proprty struct {
	Title string
	Desc  string
}
