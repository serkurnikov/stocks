package dao

//go:generate sqlgen -file currency.go -type ResultCurrency -pkg dao -o result_currency_sql.go -db postgres

type ResultCurrency struct {
	ID      int64  `db:"ID"`
	Raw     string `db:"Raw"`
	Display string `db:"Display"`
}
