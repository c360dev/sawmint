// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetReceiptsParams creates a new GetReceiptsParams object
// no default values defined in spec.
func NewGetReceiptsParams() GetReceiptsParams {

	return GetReceiptsParams{}
}

// GetReceiptsParams contains all the bound params for the get receipts operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetReceipts
type GetReceiptsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A comma-separated list of transaction ids
	  Required: true
	  In: query
	*/
	ID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetReceiptsParams() beforehand.
func (o *GetReceiptsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qID, qhkID, _ := qs.GetOK("id")
	if err := o.bindID(qID, qhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from query.
func (o *GetReceiptsParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("id", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("id", "query", raw); err != nil {
		return err
	}

	o.ID = raw

	return nil
}
