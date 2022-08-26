package database

import (
	"os"

	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	connStr := os.Getenv("DB_HOST")
	opt, err := pg.ParseURL(connStr)
	if err != nil {
		panic(err) //we can just panic for now :)
	}

	db := pg.Connect(opt)
	if _, DBStatus := db.Exec("SELECT 1"); DBStatus != nil {
		panic("Postgres is down!!")
	}

	return db
}
