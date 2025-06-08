package db_sql

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/Rombond/budgestify/api/types/transaction"
)

func GetTransaction(db *sql.DB, id int) (transaction.Transaction, error) {
	var transaction transaction.Transaction
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE id = ?;", Tables[TransactionIdx].Key)
	err := db.QueryRow(query, id).Scan(&transaction.Id, &transaction.Name, &transaction.CategoryID, &transaction.PayerID, &transaction.PayerAccountID, &transaction.PayDate, &transaction.Currency, &transaction.ConversionRate)
	if err != nil {
		slog.Error("[GetTransaction] Error querying transaction: " + err.Error())
	}
	return transaction, err
}

func GetTransactionsByHouseUser(db *sql.DB, house_user int) ([]transaction.Transaction, error) {
	var transactions []transaction.Transaction
	var transaction transaction.Transaction
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE payer = ?;", Tables[TransactionIdx].Key)
	rows, err := db.Query(query, house_user)
	if err != nil {
		slog.Error("[GetTransactionsByHouseUser] Error querying transactions: " + err.Error())
		return transactions, err
	}
	for rows.Next() {
		rows.Scan(&transaction.Id, &transaction.Name, &transaction.CategoryID, &transaction.PayerID, &transaction.PayerAccountID, &transaction.PayDate, &transaction.Currency, &transaction.ConversionRate)
		transactions = append(transactions, transaction)
	}
	if err = rows.Err(); err != nil {
		return transactions, err
	}
	return transactions, err
}

func GetTransactionsByAccount(db *sql.DB, account int) ([]transaction.Transaction, error) {
	var transactions []transaction.Transaction
	var transaction transaction.Transaction
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE payer_account = ?;", Tables[TransactionIdx].Key)
	rows, err := db.Query(query, account)
	if err != nil {
		slog.Error("[GetTransactionsByAccount] Error querying transactions: " + err.Error())
		return transactions, err
	}
	for rows.Next() {
		rows.Scan(&transaction.Id, &transaction.Name, &transaction.CategoryID, &transaction.PayerID, &transaction.PayerAccountID, &transaction.PayDate, &transaction.Currency, &transaction.ConversionRate)
		transactions = append(transactions, transaction)
	}
	if err = rows.Err(); err != nil {
		return transactions, err
	}
	return transactions, err
}

func CreateTransaction(db *sql.DB, name string, category *int, amount float64, payer int, payerAccount *int, payDate time.Time, currency string, conversionRate float64) (int, error) {
	properties := "`name`"
	placeholders := "?"
	args := []any{name}

	if category != nil {
		properties += ", `category`"
		placeholders += ", ?"
		args = append(args, *category)
	}

	properties += ", `amount`, `payer`"
	placeholders += ", ?, ?"
	args = append(args, amount)
	args = append(args, payer)

	if payerAccount != nil {
		properties += ", `payer_account`"
		placeholders += ", ?"
		args = append(args, *payerAccount)
	}

	properties += ", `pay_date`, `currency`, `conversion_rate`"
	placeholders += ", ?, ?, ?"
	args = append(args, payDate)
	args = append(args, currency)
	args = append(args, conversionRate)

	query := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%s);", Tables[TransactionIdx].Key, properties, placeholders)

	results, err := db.Exec(query, args...)
	if err != nil {
		slog.Error("[CreateTransaction] Error while inserting new transaction: " + err.Error())
		return -1, err
	}

	id, err := results.LastInsertId()
	if err != nil {
		slog.Error("[CreateTransaction] Error while getting lastInsertId: " + err.Error())
		return -1, err
	}
	return int(id), nil
}

func ChangeTransaction(db *sql.DB, id int, name string, category *int, amount float64, payer int, payerAccount *int, payDate time.Time, currency string, conversionRate float64) (bool, error) {
	params := "`name` = ? "
	args := []any{}

	if category != nil {
		params += "`category` = ? "
		args = append(args, *category)
	}

	params += "`amount` = ? `payer` = ? "
	args = append(args, amount)
	args = append(args, payer)

	if payerAccount != nil {
		params += "`payer_account` = ? "
		args = append(args, *payerAccount)
	}

	params += "`pay_date`= ? `currency` = ? `conversion_rate` = ?"
	args = append(args, payDate)
	args = append(args, currency)
	args = append(args, conversionRate)
	args = append(args, id)

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = ?", Tables[TransactionIdx].Key, params)
	_, err := db.Exec(query, args...)
	if err != nil {
		slog.Error("[ChangeTransaction] Error while updating transaction: " + err.Error())
		return false, err
	}
	return true, nil
}

func DeleteTransaction(db *sql.DB, id int) (bool, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", Tables[TransactionIdx].Key)
	_, err := db.Exec(query, id)
	if err != nil {
		slog.Error("[DeleteTransaction] Error while deleting transaction: " + err.Error())
		return false, err
	}
	return true, nil
}
