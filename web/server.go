package main

import (
	"github.com/a-h/templ"
	"github.com/diogor/oculto/web/templates"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func main() {
	app := fiber.New()
	app.Use(compress.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return Render(c, templates.Index())
	})

	app.Use(NotFoundMiddleware)
	log.Fatal(app.Listen(":3000"))
}

func NotFoundMiddleware(c *fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	return Render(c, templates.NotFound())
}

func Render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}
