package main

import (
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
	Id        uuid.UUID
	Email     string
	Password  string
	FirstName string
	LastName  string
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func handleAddTodo(c echo.Context) error {}

/*
POST to do list item /item
DEL to do list item /item/:id
PUT to do list item /item
GET to do list items for a specific group /item/:group_id
*/
