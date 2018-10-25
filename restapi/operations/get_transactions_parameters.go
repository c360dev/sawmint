// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTransactionsParams creates a new GetTransactionsParams object
// with the default values initialized.
func NewGetTransactionsParams() GetTransactionsParams {

	var (
		// initialize parameters with default values

		headDefault  = string("latest")
		limitDefault = int64(1000)
	)

	return GetTransactionsParams{
		Head: &headDefault,

		Limit: &limitDefault,
	}
}

// GetTransactionsParams contains all the bound params for the get transactions operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetTransactions
type GetTransactionsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Index or id of head block
	  In: query
	  Default: "latest"
	*/
	Head *string
	/*Number of items to return
	  In: query
	  Default: 1000
	*/
	Limit *int64
	/*If the list should be reversed
	  In: query
	*/
	Reverse *string
	/*Id to start paging (inclusive)
	  In: query
	*/
	Start *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetTransactionsParams() beforehand.
func (o *GetTransactionsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qHead, qhkHead, _ := qs.GetOK("head")
	if err := o.bindHead(qHead, qhkHead, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qReverse, qhkReverse, _ := qs.GetOK("reverse")
	if err := o.bindReverse(qReverse, qhkReverse, route.Formats); err != nil {
		res = append(res, err)
	}

	qStart, qhkStart, _ := qs.GetOK("start")
	if err := o.bindStart(qStart, qhkStart, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindHead binds and validates parameter Head from query.
func (o *GetTransactionsParams) bindHead(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetTransactionsParams()
		return nil
	}

	o.Head = &raw

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetTransactionsParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetTransactionsParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	return nil
}

// bindReverse binds and validates parameter Reverse from query.
func (o *GetTransactionsParams) bindReverse(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Reverse = &raw

	return nil
}

// bindStart binds and validates parameter Start from query.
func (o *GetTransactionsParams) bindStart(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Start = &raw

	return nil
}
