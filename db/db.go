package db

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
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

func setupDb(dbLocation string) (*gorm.DB, error) {
	if len(dbLocation) == 0 {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "db file location required")
	}

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "could not open db")
	}

	return db, nil

}
