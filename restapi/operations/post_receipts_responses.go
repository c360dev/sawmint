// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/c360dev/sawmint/models"
)

// PostReceiptsOKCode is the HTTP code returned for type PostReceiptsOK
const PostReceiptsOKCode int = 200

/*PostReceiptsOK Successfully retrieved transaction receipts

swagger:response postReceiptsOK
*/
type PostReceiptsOK struct {

	/*
	  In: Body
	*/
	Payload *PostReceiptsOKBody `json:"body,omitempty"`
}

// NewPostReceiptsOK creates PostReceiptsOK with default headers values
func NewPostReceiptsOK() *PostReceiptsOK {

	return &PostReceiptsOK{}
}

// WithPayload adds the payload to the post receipts o k response
func (o *PostReceiptsOK) WithPayload(payload *PostReceiptsOKBody) *PostReceiptsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post receipts o k response
func (o *PostReceiptsOK) SetPayload(payload *PostReceiptsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostReceiptsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostReceiptsBadRequestCode is the HTTP code returned for type PostReceiptsBadRequest
const PostReceiptsBadRequestCode int = 400

/*PostReceiptsBadRequest Request was malformed

swagger:response postReceiptsBadRequest
*/
type PostReceiptsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostReceiptsBadRequest creates PostReceiptsBadRequest with default headers values
func NewPostReceiptsBadRequest() *PostReceiptsBadRequest {

	return &PostReceiptsBadRequest{}
}

// WithPayload adds the payload to the post receipts bad request response
func (o *PostReceiptsBadRequest) WithPayload(payload *models.Error) *PostReceiptsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post receipts bad request response
func (o *PostReceiptsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostReceiptsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostReceiptsInternalServerErrorCode is the HTTP code returned for type PostReceiptsInternalServerError
const PostReceiptsInternalServerErrorCode int = 500

/*PostReceiptsInternalServerError Something went wrong within the validator

swagger:response postReceiptsInternalServerError
*/
type PostReceiptsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostReceiptsInternalServerError creates PostReceiptsInternalServerError with default headers values
func NewPostReceiptsInternalServerError() *PostReceiptsInternalServerError {

	return &PostReceiptsInternalServerError{}
}

// WithPayload adds the payload to the post receipts internal server error response
func (o *PostReceiptsInternalServerError) WithPayload(payload *models.Error) *PostReceiptsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post receipts internal server error response
func (o *PostReceiptsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostReceiptsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostReceiptsServiceUnavailableCode is the HTTP code returned for type PostReceiptsServiceUnavailable
const PostReceiptsServiceUnavailableCode int = 503

/*PostReceiptsServiceUnavailable API is unable to reach the validator

swagger:response postReceiptsServiceUnavailable
*/
type PostReceiptsServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostReceiptsServiceUnavailable creates PostReceiptsServiceUnavailable with default headers values
func NewPostReceiptsServiceUnavailable() *PostReceiptsServiceUnavailable {

	return &PostReceiptsServiceUnavailable{}
}

// WithPayload adds the payload to the post receipts service unavailable response
func (o *PostReceiptsServiceUnavailable) WithPayload(payload *models.Error) *PostReceiptsServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post receipts service unavailable response
func (o *PostReceiptsServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostReceiptsServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
