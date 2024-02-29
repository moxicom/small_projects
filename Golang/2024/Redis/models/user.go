package models

type User struct {
	Id   int    `redis:"id"`
	Name string `redis:"name"`
	Age  int    `redis:"age"`
}
