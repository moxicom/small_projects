package main

import (
	"flag"
	"fmt"
	"todo"
)

func main() {
	filename := "info.json"

	add := flag.String("add", "", "add a new todo")
	complete := flag.Int("complete", 0, "mark a task as completed")
	del := flag.Int("del", 0, "delete a row")
	get := flag.Bool("get", false, "get all rows")

	flag.Parse()

	tasks := &todo.Tasks{}

	if err := tasks.OpenTasks(filename); err != nil {
		fmt.Println(err)
		return
	}

	switch {
	case *add != "":
		tasks.Add(*add)
		err := tasks.UpdateJson(filename)
		if err != nil {
			fmt.Println(err)
		}
	case *complete > 0:
		err := tasks.Complete(*complete)
		if err != nil {
			fmt.Println(err)
		}
		err = tasks.UpdateJson(filename)
		if err != nil {
			fmt.Println(err)
		}
	case *del > 0:
		fmt.Println(*del)
		err := tasks.Remove(*del)
		if err != nil {
			fmt.Println(err)
		}
		err = tasks.UpdateJson(filename)
		if err != nil {
			fmt.Println(err)
		}
	case *get:
		tasks.Print()
	}
}
