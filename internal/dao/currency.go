package dao

//go:generate sqlxgen -file currency.go -type ResultCurrency -pkg dao -o result_currency_sql.go -db postgres

var CurrencyValuesConstants = map[string]string{
	"CHANGE24HOUR":    "CHANGE24HOUR",
	"CHANGEPCT24HOUR": "CHANGEPCT24HOUR",
	"OPEN24HOUR":      "OPEN24HOUR",
	"VOLUME24HOUR":    "VOLUME24HOUR",
	"VOLUME24HOURTO":  "VOLUME24HOURTO",
	"LOW24HOUR":       "LOW24HOUR",
	"HIGH24HOUR":      "HIGH24HOUR",
	"PRICE":           "PRICE",
	"SUPPLY":          "SUPPLY",
	"MKTCAP":          "MKTCAP",
}

type ResultCurrency struct {
	ID      int64  `db:"ID"`
	Raw     string `db:"Raw"`
	Display string `db:"Display"`
}
