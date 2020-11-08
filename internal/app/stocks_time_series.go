package app

import "stocks/internal/alphavantageapi"

func (a App) GetTimeSeriesIntraday(ctx Ctx, params alphavantageapi.TimeSeriesIntradayParams) (*alphavantageapi.TimeSeriesIntraday, error) {
	return a.alphaApi.GetTimeSeriesIntraday(params)
}
