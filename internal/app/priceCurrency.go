package app

import (
	"github.com/Jeffail/gabs/v2"
	"stocks/internal/cryptocompareapi"
)

func (a App) GetCurrencyPrice(ctx Ctx, params cryptocompareapi.CurrencyParams) (*gabs.Container, error) {
	return a.alphaApi.GetCurrencyPrice(params)
}

func (a *App) AddPriceCurrency(ctx Ctx, name string) (*PriceCurrency, error) {
	id, err := a.repo.AddPriceCurrency(ctx, name)
	if err != nil {
		return nil, err
	}
	c := &PriceCurrency{
		ID:   id,
		Name: name,
	}
	return c, nil
}