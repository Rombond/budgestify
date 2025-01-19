package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cfg := mysql.Config{
		User:   "root",
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   "mysql_db:" + os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
	}

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		state := true
		pingErr := db.Ping()
		if pingErr != nil {
			log.Fatal(pingErr)
			state = false
		}
		ctx.JSON(http.StatusOK, gin.H{"status": state})
	})
	router.Run(":" + os.Getenv("API_PORT"))
}
