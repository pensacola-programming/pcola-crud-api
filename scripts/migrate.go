package main

import (
	"fmt"
	"github.com/mulfdev/pcola-crud-api/db"
	"log"
)

func main() {
	_, err := db.New("database.db")

	if err != nil {
		log.Fatal("couldnt migrate db")
	}
	fmt.Println("Migrating DB")

	db.DB().AutoMigrate(&db.Todo{})

	fmt.Println("Migration complete")
}
