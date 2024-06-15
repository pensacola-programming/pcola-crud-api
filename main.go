package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/mulfdev/pcola-crud-api/db"
)

/*

id - PK integer not null
email - varchar
password - varchar
first_name - varchar
last_name - varchar
*/

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

	return echo.NewHTTPError(http.StatusInternalServerError, "handler not implemented")
}

func handleUpdateTodo(c echo.Context) error {
	id := c.Param("id")

	if len(id) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "missing id param")
	}

	fmt.Println(id)
	panic("not implemented")
}

func handleGetTodo(c echo.Context) error {
	t := &Todo{
		Id:        uuid.New(),
		Title:     "todo title",
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
