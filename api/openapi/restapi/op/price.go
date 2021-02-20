// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PriceHandlerFunc turns a function with the right signature into a price handler
type PriceHandlerFunc func(PriceParams) PriceResponder

// Handle executing the request and returning a response
func (fn PriceHandlerFunc) Handle(params PriceParams) PriceResponder {
	return fn(params)
}

// PriceHandler interface for that can handle valid price params
type PriceHandler interface {
	Handle(PriceParams) PriceResponder
}

// NewPrice creates a new http.Handler for the price operation
func NewPrice(ctx *middleware.Context, handler PriceHandler) *Price {
	return &Price{Context: ctx, Handler: handler}
}

/*Price swagger:route GET /price price

get Currency Price

*/
type Price struct {
	Context *middleware.Context
	Handler PriceHandler
}

func (o *Price) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPriceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PriceOKBody price o k body
//
// swagger:model PriceOKBody
type PriceOKBody struct {

	// result
	// Required: true
	Result interface{} `json:"result"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (o *PriceOKBody) UnmarshalJSON(data []byte) error {
	var props struct {

		// result
		// Required: true
		Result interface{} `json:"result"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	o.Result = props.Result
	return nil
}

// Validate validates this price o k body
func (o *PriceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PriceOKBody) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("priceOK"+"."+"result", "body", o.Result); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PriceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PriceOKBody) UnmarshalBinary(b []byte) error {
	var res PriceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}