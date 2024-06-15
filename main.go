package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

/*

id - PK integer not null
email - varchar
password - varchar
first_name - varchar
last_name - varchar
*/

type Todo struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Created   time.Time `json:"created"`
	Due       time.Time `json:"due"`
	Completed bool      `json:"completed"`
	GroupId   int       `json:"groupId"`
}

type User struct {
	Id        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func assert(condition bool, errMsg string) error {
	if !condition {
		return fmt.Errorf(errMsg)
	}

	return nil
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/todo", handleGetTodo)
	e.POST("/todo", handleAddTodo)

	e.PUT("/todo/:id", handleUpdateTodo)
	e.Logger.Fatal(e.Start(":1323"))
}

func handleDeleteTodo(c echo.Context) error {
	panic("Not implemented")
}

func handleAddTodo(c echo.Context) error {
	id := c.Param("id")
	assert(len(id) > 0, "id missing")
	panic("not implemented")
}

func handleUpdateTodo(c echo.Context) error {
	panic("not implemented")
}

func handleGetTodo(c echo.Context) error {
	t := &Todo{
		Id:        uuid.New(),
		Title:     "todo title",
		Created:   time.Now(),
		Due:       time.Now().AddDate(0, 0, 10),
		Completed: false,
		GroupId:   1,
	}
	return c.JSON(http.StatusOK, t)
}

/*
POST to do list item /item
DEL to do list item /item/:id
PUT to do list item /item
GET to do list items for a specific group /item/:group_id
*/
