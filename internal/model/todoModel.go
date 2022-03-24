package model

import "gorm.io/gorm"

type ToDo struct {
	gorm.Model
	ToDo     string `json:"todo"`
	Finished bool   `json:"finished"`
	Failed   bool   `json:"failed"`
}

func GetTodos() *[]ToDo {
	var todos []ToDo
	return &todos
}

func GetTodo() *ToDo {
	var todo ToDo
	return &todo
}
