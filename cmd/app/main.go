package main

import (
	"context"
	"log"
	"os"

	"github.com/AguileraFacundo/caja-simple/internal/api"
	db "github.com/AguileraFacundo/caja-simple/internal/db/sqlc"
	"github.com/jackc/pgx/v5"
)

var (
	dbUrl         = os.Getenv("DB_URL")
	serverAddress = os.Getenv("SERVER_ADDRESS")
)

func main() {
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("failed to connect db %v", err)
	}

	db := db.New(conn)
	server := api.NewServer(db)

	server.Start(serverAddress)

}
