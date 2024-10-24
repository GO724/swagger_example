// Code generated by go-swagger; DO NOT EDIT.

package admins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"swagger_example/models"
)

// GetCardsOKCode is the HTTP code returned for type GetCardsOK
const GetCardsOKCode int = 200

/*
GetCardsOK search results matching criteria

swagger:response getCardsOK
*/
type GetCardsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Card `json:"body,omitempty"`
}

// NewGetCardsOK creates GetCardsOK with default headers values
func NewGetCardsOK() *GetCardsOK {

	return &GetCardsOK{}
}

// WithPayload adds the payload to the get cards o k response
func (o *GetCardsOK) WithPayload(payload []*models.Card) *GetCardsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get cards o k response
func (o *GetCardsOK) SetPayload(payload []*models.Card) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCardsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Card, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetCardsBadRequestCode is the HTTP code returned for type GetCardsBadRequest
const GetCardsBadRequestCode int = 400

/*
GetCardsBadRequest bad input parameter

swagger:response getCardsBadRequest
*/
type GetCardsBadRequest struct {
}

// NewGetCardsBadRequest creates GetCardsBadRequest with default headers values
func NewGetCardsBadRequest() *GetCardsBadRequest {

	return &GetCardsBadRequest{}
}

// WriteResponse to the client
func (o *GetCardsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
