package repositories

import (
	"database/sql"
	"golang-rest-api/structs"
)

func GetAllTransaction(db *sql.DB) (results []structs.Record, err error) {
	// sql := `SELECT * FROM transaction`
	sql := `select
	t.id as id, 
  t.sender  as sender_id,
  a."name" as sender_name,
  t.beneficiary as beneficiary_id,
  a2."name" as beneficiary_name,
  t.amount as amount,
  t.remark as remark
FROM
  "transaction" t 
  INNER JOIN account a  ON t.sender  = a.id
  INNER JOIN account a2 ON t.beneficiary = a2.id;`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var record = structs.Record{}

		err = rows.Scan(&record.ID, &record.SenderID, &record.SenderName, &record.BeneficiaryID, &record.BeneficiaryName, &record.Remark, &record.Amount)
		results = append(results, record)
	}
	if err != nil {
		panic(err)
	}
	return
}

func InsertTransaction(db *sql.DB, transaction structs.Transaction) (err error) {
	sql := `INSERT INTO transaction (sender, beneficiary, amount, remark) VALUES ($1, $2, $3, $4)`
	errs := db.QueryRow(sql, &transaction.Sender, &transaction.Beneficiary, &transaction.Amount, &transaction.Remark)
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
