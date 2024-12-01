package main

import (
	"github.com/diogor/oculto/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func main() {
	app := fiber.New()
	app.Use(compress.New())

	app.Get("/", web.IndexHandler)
	app.Get("/:id", web.GetGameHandler)
	app.Post("/game", web.CreateGameHandler)

	app.Use(web.NotFoundMiddleware)
	log.Fatal(app.Listen(":3000"))
}
