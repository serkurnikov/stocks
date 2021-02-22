package dal

import (
	"context"
	"github.com/Jeffail/gabs/v2"
	"github.com/jmoiron/sqlx"
	"stocks/internal/app"
	"stocks/internal/cryptocompareapi"
	"stocks/internal/dao"
)

type Ctx = context.Context

type Repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) app.Repo {
	return &Repo{db: db}
}

func (r Repo) SavePriceCurrency(ctx app.Ctx, container *gabs.Container) (id int, err error) {
	var v = dao.ResultCurrency{
		Raw: container.S(cryptocompareapi.RAW).String(),
		Display: container.S(cryptocompareapi.DISPLAY).String(),
	}
	dao.InsertResultCurrency(r.db, dao.InsertResultCurrenyStmt, &v)
	return 0, err
}
