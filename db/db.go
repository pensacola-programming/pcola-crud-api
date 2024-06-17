package db

import (
	"sync"

	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	once     sync.Once
	instance *gorm.DB
	err      error
)

type Todo struct {
	Id        uuid.UUID `json:"id" param:"id"`
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

func New(dbLocation string) (*gorm.DB, error) {
	once.Do(func() {
		instance, err = gorm.Open(sqlite.Open(dbLocation), &gorm.Config{})
		if err != nil {
			log.Fatalf("could not open db: %v", err)
		}

		if err = instance.AutoMigrate(&Todo{}, &User{}); err != nil {
			log.Fatalf("could not migrate db: %v", err)
		}
	})

	return instance, err
}

func DB() *gorm.DB {
	return instance
}
