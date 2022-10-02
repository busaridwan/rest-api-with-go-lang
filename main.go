package main

import (
	// "fmt"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Terminal - go mod init example/directory
// go get github.com/gin-gonic/gin

type Todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{
		ID:        "1",
		Item:      "Clean room",
		Completed: false,
	},
	{
		ID:        "2",
		Item:      "Read book",
		Completed: false,
	},
	{
		ID:        "3",
		Item:      "Record video",
		Completed: false,
	},
}

func getTodos(context *gin.Context) {
	// Covert todo array to json
	context.IndentedJSON(http.StatusOK, todos)

}
func addTodo(context *gin.Context) {
	// Create new todo
	var newTodo Todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	// Append the newTodo to the existing list
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)

}
func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found "})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}
func updateTodo(context *gin.Context) {

	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found "})
		return
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)

}
func getTodoById(id string) (*Todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func main() {

	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", updateTodo)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}
