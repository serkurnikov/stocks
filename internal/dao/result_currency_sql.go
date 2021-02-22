package dao

// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

func ScanResultCurrency(row *sql.Row) (*ResultCurrency, error) {
	var v0 int64
	var v1 string
	var v2 string

	err := row.Scan(
		&v0,
		&v1,
		&v2,
	)
	if err != nil {
		return nil, err
	}

	v := &ResultCurrency{}
	v.ID = v0
	v.Raw = v1
	v.Display = v2

	return v, nil
}

func ScanResultCurrencys(rows *sql.Rows) ([]*ResultCurrency, error) {
	var err error
	var vv []*ResultCurrency

	var v0 int64
	var v1 string
	var v2 string

	for rows.Next() {
		err = rows.Scan(
			&v0,
			&v1,
			&v2,
		)
		if err != nil {
			return vv, err
		}

		v := &ResultCurrency{}
		v.ID = v0
		v.Raw = v1
		v.Display = v2

		vv = append(vv, v)
	}
	return vv, rows.Err()
}

func SliceResultCurrency(v *ResultCurrency) []interface{} {
	var v0 int64
	var v1 string
	var v2 string

	v0 = v.ID
	v1 = v.Raw
	v2 = v.Display

	return []interface{}{
		v0,
		v1,
		v2,
	}
}

func SelectResultCurrency(db *sqlx.DB, query string, args ...interface{}) (*ResultCurrency, error) {
	row := db.QueryRow(query, args...)
	return ScanResultCurrency(row)
}

func SelectResultCurrencys(db *sqlx.DB, query string, args ...interface{}) ([]*ResultCurrency, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanResultCurrencys(rows)
}

func InsertResultCurrency(db *sqlx.DB, query string, v *ResultCurrency) error {

	res, err := db.Exec(query, SliceResultCurrency(v)[1:]...)
	if err != nil {
		return err
	}

	v.ID, err = res.LastInsertId()
	return err
}

func UpdateResultCurrency(db *sqlx.DB, query string, v *ResultCurrency) error {

	args := SliceResultCurrency(v)[1:]
	args = append(args, v.ID)
	_, err := db.Exec(query, args...)
	return err
}

const CreateResultCurrenyStmt = `
CREATE TABLE IF NOT EXISTS result_currenies (
 result_currency_id      SERIAL PRIMARY KEY 
,result_currency_raw     jsonb
,result_currency_display jsonb
);
`

const InsertResultCurrenyStmt = `
INSERT INTO result_currenies (
 result_currency_raw
,result_currency_display
) VALUES ($1,$2)
`

const SelectResultCurrenyStmt = `
SELECT 
 result_currency_id
,result_currency_raw
,result_currency_display
FROM result_currenies 
`

const SelectResultCurrenyRangeStmt = `
SELECT 
 result_currency_id
,result_currency_raw
,result_currency_display
FROM result_currenies 
LIMIT $1 OFFSET $2
`

const SelectResultCurrenyCountStmt = `
SELECT count(1)
FROM result_currenies 
`

const SelectResultCurrenyPkeyStmt = `
SELECT 
 result_currency_id
,result_currency_raw
,result_currency_display
FROM result_currenies 
WHERE result_currency_id=$1
`

const UpdateResultCurrenyPkeyStmt = `
UPDATE result_currenies SET 
 result_currency_id=$1
,result_currency_raw=$2
,result_currency_display=$3 
WHERE result_currency_id=$4
`

const DeleteResultCurrenyPkeyStmt = `
DELETE FROM result_currenies 
WHERE result_currency_id=$1
`
