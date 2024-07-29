package templates

const FiberTemplate = `package main
import (
    "log"

    "github.com/gofiber/fiber/v3"
)

func main() {
    // Initialize a new Fiber app
    app := fiber.New()

    app.Get("/", func(c fiber.Ctx) error {
        return c.SendString("Hello, Fiber 👋!")
    })

    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}
`
