package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mulfdev/pcola-crud-api/db"
)

func main() {
	e := echo.New()
	_, err := db.New("database.db")

	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/todo/:id", handleGetTodo)
	e.POST("/todo", handleAddTodo)
	e.DELETE("/todo/:id", handleDeleteTodo)
	e.PUT("/todo/:id", handleUpdateTodo)
	e.Logger.Fatal(e.Start(":1323"))
}

func handleDeleteTodo(c echo.Context) error {

	var todo db.Todo
	err := c.Bind(&todo)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "todo id missing")
	}

	result := db.DB().Delete(&todo, todo.Id)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not find todo")
	}

	return c.JSON(http.StatusOK, "deleted")
}

func handleAddTodo(c echo.Context) error {

	var todo db.Todo
	err := c.Bind(&todo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	result := db.DB().Create(&todo)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not create todo")
	}
	return c.JSON(http.StatusCreated, &todo)
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

	var todo db.Todo
	err := c.Bind(&todo)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not find todo")
	}

	result := db.DB().Where("id = ?", todo.Id).First(&todo)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not load todo")
	}

	return c.JSON(http.StatusOK, todo)
}

/*
POST to do list item /item
DEL to do list item /item/:id
PUT to do list item /item
GET to do list items for a specific group /item/:group_id
*/
