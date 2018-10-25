// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TransactionReceipt transaction receipt
// swagger:model TransactionReceipt
type TransactionReceipt struct {

	// data
	Data []io.ReadCloser `json:"data"`

	// events
	Events []*TransactionReceiptEventsItems0 `json:"events"`

	// state changes
	StateChanges []*TransactionReceiptStateChangesItems0 `json:"state_changes"`

	// transaction id
	TransactionID string `json:"transaction_id,omitempty"`
}

// Validate validates this transaction receipt
func (m *TransactionReceipt) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateChanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionReceipt) validateEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.Events) { // not required
		return nil
	}

	for i := 0; i < len(m.Events); i++ {
		if swag.IsZero(m.Events[i]) { // not required
			continue
		}

		if m.Events[i] != nil {
			if err := m.Events[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TransactionReceipt) validateStateChanges(formats strfmt.Registry) error {

	if swag.IsZero(m.StateChanges) { // not required
		return nil
	}

	for i := 0; i < len(m.StateChanges); i++ {
		if swag.IsZero(m.StateChanges[i]) { // not required
			continue
		}

		if m.StateChanges[i] != nil {
			if err := m.StateChanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("state_changes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionReceipt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionReceipt) UnmarshalBinary(b []byte) error {
	var res TransactionReceipt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TransactionReceiptEventsItems0 transaction receipt events items0
// swagger:model TransactionReceiptEventsItems0
type TransactionReceiptEventsItems0 struct {

	// attributes
	Attributes []*TransactionReceiptEventsItems0AttributesItems0 `json:"attributes"`

	// data
	// Format: binary
	Data io.ReadCloser `json:"data,omitempty"`

	// event type
	EventType string `json:"event_type,omitempty"`
}

// Validate validates this transaction receipt events items0
func (m *TransactionReceiptEventsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionReceiptEventsItems0) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	for i := 0; i < len(m.Attributes); i++ {
		if swag.IsZero(m.Attributes[i]) { // not required
			continue
		}

		if m.Attributes[i] != nil {
			if err := m.Attributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionReceiptEventsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionReceiptEventsItems0) UnmarshalBinary(b []byte) error {
	var res TransactionReceiptEventsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TransactionReceiptEventsItems0AttributesItems0 transaction receipt events items0 attributes items0
// swagger:model TransactionReceiptEventsItems0AttributesItems0
type TransactionReceiptEventsItems0AttributesItems0 struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this transaction receipt events items0 attributes items0
func (m *TransactionReceiptEventsItems0AttributesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TransactionReceiptEventsItems0AttributesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionReceiptEventsItems0AttributesItems0) UnmarshalBinary(b []byte) error {
	var res TransactionReceiptEventsItems0AttributesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TransactionReceiptStateChangesItems0 transaction receipt state changes items0
// swagger:model TransactionReceiptStateChangesItems0
type TransactionReceiptStateChangesItems0 struct {

	// address
	Address string `json:"address,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// value
	// Format: binary
	Value io.ReadCloser `json:"value,omitempty"`
}

// Validate validates this transaction receipt state changes items0
func (m *TransactionReceiptStateChangesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TransactionReceiptStateChangesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionReceiptStateChangesItems0) UnmarshalBinary(b []byte) error {
	var res TransactionReceiptStateChangesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
