package controllers

import (
	"github.com/Nevojt/react-go-todo/models"
	"github.com/gofiber/fiber/v2"
)

func GetTodos(c *fiber.Ctx) error {
	todos, err := models.GetTodos() // Використовуйте функцію з models
	if err != nil {
		// Якщо є помилка при отриманні todos, поверніть JSON з помилкою
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Якщо все добре, поверніть todos як JSON
	return c.JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	// Створення структури Todo для зберігання отриманих даних
	todo := new(models.Todo)
	// Розбір JSON з тіла запиту до структури Todo
	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannon parse JSON",
		})
	}
	// Збереження нового Todo в базу даних
	createdTodo := models.CreateTodo(todo)
	if createdTodo == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create todo",
		})
	}
	return c.Status(fiber.StatusOK).JSON(createdTodo)
}
