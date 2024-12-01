package web

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"os"
)

func GetConnection() *pgx.Conn {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	return conn
}
