package db_sql

import (
	"database/sql"
	"fmt"
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
	tableName := "Category"
	if isTableMissing(db, tableName) {
		createTable(db, tableName, "id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), name VARCHAR(255) NOT NULL, icons VARCHAR(255), parent INT, FOREIGN KEY (parent) REFERENCES Category(id)")
	}

	tableName = "User"
	if isTableMissing(db, tableName) {
		createTable(db, tableName, "id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), name VARCHAR(255) NOT NULL")
	}

	tableName = "AccountHouse"
	if isTableMissing(db, tableName) {
		createTable(db, tableName, "id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), name VARCHAR(255) NOT NULL, amount int NOT NULL")
	}

	tableName = "House"
	if isTableMissing(db, tableName) {
		createTable(db, tableName, "id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), name VARCHAR(255) NOT NULL, account INT, FOREIGN KEY (account) REFERENCES AccountHouse(id)")
	}

	tableName = "House_User"
	if isTableMissing(db, tableName) {
		createTable(db, tableName, "id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), house INT NOT NULL, user INT NOT NULL, FOREIGN KEY (house) REFERENCES House(id), FOREIGN KEY (user) REFERENCES User(id)")
	}

	tableName = "Account"
	if isTableMissing(db, tableName) {
		createTable(db, tableName, "id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), name VARCHAR(255) NOT NULL, amount int NOT NULL, house_user INT NOT NULL, FOREIGN KEY (house_user) REFERENCES House_User(id)")
	}

	tableName = "Transaction"
	if isTableMissing(db, tableName) {
		createTable(db, tableName, "id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), name VARCHAR(255) NOT NULL, amount int NOT NULL, payer INT NOT NULL, category INT, FOREIGN KEY (payer) REFERENCES House_User(id), FOREIGN KEY (category) REFERENCES Category(id)")
	}

	tableName = "Participant"
	if isTableMissing(db, tableName) {
		createTable(db, tableName, "id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), house_user INT NOT NULL, transaction int NOT NULL,FOREIGN KEY (house_user) REFERENCES House_User(id), FOREIGN KEY (transaction) REFERENCES Transaction(id)")
	}
}

func isTableMissing(db *sql.DB, wantedTable string) bool {
	var tableName string
	isMissing := true
	query := fmt.Sprintf("show tables like '%s'", wantedTable)
	results, err := db.Query(query)
	if err != nil {
		slog.Error("[isTableMissing] Error querying tables: " + err.Error())
		os.Exit(1)
	}
	defer results.Close()
	for results.Next() {
		err = results.Scan(&tableName)
		if err != nil {
			slog.Error("[isTableMissing] Error querying tables: " + err.Error())
			os.Exit(1)
		}
		if tableName == wantedTable {
			isMissing = false
			break
		}
	}
	if err = results.Err(); err != nil {
		slog.Error("[isTableMissing] Error on results: " + err.Error())
		os.Exit(1)
	}
	return isMissing
}

func createTable(db *sql.DB, tableName string, columns string) {
	slog.Info("[createTable] table " + tableName + " is missing creating it")
	query := fmt.Sprintf("CREATE TABLE %s (%s)", tableName, columns)
	_, err := db.Exec(query)
	if err != nil {
		slog.Error("[createTable] Error creating" + tableName + " table: " + err.Error())
		os.Exit(1)
	}
}
