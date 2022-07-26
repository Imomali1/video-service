package main

import (
	"api/pkg/db"
	"api/pkg/utils/reader"
)

func main() {
	db.ConnDB()

	initTable := reader.ReadFile("pkg/db/SQL/table.sql")

	_, err := db.MyDB.Exec(string(initTable))
	if err != nil {
		panic(err)
	}

	defer db.CloseDB(db.MyDB)
}
