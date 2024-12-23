package models

import (
	"log"

	"github.com/Nevojt/react-go-todo/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	config.Connect()
	db = config.GetDB()
	err = db.AutoMigrate(&Todo{})
	if err != nil {
		log.Fatalf("Failed to automigrate: %v", err)
	}
}

type Todo struct {
	gorm.Model
	ID        int    `gorm:"primaryKey" json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func CreateTodo(t *Todo) *Todo {
	result := db.Create(t)
	if result.Error != nil {
		log.Printf("Error creating todo: %v", result.Error)
		return nil
	}
	return t
}

func GetTodos() ([]Todo, error) {
	var todos []Todo
	result := db.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}
