// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// User user
// swagger:model User
type User struct {
	Timedef

	// available balance of user
	Balance string `json:"balance,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// phone
	Phone string `json:"phone,omitempty"`

	// roles
	Roles []string `json:"roles"`

	// status
	// Enum: [OFFLINE ONLINE]
	Status string `json:"status,omitempty"`

	// total fund
	TotalFund string `json:"totalFund,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *User) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Timedef
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Timedef = aO0

	// AO1
	var dataAO1 struct {
		Balance string `json:"balance,omitempty"`

		Email string `json:"email,omitempty"`

		ID string `json:"id,omitempty"`

		Phone string `json:"phone,omitempty"`

		Roles []string `json:"roles,omitempty"`

		Status string `json:"status,omitempty"`

		TotalFund string `json:"totalFund,omitempty"`

		Username string `json:"username,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Balance = dataAO1.Balance

	m.Email = dataAO1.Email

	m.ID = dataAO1.ID

	m.Phone = dataAO1.Phone

	m.Roles = dataAO1.Roles

	m.Status = dataAO1.Status

	m.TotalFund = dataAO1.TotalFund

	m.Username = dataAO1.Username

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m User) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Timedef)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Balance string `json:"balance,omitempty"`

		Email string `json:"email,omitempty"`

		ID string `json:"id,omitempty"`

		Phone string `json:"phone,omitempty"`

		Roles []string `json:"roles,omitempty"`

		Status string `json:"status,omitempty"`

		TotalFund string `json:"totalFund,omitempty"`

		Username string `json:"username,omitempty"`
	}

	dataAO1.Balance = m.Balance

	dataAO1.Email = m.Email

	dataAO1.ID = m.ID

	dataAO1.Phone = m.Phone

	dataAO1.Roles = m.Roles

	dataAO1.Status = m.Status

	dataAO1.TotalFund = m.TotalFund

	dataAO1.Username = m.Username

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Timedef
	if err := m.Timedef.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var userRolesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["admin","user"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userRolesItemsEnum = append(userRolesItemsEnum, v)
	}
}

func (m *User) validateRolesItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, userRolesItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *User) validateRoles(formats strfmt.Registry) error {

	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	for i := 0; i < len(m.Roles); i++ {

		// value enum
		if err := m.validateRolesItemsEnum("roles"+"."+strconv.Itoa(i), "body", m.Roles[i]); err != nil {
			return err
		}

	}

	return nil
}

var userTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OFFLINE","ONLINE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userTypeStatusPropEnum = append(userTypeStatusPropEnum, v)
	}
}

// property enum
func (m *User) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, userTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *User) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}