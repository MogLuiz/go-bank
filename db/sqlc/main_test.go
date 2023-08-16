package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	DB_DRIVER = "postgres"
	DB_SOURCE = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var TEST_QUERIES *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(DB_DRIVER, DB_SOURCE)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer conn.Close()

	TEST_QUERIES = New(conn)

	os.Exit(m.Run())
}
