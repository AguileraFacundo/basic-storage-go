package test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"

	db "github.com/AguileraFacundo/caja-simple/internal/db/sqlc"
)

var (
	TestQueries *db.Queries
	Conn        *sql.DB
)

const dbUrl = "postgres://root:mypw@localhost:5432/cajita?sslmode=disable"

func TestMain(m *testing.M) {
	Conn, err := pgx.Connect(context.Background(), dbUrl)
	fmt.Println(dbUrl)
	if err != nil {
		log.Fatalf("failed to connect database %v", err)
	}
	TestQueries = db.New(Conn)
	os.Exit(m.Run())
}
