package main

import (
	"gitodo/domain"
	"gitodo/repository"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var todoRepo repository.Todo

func init() {
	database, err := gorm.Open(sqlite.Open("database/todos.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&domain.Todo{})
	todoRepo.Db = database
}

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./static", false)))
	r.GET("/todos", getAllTodos)
	r.Run()
}

func getAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": todoRepo.GetAll()})
}
