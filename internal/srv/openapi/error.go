package openapi

//go:generate gobin -m -run github.com/cheekybits/genny -in=$GOFILE -out=gen.$GOFILE gen "GetPrice=Currency"
//go:generate sed -i -e "\\,^//go:generate,d" gen.$GOFILE

import (
	_ "github.com/cheekybits/genny/generic"
	"github.com/go-openapi/swag"
	"net/http"
	"stocks/api/openapi/model"
	"stocks/api/openapi/restapi/op"
	"stocks/pkg/def"
)

func errGetPrice(log Log, err error, code errCode) op.PriceResponder {
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.extra, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.extra, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError { // Do no expose details about internal errors.
		msg = "internal error" //nolint:goconst // Duplicated by go:generate.
	}

	return op.NewPriceDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(code.extra),
		Message: swag.String(msg),
	})
}