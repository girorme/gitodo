package main

import (
	"encoding/json"
	"gitodo/domain"
	"gitodo/repository"
	"io/ioutil"
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
	r.Use(gin.Recovery())
	r.Use(static.Serve("/", static.LocalFile("./static", false)))

	api := r.Group("/api")
	api.POST("/todos", saveTodos)
	api.GET("/todos", getTodos)

	r.Run()
}

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todoRepo.GetAll())
}

func saveTodos(c *gin.Context) {
	var todos domain.Todos

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := json.Unmarshal(body, &todos); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if _, err := todoRepo.SaveMany(todos); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, todos)
}
