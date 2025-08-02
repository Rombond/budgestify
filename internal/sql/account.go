package db_sql

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/Rombond/budgestify/api/types/account"
)

func GetAccount(db *sql.DB, id int) (account.Account, error) {
	var account account.Account
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE id = ?;", Tables[AccountIdx].Key)
	err := db.QueryRow(query, id).Scan(&account.Id, &account.Name, &account.House_User, &account.Amount, &account.Currency, &account.TheoreticalAmount)
	if err != nil {
		slog.Error("[GetAccount] Error querying account: " + err.Error())
	}
	return account, err
}

func GetAccounts(db *sql.DB, house_User int) ([]account.Account, error) {
	var accounts []account.Account
	var account account.Account
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE house_user = ?;", Tables[AccountIdx].Key)
	rows, err := db.Query(query, house_User)
	if err != nil {
		slog.Error("[GetAccounts] Error querying accounts: " + err.Error())
		return accounts, err
	}
	for rows.Next() {
		rows.Scan(&account.Id, &account.Name, &account.House_User, &account.Amount, &account.Currency, &account.TheoreticalAmount)
		accounts = append(accounts, account)
	}
	if err = rows.Err(); err != nil {
		return accounts, err
	}
	return accounts, err
}

func CreateAccount(db *sql.DB, name string, house_User int, amount float64, currency string, theoreticalAmount float64) (int, error) {
	query := fmt.Sprintf("INSERT INTO `%s` (`name`, `house_user`, `amount`, `currency`, `theoreticalAmount`) VALUES (?, ?, ?, ?, ?);", Tables[AccountIdx].Key)

	results, err := db.Exec(query, name, house_User, amount, currency, theoreticalAmount)
	if err != nil {
		slog.Error("[CreateAccount] Error while inserting new account: " + err.Error())
		return -1, err
	}

	id, err := results.LastInsertId()
	if err != nil {
		slog.Error("[CreateAccount] Error while getting lastInsertId: " + err.Error())
		return -1, err
	}
	return int(id), nil
}

func ChangeAccount(db *sql.DB, id int, name string, amount float64, theoreticalAmount float64) (bool, error) {
	query := fmt.Sprintf("UPDATE %s SET `name` = ?, `amount` = ?, `theoreticalAmount` = ? WHERE id = ?", Tables[AccountIdx].Key)
	_, err := db.Exec(query, name, amount, theoreticalAmount, id)
	if err != nil {
		slog.Error("[ChangeAccount] Error while updating account: " + err.Error())
		return false, err
	}
	return true, nil
}

func DeleteAccount(db *sql.DB, id int) (bool, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", Tables[AccountIdx].Key)
	_, err := db.Exec(query, id)
	if err != nil {
		slog.Error("[DeleteAccount] Error while deleting account: " + err.Error())
		return false, err
	}
	return true, nil
}
