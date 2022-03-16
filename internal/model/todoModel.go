package model

import "gorm.io/gorm"

type ToDo struct {
	gorm.Model
	ToDo     string
	Finished bool
	Failed   bool
}

func GetTodos() []ToDo {
	var todos []ToDo
	return todos
}

func GetTodo() ToDo {
	var todo ToDo
	return todo
}
