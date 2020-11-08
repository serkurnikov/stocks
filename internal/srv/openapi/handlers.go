package openapi

import (
	"github.com/go-openapi/swag"
	"stocks/api/openapi/restapi/op"
	"stocks/internal/alphavantageapi"
)

func (srv *server) TimeSeriesIntradayHandlerFunc(params op.TimeSeriesIntradayParams) op.TimeSeriesIntradayResponder {
	ctx, _ := fromRequest(params.HTTPRequest)
	var prm  = alphavantageapi.TimeSeriesIntradayParams{
		Function:   swag.StringValue(params.Function),
		Symbol:     swag.StringValue(params.Symbol),
		Interval:   swag.StringValue(params.Interval),
		Adjusted:   "",
		Outputsize: "",
		Datatype:   "",
	}

	result, _ := srv.app.GetTimeSeriesIntraday(ctx, prm)
	return op.NewTimeSeriesIntradayOK().WithPayload(&op.TimeSeriesIntradayOKBody{Result: result})
}