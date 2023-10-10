package main

import (
	"github.com/scubass/htmx-golang-crud/db"
)

func main() {
	dbConn := db.SetupDb()
	defer dbConn.Close()
}
