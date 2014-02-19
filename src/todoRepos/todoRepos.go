package todoRepos

import (
	"fmt"
)

var todos = map[int]Todo{123: Todo{123, "Bals Peru!"}}

func getTodos() []Todo {
	result := make([]Todo, len(todos))
	for _, todo := range todos {
		result = append(result, todo)
	}
	return result
}

func GetAll() []Todo {
	return getTodos()
}

func Get(id int) (Todo, error) {
	if todo, present := todos[id]; present {
		return todo, nil
	}
	return Todo{}, fmt.Errorf("todo %v not found", id)
}

func Insert(id int, text string) {
	todos[id] = Todo{id, text}
}

func Delete(id int) {
	if _, present := todos[id]; present {
		delete(todos, id)
	}
}

func Update(id int, text string) {
	if _, present := todos[id]; present {
		todo := todos[id]
		todo.Text = text
	}
}
