package main

import (
	// go modで記載しているmodule名/importするpackage名

	"fmt"
	"gin_todo/internal/infra/database"
)

func main() {
	dbConn := database.NewDB()
	defer fmt.Println("Successfully migrated")
	defer database.CloseDB(dbConn)
}

