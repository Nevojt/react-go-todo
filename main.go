package react_go_todo

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	fmt.Println("Hello, World!")
	app := fiber.New()

	log.Fatal(app.Listen(":4000"))
}
