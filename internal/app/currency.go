package app

import (
	"github.com/Jeffail/gabs/v2"
	"stocks/internal/cryptocompareapi"
)

func (a App) GetCurrencyPrice(ctx Ctx, params *cryptocompareapi.CurrencyParams) (*gabs.Container, error) {
	return a.alphaApi.GetCurrencyPrice(params)
}

func (a App) SavePriceCurrency(ctx Ctx, name string) (*Currency, error) {
	id, err := a.repo.SavePriceCurrency(ctx, name)
	if err != nil {
		return nil, err
	}
	c := &Currency{
		ID:   id,
		Name: name,
	}
	return c, nil
}

func (a App) GetCurrencyParamsFromYaml(ctx Ctx) (params *cryptocompareapi.CurrencyParams) {
	return a.resourseData.GetCurrencyParamsFromYaml(ctx)
}