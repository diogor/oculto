// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: copyfrom.go

package orm

import (
	"context"
)

// iteratorForCreatePlayers implements pgx.CopyFromSource.
type iteratorForCreatePlayers struct {
	rows                 []CreatePlayersParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreatePlayers) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreatePlayers) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].ID,
		r.rows[0].Name,
		r.rows[0].GameID,
	}, nil
}

func (r iteratorForCreatePlayers) Err() error {
	return nil
}

func (q *Queries) CreatePlayers(ctx context.Context, arg []CreatePlayersParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"player"}, []string{"id", "name", "game_id"}, &iteratorForCreatePlayers{rows: arg})
}
