package app

import (
	"github.com/Jeffail/gabs/v2"
)

func (a App) GetCurrencyPrice(ctx Ctx) (*gabs.Container, error) {
	prm := a.resourseData.GetCurrencyParamsFromYaml(ctx)
	var result *gabs.Container
	if prm != nil {
		result, _ = a.alphaApi.GetCurrencyPrice(prm)
	}

	a.repo.SavePriceCurrency(ctx, result)
	return result, nil
}