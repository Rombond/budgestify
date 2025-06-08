package db_sql

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/Rombond/budgestify/api/types/recurrence"
)

func GetRecurrence(db *sql.DB, id int) (recurrence.Recurrence, error) {
	var recurrence recurrence.Recurrence
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE id = ?;", Tables[RecurrenceIdx].Key)
	err := db.QueryRow(query, id).Scan(&recurrence.Id, &recurrence.Name, &recurrence.House_User, &recurrence.PayerAccountID, &recurrence.Category, &recurrence.Amount, &recurrence.Currency, &recurrence.ConversionRate, &recurrence.PayDate, &recurrence.DayCycle)
	if err != nil {
		slog.Error("[GetRecurrence] Error querying recurrence: " + err.Error())
	}
	return recurrence, err
}

func GetRecurrencesByHouseUser(db *sql.DB, house_user int) ([]recurrence.Recurrence, error) {
	var recurrences []recurrence.Recurrence
	var recurrence recurrence.Recurrence
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE house_user = ?;", Tables[RecurrenceIdx].Key)
	rows, err := db.Query(query, house_user)
	if err != nil {
		slog.Error("[GetRecurrencesByHouseUser] Error querying recurrences: " + err.Error())
		return recurrences, err
	}
	for rows.Next() {
		rows.Scan(&recurrence.Id, &recurrence.Name, &recurrence.House_User, &recurrence.PayerAccountID, &recurrence.Category, &recurrence.Amount, &recurrence.Currency, &recurrence.ConversionRate, &recurrence.PayDate, &recurrence.DayCycle)
		recurrences = append(recurrences, recurrence)
	}
	if err = rows.Err(); err != nil {
		return recurrences, err
	}
	return recurrences, err
}

func GetRecurrencesByAccount(db *sql.DB, account int) ([]recurrence.Recurrence, error) {
	var recurrences []recurrence.Recurrence
	var recurrence recurrence.Recurrence
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE payer_account = ?;", Tables[RecurrenceIdx].Key)
	rows, err := db.Query(query, account)
	if err != nil {
		slog.Error("[GetRecurrencesByAccount] Error querying recurrences: " + err.Error())
		return recurrences, err
	}
	for rows.Next() {
		rows.Scan(&recurrence.Id, &recurrence.Name, &recurrence.House_User, &recurrence.PayerAccountID, &recurrence.Category, &recurrence.Amount, &recurrence.Currency, &recurrence.ConversionRate, &recurrence.PayDate, &recurrence.DayCycle)
		recurrences = append(recurrences, recurrence)
	}
	if err = rows.Err(); err != nil {
		return recurrences, err
	}
	return recurrences, err
}

func CreateRecurrence(db *sql.DB, name string, house_user int, payerAccount int, category *int, amount float64, currency string, conversionRate float64, payDate time.Time, dayCycle int) (int, error) {
	properties := "`name`, `house_user`, `payer_account`"
	placeholders := "?, ?, ?"
	args := []any{name, house_user, payerAccount}

	if category != nil {
		properties += ", `category`"
		placeholders += ", ?"
		args = append(args, *category)
	}

	properties += ", `amount`, `currency`, `conversion_rate`, `pay_date`, `day_cycle`"
	placeholders += ", ?, ?, ?, ?, ?"
	args = append(args, amount)
	args = append(args, currency)
	args = append(args, conversionRate)
	args = append(args, payDate)
	args = append(args, dayCycle)

	query := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%s);", Tables[RecurrenceIdx].Key, properties, placeholders)

	results, err := db.Exec(query, args...)
	if err != nil {
		slog.Error("[CreateRecurrence] Error while inserting new recurrence: " + err.Error())
		return -1, err
	}

	id, err := results.LastInsertId()
	if err != nil {
		slog.Error("[CreateRecurrence] Error while getting lastInsertId: " + err.Error())
		return -1, err
	}
	return int(id), nil
}

func ChangeRecurrence(db *sql.DB, id int, name string, house_user int, payerAccount int, category *int, amount float64, currency string, conversionRate float64, payDate time.Time, dayCycle int) (bool, error) {
	params := "`name` = ? `house_user` = ? `payer_account` = ? "
	args := []any{name, house_user, payerAccount}

	if category != nil {
		params += "`category` = ? "
		args = append(args, *category)
	}

	params += "`amount` = ? `currency` = ? `conversion_rate` = ? `pay_date` = ? `day_cycle` = ?"
	args = append(args, amount)
	args = append(args, currency)
	args = append(args, conversionRate)
	args = append(args, payDate)
	args = append(args, dayCycle)

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = ?", Tables[RecurrenceIdx].Key, params)
	_, err := db.Exec(query, args...)
	if err != nil {
		slog.Error("[ChangeRecurrence] Error while updating recurrence: " + err.Error())
		return false, err
	}
	return true, nil
}

func DeleteRecurrence(db *sql.DB, id int) (bool, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", Tables[RecurrenceIdx].Key)
	_, err := db.Exec(query, id)
	if err != nil {
		slog.Error("[DeleteRecurrence] Error while deleting recurrence: " + err.Error())
		return false, err
	}
	return true, nil
}
