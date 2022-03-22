package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewConnection(ctx context.Context) *pgxpool.Pool{
	addr := "postgres://postgres:secret@localhost:5432/test-api?sslmode=disable"
	conn, err := pgxpool.Connect(ctx,addr)
	if err != nil{
		log.Fatalf("Error connecting to db")
	}
	if err = conn.Ping(ctx); err != nil {
		log.Fatalf("Error ping database")
	}
	return conn
}
