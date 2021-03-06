// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/c360dev/sawmint/models"
)

// GetStateOKCode is the HTTP code returned for type GetStateOK
const GetStateOKCode int = 200

/*GetStateOK Successfully retrieved state data

swagger:response getStateOK
*/
type GetStateOK struct {

	/*
	  In: Body
	*/
	Payload *GetStateOKBody `json:"body,omitempty"`
}

// NewGetStateOK creates GetStateOK with default headers values
func NewGetStateOK() *GetStateOK {

	return &GetStateOK{}
}

// WithPayload adds the payload to the get state o k response
func (o *GetStateOK) WithPayload(payload *GetStateOKBody) *GetStateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get state o k response
func (o *GetStateOK) SetPayload(payload *GetStateOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetStateBadRequestCode is the HTTP code returned for type GetStateBadRequest
const GetStateBadRequestCode int = 400

/*GetStateBadRequest Request was malformed

swagger:response getStateBadRequest
*/
type GetStateBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetStateBadRequest creates GetStateBadRequest with default headers values
func NewGetStateBadRequest() *GetStateBadRequest {

	return &GetStateBadRequest{}
}

// WithPayload adds the payload to the get state bad request response
func (o *GetStateBadRequest) WithPayload(payload *models.Error) *GetStateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get state bad request response
func (o *GetStateBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetStateInternalServerErrorCode is the HTTP code returned for type GetStateInternalServerError
const GetStateInternalServerErrorCode int = 500

/*GetStateInternalServerError Something went wrong within the validator

swagger:response getStateInternalServerError
*/
type GetStateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetStateInternalServerError creates GetStateInternalServerError with default headers values
func NewGetStateInternalServerError() *GetStateInternalServerError {

	return &GetStateInternalServerError{}
}

// WithPayload adds the payload to the get state internal server error response
func (o *GetStateInternalServerError) WithPayload(payload *models.Error) *GetStateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get state internal server error response
func (o *GetStateInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetStateServiceUnavailableCode is the HTTP code returned for type GetStateServiceUnavailable
const GetStateServiceUnavailableCode int = 503

/*GetStateServiceUnavailable API is unable to reach the validator

swagger:response getStateServiceUnavailable
*/
type GetStateServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetStateServiceUnavailable creates GetStateServiceUnavailable with default headers values
func NewGetStateServiceUnavailable() *GetStateServiceUnavailable {

	return &GetStateServiceUnavailable{}
}

// WithPayload adds the payload to the get state service unavailable response
func (o *GetStateServiceUnavailable) WithPayload(payload *models.Error) *GetStateServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get state service unavailable response
func (o *GetStateServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStateServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
