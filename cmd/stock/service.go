package main

import (
	"context"
	"github.com/powerman/structlog"
	"stocks/api/openapi/restapi"
	"stocks/internal/app"
	"stocks/internal/config"
	"stocks/internal/cryptocompareapi"
	"stocks/internal/dal"
	"stocks/internal/srv/openapi"
	"stocks/pkg/concurrent"
	"stocks/pkg/serve"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

type service struct {
	cfg *config.ServeConfig
	srv *restapi.Server
}

func (s *service) runServe(ctxStartup, ctxShutdown Ctx, shutdown func()) (err error) {
	log := structlog.FromContext(ctxShutdown, nil)

	db, err := connectDB()
	if err != nil {
		return log.Err("err", err)
	}

	if err = migrationDB(db); err != nil {
		return log.Err("err", err)
	}

	cryptoApi := cryptocompareapi.NewCryptoCompare()
	go cryptoApi.UpdateCurrency()

	repo := dal.New(db)
	resourseData := dal.Init()
	appl := app.NewAppl(repo, resourseData, cryptoApi)
	s.srv, err = openapi.NewServer(appl)
	if err != nil {
		return log.Err("failed to openapi.NewServer", "err", err)
	}

	err = concurrent.Serve(ctxShutdown, shutdown,
		s.serveOpenAPI,
	)
	if err != nil {
		return log.Err("failed to serve", "err", err)
	}
	return nil
}

func (s *service) serveOpenAPI(ctx Ctx) error {
	return serve.OpenAPI(ctx, s.srv, "OpenAPI")
}
