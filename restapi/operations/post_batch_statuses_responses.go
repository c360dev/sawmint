// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/faddat/sawmint/models"
)

// PostBatchStatusesOKCode is the HTTP code returned for type PostBatchStatusesOK
const PostBatchStatusesOKCode int = 200

/*PostBatchStatusesOK Successfully retrieved statuses

swagger:response postBatchStatusesOK
*/
type PostBatchStatusesOK struct {

	/*
	  In: Body
	*/
	Payload *PostBatchStatusesOKBody `json:"body,omitempty"`
}

// NewPostBatchStatusesOK creates PostBatchStatusesOK with default headers values
func NewPostBatchStatusesOK() *PostBatchStatusesOK {

	return &PostBatchStatusesOK{}
}

// WithPayload adds the payload to the post batch statuses o k response
func (o *PostBatchStatusesOK) WithPayload(payload *PostBatchStatusesOKBody) *PostBatchStatusesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post batch statuses o k response
func (o *PostBatchStatusesOK) SetPayload(payload *PostBatchStatusesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostBatchStatusesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostBatchStatusesBadRequestCode is the HTTP code returned for type PostBatchStatusesBadRequest
const PostBatchStatusesBadRequestCode int = 400

/*PostBatchStatusesBadRequest Request was malformed

swagger:response postBatchStatusesBadRequest
*/
type PostBatchStatusesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostBatchStatusesBadRequest creates PostBatchStatusesBadRequest with default headers values
func NewPostBatchStatusesBadRequest() *PostBatchStatusesBadRequest {

	return &PostBatchStatusesBadRequest{}
}

// WithPayload adds the payload to the post batch statuses bad request response
func (o *PostBatchStatusesBadRequest) WithPayload(payload *models.Error) *PostBatchStatusesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post batch statuses bad request response
func (o *PostBatchStatusesBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostBatchStatusesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostBatchStatusesInternalServerErrorCode is the HTTP code returned for type PostBatchStatusesInternalServerError
const PostBatchStatusesInternalServerErrorCode int = 500

/*PostBatchStatusesInternalServerError Something went wrong within the validator

swagger:response postBatchStatusesInternalServerError
*/
type PostBatchStatusesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostBatchStatusesInternalServerError creates PostBatchStatusesInternalServerError with default headers values
func NewPostBatchStatusesInternalServerError() *PostBatchStatusesInternalServerError {

	return &PostBatchStatusesInternalServerError{}
}

// WithPayload adds the payload to the post batch statuses internal server error response
func (o *PostBatchStatusesInternalServerError) WithPayload(payload *models.Error) *PostBatchStatusesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post batch statuses internal server error response
func (o *PostBatchStatusesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostBatchStatusesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostBatchStatusesServiceUnavailableCode is the HTTP code returned for type PostBatchStatusesServiceUnavailable
const PostBatchStatusesServiceUnavailableCode int = 503

/*PostBatchStatusesServiceUnavailable API is unable to reach the validator

swagger:response postBatchStatusesServiceUnavailable
*/
type PostBatchStatusesServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostBatchStatusesServiceUnavailable creates PostBatchStatusesServiceUnavailable with default headers values
func NewPostBatchStatusesServiceUnavailable() *PostBatchStatusesServiceUnavailable {

	return &PostBatchStatusesServiceUnavailable{}
}

// WithPayload adds the payload to the post batch statuses service unavailable response
func (o *PostBatchStatusesServiceUnavailable) WithPayload(payload *models.Error) *PostBatchStatusesServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post batch statuses service unavailable response
func (o *PostBatchStatusesServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostBatchStatusesServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}