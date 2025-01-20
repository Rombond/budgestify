package db_sql

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/go-sql-driver/mysql"
)

func ConnectDatabase() *sql.DB {
	cfg := mysql.Config{
		User:   "root",
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   "mysql_db:" + os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	return db
}

func InitDatabase(db *sql.DB) {
	if isTableMissing(db, "Category") {
		slog.Info("table 'Category''  is missing creating it")
		_, err := db.Exec("CREATE TABLE Category (id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL, icons VARCHAR(255), parent INT UNSIGNED, FOREIGN KEY (parent) REFERENCES Category(id) ON DELETE CASCADE")
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
	}
}

func isTableMissing(db *sql.DB, wantedTable string) bool {
	var tableName string
	isMissing := true
	results, err := db.Query("show tables like '?'", wantedTable)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer results.Close()
	for results.Next() {
		err = results.Scan(&tableName)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		if tableName == wantedTable {
			isMissing = false
			break
		}
	}
	if err = results.Err(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	return isMissing
}
