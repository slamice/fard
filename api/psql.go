package api

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/coopernurse/gorp"
)

// Postgres database environment settings.
const (
    ConnString = "user=postgres host=localhost dbname=frd"
)

fun connectDb() {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open(ConnString)

	if err != nil {
		log.Fatal("Postgres connection failed: ", err)
	}

	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

}

func Close() {
	tableMaps = []*gorp.TableMap{}
	dbmap.Db.Close()
}