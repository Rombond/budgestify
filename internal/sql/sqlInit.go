package db_sql

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Pair struct {
	Key   string
	Value string
}

var Tables = []Pair{
	{"User", "id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), name VARCHAR(255) NOT NULL, login VARCHAR(255) NOT NULL, hash BINARY(64) NOT NULL"},
	{"House", "id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), name VARCHAR(255) NOT NULL, account INT, FOREIGN KEY (account) REFERENCES AccountHouse(id)"},
	{"Category", "id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), name VARCHAR(255) NOT NULL, icons VARCHAR(255), parent INT, house INT, FOREIGN KEY (parent) REFERENCES Category(id), FOREIGN KEY (house) REFERENCES House(id)"},
	{"House_User", "id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), house INT NOT NULL, user INT NOT NULL, admin BOOLEAN NOT NULL, FOREIGN KEY (house) REFERENCES House(id), FOREIGN KEY (user) REFERENCES User(id)"},
	{"Account", "id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), name VARCHAR(255) NOT NULL, amount FLOAT NOT NULL, currency CHAR(3) NOT NULL, theoricalAmount FLOAT NOT NULL, house_user INT NOT NULL, FOREIGN KEY (house_user) REFERENCES House_User(id)"},
	{"Transaction", "id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), name VARCHAR(255) NOT NULL, amount FLOAT NOT NULL, payer INT NOT NULL, payer_account INT, house_account INT, category INT, pay_date DATE NOT NULL, currency CHAR(3) NOT NULL, conversion_rate FLOAT NOT NULL DEFAULT 1, FOREIGN KEY (payer) REFERENCES House_User(id), FOREIGN KEY (payer_account) REFERENCES Account(id), FOREIGN KEY (house_account) REFERENCES AccountHouse(id), FOREIGN KEY (category) REFERENCES Category(id)"},
	{"Participant", "id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), house_user INT NOT NULL, transaction INT NOT NULL, has_repay BOOLEAN NOT NULL, repay_date DATE, FOREIGN KEY (house_user) REFERENCES House_User(id), FOREIGN KEY (transaction) REFERENCES Transaction(id)"},
}

const (
	UserIdx        = 0
	HouseIdx       = 1
	CategoryIdx    = 2
	House_UserIdx  = 3
	AccountIdx     = 4
	TransactionIdx = 5
	ParticipantIdx = 6
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
		slog.Error("[ConnectDatabase] Error on Open: " + err.Error())
		os.Exit(1)
	}
	return db
}

func InitDatabase(db *sql.DB) {
	for _, pair := range Tables {
		if IsTableMissing(db, pair.Key) {
			createTable(db, pair.Key, pair.Value)
		}
	}
}

func IsTableMissing(db *sql.DB, wantedTable string) bool {
	var tableName string
	isMissing := true
	query := fmt.Sprintf("show tables like '%s'", wantedTable)
	results, err := db.Query(query)
	if err != nil {
		slog.Error("[IsTableMissing] Error querying tables: " + err.Error())
		os.Exit(1)
	}
	defer results.Close()
	for results.Next() {
		err = results.Scan(&tableName)
		if err != nil {
			slog.Error("[IsTableMissing] Error querying tables: " + err.Error())
			os.Exit(1)
		}
		if tableName == wantedTable {
			isMissing = false
			break
		}
	}
	if err = results.Err(); err != nil {
		slog.Error("[IsTableMissing] Error on results: " + err.Error())
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
