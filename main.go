package main

import (
	"log/slog"
	"os"

	router "github.com/Rombond/budgestify/internal/router"
	db_sql "github.com/Rombond/budgestify/internal/sql"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("Error while loading .env file: " + err.Error())
		os.Exit(1)
	}

	db := db_sql.ConnectDatabase(os.Getenv("DB_NAME"))
	db_sql.UpdateStateSetup(db, -1)
	state := db_sql.GetStateSetup()
	if !state.IsDbInitialized {
		db_sql.InitDatabase(db)
	}

	router.InitRouter(db)
}
