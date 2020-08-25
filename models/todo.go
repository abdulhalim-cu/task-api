package models

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title 			string 	`json:"title"`
	Description 	string 	`json:"description"`
	Done			bool 	`json:"done"`
}

func CreateTodo(title, description string) *Todo {
	return &Todo{
		Title: title,
		Description: description,
		Done: false,
	}
}


