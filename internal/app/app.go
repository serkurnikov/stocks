//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock.$GOFILE Appl,Repo

// Package app provides business logic.
package app

import (
	"context"
	"github.com/Jeffail/gabs/v2"
	"stocks/internal/cryptocompareapi"
)

type (
	Ctx = context.Context

	Appl interface {
		GetCurrencyPrice(ctx Ctx) (*gabs.Container, error)
	}

	Repo interface {
		SavePriceCurrency(ctx Ctx, container *gabs.Container) (id int, err error)
	}

	ResourseData interface {
		GetCurrencyParamsFromYaml() (params *cryptocompareapi.CurrencyParams)
	}

	App struct {
		repo         Repo
		resourseData ResourseData
		alphaApi     cryptocompareapi.Api
	}
)

type (
	Currency struct {
		ID   int
		Name string
	}
)

func NewAppl(repo Repo, resourseData ResourseData, api cryptocompareapi.Api) Appl {
	return &App{
		repo:     repo,
		resourseData: resourseData,
		alphaApi: api,
	}
}
