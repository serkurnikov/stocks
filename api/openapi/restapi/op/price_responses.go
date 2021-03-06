// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"stocks/api/openapi/model"
)

// PriceOKCode is the HTTP code returned for type PriceOK
const PriceOKCode int = 200

/*PriceOK response

swagger:response priceOK
*/
type PriceOK struct {

	/*
	  In: Body
	*/
	Payload *PriceOKBody `json:"body,omitempty"`
}

// NewPriceOK creates PriceOK with default headers values
func NewPriceOK() *PriceOK {

	return &PriceOK{}
}

// WithPayload adds the payload to the price o k response
func (o *PriceOK) WithPayload(payload *PriceOKBody) *PriceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the price o k response
func (o *PriceOK) SetPayload(payload *PriceOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PriceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *PriceOK) PriceResponder() {}

/*PriceDefault General errors using same model as used by go-swagger for validation errors.

swagger:response priceDefault
*/
type PriceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *model.Error `json:"body,omitempty"`
}

// NewPriceDefault creates PriceDefault with default headers values
func NewPriceDefault(code int) *PriceDefault {
	if code <= 0 {
		code = 500
	}

	return &PriceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the price default response
func (o *PriceDefault) WithStatusCode(code int) *PriceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the price default response
func (o *PriceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the price default response
func (o *PriceDefault) WithPayload(payload *model.Error) *PriceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the price default response
func (o *PriceDefault) SetPayload(payload *model.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PriceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *PriceDefault) PriceResponder() {}

type PriceNotImplementedResponder struct {
	middleware.Responder
}

func (*PriceNotImplementedResponder) PriceResponder() {}

func PriceNotImplemented() PriceResponder {
	return &PriceNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.Price has not yet been implemented",
		),
	}
}

type PriceResponder interface {
	middleware.Responder
	PriceResponder()
}
