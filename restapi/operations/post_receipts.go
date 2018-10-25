// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "github.com/faddat/sawmint/models"
)

// PostReceiptsHandlerFunc turns a function with the right signature into a post receipts handler
type PostReceiptsHandlerFunc func(PostReceiptsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostReceiptsHandlerFunc) Handle(params PostReceiptsParams) middleware.Responder {
	return fn(params)
}

// PostReceiptsHandler interface for that can handle valid post receipts params
type PostReceiptsHandler interface {
	Handle(PostReceiptsParams) middleware.Responder
}

// NewPostReceipts creates a new http.Handler for the post receipts operation
func NewPostReceipts(ctx *middleware.Context, handler PostReceiptsHandler) *PostReceipts {
	return &PostReceipts{Context: ctx, Handler: handler}
}

/*PostReceipts swagger:route POST /receipts postReceipts

Fetches the receipts for a set of transactions

Identical to `GET /receipts`, but takes ids of transactions as a JSON
formatted POST body rather than a query parameter. This allows for many
more receipts to be fetched and should be used with more than 15 ids.

Note that because query information is not encoded in the URL, no `link`
will be returned with this request.


*/
type PostReceipts struct {
	Context *middleware.Context
	Handler PostReceiptsHandler
}

func (o *PostReceipts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostReceiptsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostReceiptsOKBody post receipts o k body
// swagger:model PostReceiptsOKBody
type PostReceiptsOKBody struct {

	// data
	Data models.TransactionReceipts `json:"data,omitempty"`

	// link
	Link models.Link `json:"link,omitempty"`
}

// Validate validates this post receipts o k body
func (o *PostReceiptsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostReceiptsOKBody) validateLink(formats strfmt.Registry) error {

	if swag.IsZero(o.Link) { // not required
		return nil
	}

	if err := o.Link.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("postReceiptsOK" + "." + "link")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostReceiptsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostReceiptsOKBody) UnmarshalBinary(b []byte) error {
	var res PostReceiptsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
