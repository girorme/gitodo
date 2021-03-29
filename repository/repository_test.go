package repository_test

import (
	"gitodo/domain"
	"gitodo/repository"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var todoRepo repository.Todo

func init() {
	database, err := gorm.Open(sqlite.Open("../database/todos-test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&domain.Todo{})
	database.Exec("DELETE FROM todos")
	todoRepo.Db = database
}

func TestGetAllTodos(t *testing.T) {
	todoRepo.NewTodo("Study tests golang")
	todos := todoRepo.GetAll()

	if len(todos) == 0 {
		t.Error("Error getting all todos")
	}
}

func TestNewTodo(t *testing.T) {
	todo := todoRepo.NewTodo("estudar tests golang")

	if todo.Title != "estudar tests golang" {
		t.Errorf("Error inserting todo")
	}
}

func TestSaveManyTodos(t *testing.T) {
	todos := []domain.Todo{
		{
			Title: "Todo 1",
		},
		{
			Title: "Todo 2",
		},
		{
			Title: "Todo 3",
		},
	}

	qtd, err := todoRepo.SaveMany(todos)

	if qtd != len(todos) {
		t.Errorf("Error saving many todos: %s", err)
	}
}

func TestGetOnlyOneTodo(t *testing.T) {
	todoRepo.NewTodo("only 1 todo")
	todos, err := todoRepo.GetOnly(1)

	if err != nil {
		t.Errorf("Error getting todos: %s", err)
	}

	if len(todos) != 1 {
		t.Errorf("Expect 1 todo, get %d", len(todos))
	}
}

func TestSaveManyAndDeletePrevious(t *testing.T) {
	todosToInsert := []domain.Todo{
		{
			Title: "Todo 1",
		},
		{
			Title: "Todo 2",
		},
		{
			Title: "Todo 3",
		},
	}

	qtd, err := todoRepo.SaveMany(todosToInsert)

	if qtd != len(todosToInsert) {
		t.Errorf("Error saving many todos: %s", err)
	}

	remainTodos, err := todoRepo.GetOnly(2)

	if err != nil {
		t.Errorf("Error getting todos: %s", err)
	}

	total, err := todoRepo.SaveMany(remainTodos)

	if total != len(remainTodos) {
		t.Errorf("Todos qty mismatch. Expect: %d, get: %d", len(remainTodos), total)
	}
}

func TestGetFirstTodo(t *testing.T) {
	expectedText := "Study tests golang"
	todoRepo.Db.Exec("DELETE FROM todos")
	todoRepo.NewTodo(expectedText)
	todo, err := todoRepo.First()

	if err != nil {
		t.Errorf("Error getting todos: %s", err)
	}

	if todo.Title != expectedText {
		t.Errorf("First todo mismatch. Expect: Study tests golang, get: %s", todo.Title)
	}
}

func TestGetOnlyWrong(t *testing.T) {
	todo, err := todoRepo.GetOnly(-1)

	if err != nil {
		t.Errorf("Error getting todos: %s", err)
	}

	if len(todo) != 1 {
		t.Errorf("Todo get only mismatch. Expect: 1, get: %d", len(todo))
	}
}
