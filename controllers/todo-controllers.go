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
