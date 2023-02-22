package repositories

import (
	"database/sql"
	"golang-rest-api/structs"
)

func GetAllAccount(db *sql.DB) (results []structs.Account, err error) {
	sql := `SELECT * FROM account`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var account = structs.Account{}

		err = rows.Scan(&account.ID, &account.Name, &account.Balance)
		results = append(results, account)
	}
	if err != nil {
		panic(err)
	}
	return
}

func InsertAccount(db *sql.DB, account structs.Account) (err error) {
	sql := `INSERT INTO account (name, balance) VALUES ($1, $2)`
	errs := db.QueryRow(sql, account.Name, account.Balance)
	return errs.Err()
}
func UpdateAccount(db *sql.DB, account structs.Account) (err error) {
	sql := `UPDATE account SET name = $1, bakance = $2 WHERE id = $3`
	errs := db.QueryRow(sql, account.Name, account.Balance, account.ID)
	return errs.Err()
}
func DeleteAccount(db *sql.DB, account structs.Account) (err error) {
	sql := `DELETE FROM account WHERE id = $1`
	errs := db.QueryRow(sql, account.ID)
	return errs.Err()
}
