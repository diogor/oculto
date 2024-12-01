package web

import (
	"context"
	"github.com/a-h/templ"
	"github.com/diogor/oculto/orm"
	"github.com/diogor/oculto/web/templates"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"strings"
)

func render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}

func NotFoundMiddleware(c *fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	return render(c, templates.NotFound())
}

func IndexHandler(c *fiber.Ctx) error {
	return render(c, templates.Index())
}

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

func GetGameHandler(c *fiber.Ctx) error {
	var id uuid.UUID
	var err error
	var game orm.Game

	id, err = uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	conn := GetConnection()
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			panic(err)
		}
	}(conn, context.Background())

	queries := orm.New(conn)
	game, err = queries.GetGame(c.Context(), id)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var players []orm.Player
	players, err = queries.GetPlayersForGame(c.Context(), id)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return render(c, templates.Game(game, players))
}
