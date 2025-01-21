package main

import (
	"log/slog"
	"os"

	router "github.com/Rombond/budgestify/internal/router"
	sqlInit "github.com/Rombond/budgestify/internal/sql"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("Error while loading .env file: " + err.Error())
		os.Exit(1)
	}

	db := sqlInit.ConnectDatabase()
	sqlInit.InitDatabase(db)

	router.InitRouter(db)
}
