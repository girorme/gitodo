package test

import (
	"gitodo/domain"
	"gitodo/repository"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	todoRepo repository.Todo
)

func init() {
	database, err := gorm.Open(sqlite.Open("./todos-test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&domain.Todo{})
	database.Create(&domain.Todo{Description: "totest"})
	todoRepo.Db = database
}

func TestGetAllTodos(t *testing.T) {
	todos := todoRepo.GetAll()

	if len(todos) == 0 {
		t.Error("Error getting all todos")
	}
}

func TestNewTodo(t *testing.T) {
	todo := todoRepo.NewTodo("estudar tests golang")

	if todo.Description != "estudar tests golang" {
		t.Errorf("Error inserting todo")
	}
}
