package db

import (
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func SetupDb() *sqlx.DB {
	connSting, found := os.LookupEnv("DB_URL")
	if !found {
		log.Fatalln("DB_URL enviroment variable not set")
	}
	db, err := sqlx.Open("pgx", connSting)
	if err != nil {
		log.Fatalf("unable to connect to database: %v\n", err)
	}

	return db
}
