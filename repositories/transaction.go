package repositories

import (
	"database/sql"
	"golang-rest-api/structs"
)

func GetAllTransaction(db *sql.DB) (results []structs.Transaction, err error) {
	sql := `SELECT * FROM transaction`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var transaction = structs.Transaction{}

		err = rows.Scan(&transaction.ID, &transaction.Sender, &transaction.Beneficiary, &transaction.Amount, &transaction.Remark)
		results = append(results, transaction)
	}
	if err != nil {
		panic(err)
	}
	return
}

func InsertTransaction(db *sql.DB, transaction structs.Transaction) (err error) {
	sql := `INSERT INTO transaction (id, first_name, last_name) VALUES ($1, $2, $3,$4,$5)`
	errs := db.QueryRow(sql, transaction.ID, &transaction.Sender, &transaction.Beneficiary, &transaction.Amount, &transaction.Remark)
	return errs.Err()
}
func UpdateTransaction(db *sql.DB, transaction structs.Transaction) (err error) {
	sql := `UPDATE transaction SET first_name = $1, last_name = $2 WHERE id = $3`
	errs := db.QueryRow(sql, &transaction.Sender, &transaction.Beneficiary, &transaction.Amount, &transaction.Remark, transaction.ID)
	return errs.Err()
}
func DeleteTransaction(db *sql.DB, transaction structs.Transaction) (err error) {
	sql := `DELETE FROM transaction WHERE id = $1`
	errs := db.QueryRow(sql, transaction.ID)
	return errs.Err()
}
