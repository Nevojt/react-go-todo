package controllers

import (
	"errors"
	"github.com/Nevojt/react-go-todo/backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
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

func GetTodoById(c *fiber.Ctx) error {
	//	Get parameter ID if URL
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}
	// Пошук Todo з заданим ID
	todo, err := models.GetTodoById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Todo not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Повернення Todo як JSON
	return c.JSON(todo)
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

func DeleteTodo(c *fiber.Ctx) error {
	//	Get parameter ID if URL
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}
	// Пошук Todo з заданим ID та видалення з бази даних
	err = models.DeleteTodo(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Todo not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Повернення Todo як JSON
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Todo deleted successfully",
		"id":      id,
		"todo":    nil, // Todo will be null since it was deleted from the database.
	})
}

func UpdateTodo(c *fiber.Ctx) error {
	//    Get parameter ID if URL
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}
	// Пошук Todo з заданим ID
	todo, err := models.GetTodoById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Todo not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Створення структури Todo для зберігання отриманих даних
	updateData := new(models.Todo)
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannon parse JSON",
		})
	}
	// Оновлення Todo в базі даних
	todo.Body = updateData.Body
	todo.Completed = updateData.Completed

	//	Збереження оновленого Todo в базі даних
	if err := models.UpdateTodo(todo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(todo)
}
