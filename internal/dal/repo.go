package dal

import (
	"context"
	"stocks/internal/app"

	"github.com/jmoiron/sqlx"
)

type Ctx = context.Context

type Repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) app.Repo {
	return &Repo{db: db}
}

func (r Repo) SavePriceCurrency(_ app.Ctx, name string) (id int, err error) {
	panic("implement me")
}