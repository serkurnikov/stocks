package openapi

import (
	"github.com/go-openapi/swag"
	"stocks/api/openapi/restapi/op"
	"stocks/internal/cryptocompareapi"
)

func (srv *server) PriceHandlerFunc(params op.PriceParams) op.PriceResponder {
	ctx, _ := fromRequest(params.HTTPRequest)
	var prm = cryptocompareapi.CurrencyParams{}

	if params != op.NewPriceParams() {
		prm = cryptocompareapi.CurrencyParams{
			Fsyms: swag.StringValue(params.Fsyms),
			Tsyms: swag.StringValue(params.Tsyms),
		}
	} else {
		//Fetch from yaml
	}

	result, _ := srv.app.GetCurrencyPrice(ctx, &prm)
	return op.NewPriceOK().WithPayload(&op.PriceOKBody{Result: result})
}
