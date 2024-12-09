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
	app.Post("/", web.CreateGameHandler)
	app.Get("/:game_id", web.GetGameHandler)
	app.Post("/pick", web.PickFriendHandler)

	app.Use(web.NotFoundMiddleware)
	log.Fatal(app.Listen(":3000"))
}
