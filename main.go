package main

import (
	"github.com/scubass/htmx-todo/db"
)

func main() {
	dbConn := db.SetupDb()
	defer dbConn.Close()
}
