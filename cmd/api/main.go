package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("GOBANK_DATABASE_USER"),
		os.Getenv("GOBANK_DATABASE_PASSWORD"),
		os.Getenv("GOBANK_DATABASE_HOST"),
		os.Getenv("GOBANK_DATABASE_PORT"),
		os.Getenv("GOBANK_DATABASE_NAME")))

	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

}
