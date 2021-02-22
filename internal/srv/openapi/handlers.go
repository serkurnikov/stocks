package openapi

import (
	"stocks/api/openapi/restapi/op"
)

func (srv *server) PriceHandlerFunc(params op.PriceParams) op.PriceResponder {
	ctx, _ := fromRequest(params.HTTPRequest)

	result, _ := srv.app.GetCurrencyPrice(ctx)
	return op.NewPriceOK().WithPayload(&op.PriceOKBody{Result: result})
}
