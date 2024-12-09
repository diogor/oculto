-- name: GetPlayersForGame :many
SELECT * FROM player
WHERE game_id = $1;

-- name: GetGame :one
SELECT * FROM game
WHERE id = $1;

-- name: GetUnpickedPlayersForGame :many
SELECT * FROM player
WHERE game_id = $1 AND id <> $2 AND NOT is_picked;

-- name: GetPlayersHasNotPickedForGame :many
SELECT * FROM player
WHERE game_id = $1 AND NOT has_picked;

-- name: CreateGame :one
INSERT INTO game (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: CreatePlayer :one
INSERT INTO player (id, name, game_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CreatePlayers :copyfrom
INSERT INTO player (id, name, game_id)
VALUES ($1, $2, $3);

-- name: CreatePick :one
INSERT INTO pick (id, game_id, picked_by, player_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdatePicker :exec
UPDATE player SET has_picked = true
WHERE id = $1;

-- name: UpdatePicked :exec
UPDATE player SET is_picked = true
WHERE id = $1;

-- name: GetPicksForGameAndPlayer :many
SELECT * FROM pick
WHERE game_id = $1 AND player_id = $2;