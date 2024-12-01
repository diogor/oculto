package web

import (
	"context"
	"fmt"
	"github.com/diogor/oculto/orm"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func CreateGameHandler(c *fiber.Ctx) error {
	name := c.FormValue("name")
	fmt.Println(c.Context().PostArgs())
	var id uuid.UUID
	var err error
	var game orm.Game

	id, err = uuid.NewV7()

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	params := orm.CreateGameParams{
		ID:   id,
		Name: name,
	}

	conn := GetConnection()
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			panic(err)
		}
	}(conn, c.Context())

	queries := orm.New(conn)
	game, err = queries.CreateGame(c.Context(), params)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Redirect("/" + game.ID.String())
}
