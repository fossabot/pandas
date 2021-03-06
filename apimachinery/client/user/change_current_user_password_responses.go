// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// ChangeCurrentUserPasswordReader is a Reader for the ChangeCurrentUserPassword structure.
type ChangeCurrentUserPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeCurrentUserPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeCurrentUserPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewChangeCurrentUserPasswordBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewChangeCurrentUserPasswordNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeCurrentUserPasswordOK creates a ChangeCurrentUserPasswordOK with default headers values
func NewChangeCurrentUserPasswordOK() *ChangeCurrentUserPasswordOK {
	return &ChangeCurrentUserPasswordOK{}
}

/*ChangeCurrentUserPasswordOK handles this case with default header values.

successful operation
*/
type ChangeCurrentUserPasswordOK struct {
}

func (o *ChangeCurrentUserPasswordOK) Error() string {
	return fmt.Sprintf("[PATCH /users/password][%d] changeCurrentUserPasswordOK ", 200)
}

func (o *ChangeCurrentUserPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeCurrentUserPasswordBadRequest creates a ChangeCurrentUserPasswordBadRequest with default headers values
func NewChangeCurrentUserPasswordBadRequest() *ChangeCurrentUserPasswordBadRequest {
	return &ChangeCurrentUserPasswordBadRequest{}
}

/*ChangeCurrentUserPasswordBadRequest handles this case with default header values.

Invalid username supplied
*/
type ChangeCurrentUserPasswordBadRequest struct {
}

func (o *ChangeCurrentUserPasswordBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /users/password][%d] changeCurrentUserPasswordBadRequest ", 400)
}

func (o *ChangeCurrentUserPasswordBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeCurrentUserPasswordNotFound creates a ChangeCurrentUserPasswordNotFound with default headers values
func NewChangeCurrentUserPasswordNotFound() *ChangeCurrentUserPasswordNotFound {
	return &ChangeCurrentUserPasswordNotFound{}
}

/*ChangeCurrentUserPasswordNotFound handles this case with default header values.

User not found
*/
type ChangeCurrentUserPasswordNotFound struct {
}

func (o *ChangeCurrentUserPasswordNotFound) Error() string {
	return fmt.Sprintf("[PATCH /users/password][%d] changeCurrentUserPasswordNotFound ", 404)
}

func (o *ChangeCurrentUserPasswordNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ChangeCurrentUserPasswordBody change current user password body
swagger:model ChangeCurrentUserPasswordBody
*/
type ChangeCurrentUserPasswordBody struct {

	// new passowrd
	NewPassowrd string `json:"new_passowrd,omitempty"`

	// old password
	OldPassword string `json:"old_password,omitempty"`
}

// Validate validates this change current user password body
func (o *ChangeCurrentUserPasswordBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeCurrentUserPasswordBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeCurrentUserPasswordBody) UnmarshalBinary(b []byte) error {
	var res ChangeCurrentUserPasswordBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
