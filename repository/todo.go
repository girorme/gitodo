package repository

import (
	"gitodo/domain"

	"gorm.io/gorm"
)

// Todo repository communicates with database
type Todo struct {
	Db *gorm.DB
}

// NewTodo create and return created row
func (t Todo) NewTodo(title string) (todo domain.Todo) {
	todo = domain.Todo{Title: title}
	t.Db.Create(&todo)
	return todo
}

// GetAll return all todos from database
func (t Todo) GetAll() (todos domain.Todos) {
	t.Db.Find(&todos)
	return todos
}

// SaveMany Saves batch of todos
func (t Todo) SaveMany(todos domain.Todos) (rowsAffected int, err error) {
	t.Db.Exec("DELETE FROM todos")

	if len(todos) == 0 {
		return 0, nil
	}

	res := t.Db.Create(todos)

	return int(res.RowsAffected), res.Error
}

func (t Todo) First() (todo domain.Todo, err error) {
	res := t.Db.First(&todo)
	return todo, res.Error
}

func (t Todo) GetOnly(qtd int) (todos domain.Todos, err error) {
	if qtd < 1 {
		qtd = 1
	}

	res := t.Db.Limit(qtd).Find(&todos)
	return todos, res.Error
}
