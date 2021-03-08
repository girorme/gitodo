package repository

import (
	"gitodo/domain"

	"gorm.io/gorm"
)

// Todo repository communicates with databse
type Todo struct {
	Db *gorm.DB
}

// NewTodo create and return created row
func (t Todo) NewTodo(title string) domain.Todo {
	todo := domain.Todo{Title: title}
	t.Db.Create(&todo)
	return todo
}

// GetAll return all todos from database
func (t Todo) GetAll() domain.Todos {
	var todos domain.Todos
	t.Db.Find(&todos)
	return todos
}
