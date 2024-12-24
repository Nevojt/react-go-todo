package main

import (
	"github.com/Nevojt/react-go-todo/backend/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading.env file")
	}

	app.Get("/api/todos", controllers.GetTodos)
	app.Get("/api/todos/:id", controllers.GetTodoById)
	app.Post("/api/todos", controllers.CreateTodo)
	app.Patch("/api/todos/:id", controllers.UpdateTodo)
	app.Delete("/api/todos/:id", controllers.DeleteTodo)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "5000"
	}
	log.Fatal(app.Listen("0.0.0.0:" + PORT))
}
