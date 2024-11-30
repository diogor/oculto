// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package orm

import (
	"context"

	"github.com/google/uuid"
)

const createGame = `-- name: CreateGame :one
INSERT INTO game (id, name)
VALUES ($1, $2)
RETURNING id, name
`

type CreateGameParams struct {
	ID   uuid.UUID
	Name string
}

func (q *Queries) CreateGame(ctx context.Context, arg CreateGameParams) (Game, error) {
	row := q.db.QueryRow(ctx, createGame, arg.ID, arg.Name)
	var i Game
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const createPick = `-- name: CreatePick :one
INSERT INTO pick (id, game_id, picked_by, player_id)
VALUES ($1, $2, $3, $4)
RETURNING id, game_id, picked_by, player_id
`

type CreatePickParams struct {
	ID       uuid.UUID
	GameID   uuid.UUID
	PickedBy uuid.UUID
	PlayerID uuid.UUID
}

func (q *Queries) CreatePick(ctx context.Context, arg CreatePickParams) (Pick, error) {
	row := q.db.QueryRow(ctx, createPick,
		arg.ID,
		arg.GameID,
		arg.PickedBy,
		arg.PlayerID,
	)
	var i Pick
	err := row.Scan(
		&i.ID,
		&i.GameID,
		&i.PickedBy,
		&i.PlayerID,
	)
	return i, err
}

const createPlayer = `-- name: CreatePlayer :one
INSERT INTO player (id, name, game_id)
VALUES ($1, $2, $3)
RETURNING id, name, game_id, has_picked, is_picked
`

type CreatePlayerParams struct {
	ID     uuid.UUID
	Name   string
	GameID uuid.UUID
}

func (q *Queries) CreatePlayer(ctx context.Context, arg CreatePlayerParams) (Player, error) {
	row := q.db.QueryRow(ctx, createPlayer, arg.ID, arg.Name, arg.GameID)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.GameID,
		&i.HasPicked,
		&i.IsPicked,
	)
	return i, err
}

const getPlayersForGame = `-- name: GetPlayersForGame :many
SELECT id, name, game_id, has_picked, is_picked FROM player
WHERE game_id = $1
`

func (q *Queries) GetPlayersForGame(ctx context.Context, gameID uuid.UUID) ([]Player, error) {
	rows, err := q.db.Query(ctx, getPlayersForGame, gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Player
	for rows.Next() {
		var i Player
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.GameID,
			&i.HasPicked,
			&i.IsPicked,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUnpickedPlayersForGame = `-- name: GetUnpickedPlayersForGame :many
SELECT id, name, game_id, has_picked, is_picked FROM player
WHERE game_id = $1 AND NOT is_picked
`

func (q *Queries) GetUnpickedPlayersForGame(ctx context.Context, gameID uuid.UUID) ([]Player, error) {
	rows, err := q.db.Query(ctx, getUnpickedPlayersForGame, gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Player
	for rows.Next() {
		var i Player
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.GameID,
			&i.HasPicked,
			&i.IsPicked,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}