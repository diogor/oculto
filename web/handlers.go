package web

import (
	"context"
	"github.com/diogor/oculto/orm"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"strings"
)

func CreateGameHandler(c *fiber.Ctx) error {
	name := c.FormValue("name")
	players := c.FormValue("players")

	playerList := strings.Split(players, ",")

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

	var createPlayersParams []orm.CreatePlayersParams

	for _, player := range playerList {
		var playerID uuid.UUID
		playerID, err = uuid.NewV7()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		createPlayersParams = append(createPlayersParams, orm.CreatePlayersParams{
			ID:     playerID,
			Name:   player,
			GameID: game.ID,
		})
	}

	queries.CreatePlayers(c.Context(), createPlayersParams)

	return c.Redirect("/" + game.ID.String())
}
