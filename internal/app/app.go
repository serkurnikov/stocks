//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock.$GOFILE Appl,Repo

// Package app provides business logic.
package app

import (
	"context"
	"stocks/internal/alphavantageapi"
)

type (
	Ctx = context.Context

	Appl interface{
		GetTimeSeriesIntraday(ctx Ctx, params alphavantageapi.TimeSeriesIntradayParams) (*alphavantageapi.TimeSeriesIntraday, error)
	}

	Repo interface{}

	App struct {
		repo     Repo
		alphaApi alphavantageapi.Api
	}
)

func NewAppl(repo Repo, api alphavantageapi.Api) Appl {
	return &App{
		repo:     repo,
		alphaApi: api,
	}
}
