// Code generated by go-swagger; DO NOT EDIT.

package admins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"swagger_example/models"
)

// GetPersonsOKCode is the HTTP code returned for type GetPersonsOK
const GetPersonsOKCode int = 200

/*
GetPersonsOK search results matching criteria

swagger:response getPersonsOK
*/
type GetPersonsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Person `json:"body,omitempty"`
}

// NewGetPersonsOK creates GetPersonsOK with default headers values
func NewGetPersonsOK() *GetPersonsOK {

	return &GetPersonsOK{}
}

// WithPayload adds the payload to the get persons o k response
func (o *GetPersonsOK) WithPayload(payload []*models.Person) *GetPersonsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get persons o k response
func (o *GetPersonsOK) SetPayload(payload []*models.Person) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPersonsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Person, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetPersonsBadRequestCode is the HTTP code returned for type GetPersonsBadRequest
const GetPersonsBadRequestCode int = 400

/*
GetPersonsBadRequest bad input parameter

swagger:response getPersonsBadRequest
*/
type GetPersonsBadRequest struct {
}

// NewGetPersonsBadRequest creates GetPersonsBadRequest with default headers values
func NewGetPersonsBadRequest() *GetPersonsBadRequest {

	return &GetPersonsBadRequest{}
}

// WriteResponse to the client
func (o *GetPersonsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
