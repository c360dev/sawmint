// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/c360dev/sawmint/models"
)

// GetPeersOKCode is the HTTP code returned for type GetPeersOK
const GetPeersOKCode int = 200

/*GetPeersOK Successfully retrieved peers

swagger:response getPeersOK
*/
type GetPeersOK struct {

	/*
	  In: Body
	*/
	Payload *GetPeersOKBody `json:"body,omitempty"`
}

// NewGetPeersOK creates GetPeersOK with default headers values
func NewGetPeersOK() *GetPeersOK {

	return &GetPeersOK{}
}

// WithPayload adds the payload to the get peers o k response
func (o *GetPeersOK) WithPayload(payload *GetPeersOKBody) *GetPeersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get peers o k response
func (o *GetPeersOK) SetPayload(payload *GetPeersOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPeersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPeersBadRequestCode is the HTTP code returned for type GetPeersBadRequest
const GetPeersBadRequestCode int = 400

/*GetPeersBadRequest Request was malformed

swagger:response getPeersBadRequest
*/
type GetPeersBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPeersBadRequest creates GetPeersBadRequest with default headers values
func NewGetPeersBadRequest() *GetPeersBadRequest {

	return &GetPeersBadRequest{}
}

// WithPayload adds the payload to the get peers bad request response
func (o *GetPeersBadRequest) WithPayload(payload *models.Error) *GetPeersBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get peers bad request response
func (o *GetPeersBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPeersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPeersInternalServerErrorCode is the HTTP code returned for type GetPeersInternalServerError
const GetPeersInternalServerErrorCode int = 500

/*GetPeersInternalServerError Something went wrong within the validator

swagger:response getPeersInternalServerError
*/
type GetPeersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPeersInternalServerError creates GetPeersInternalServerError with default headers values
func NewGetPeersInternalServerError() *GetPeersInternalServerError {

	return &GetPeersInternalServerError{}
}

// WithPayload adds the payload to the get peers internal server error response
func (o *GetPeersInternalServerError) WithPayload(payload *models.Error) *GetPeersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get peers internal server error response
func (o *GetPeersInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPeersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPeersServiceUnavailableCode is the HTTP code returned for type GetPeersServiceUnavailable
const GetPeersServiceUnavailableCode int = 503

/*GetPeersServiceUnavailable API is unable to reach the validator

swagger:response getPeersServiceUnavailable
*/
type GetPeersServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPeersServiceUnavailable creates GetPeersServiceUnavailable with default headers values
func NewGetPeersServiceUnavailable() *GetPeersServiceUnavailable {

	return &GetPeersServiceUnavailable{}
}

// WithPayload adds the payload to the get peers service unavailable response
func (o *GetPeersServiceUnavailable) WithPayload(payload *models.Error) *GetPeersServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get peers service unavailable response
func (o *GetPeersServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPeersServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
