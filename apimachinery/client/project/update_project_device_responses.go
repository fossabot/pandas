// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateProjectDeviceReader is a Reader for the UpdateProjectDevice structure.
type UpdateProjectDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateProjectDeviceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateProjectDeviceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateProjectDeviceOK creates a UpdateProjectDeviceOK with default headers values
func NewUpdateProjectDeviceOK() *UpdateProjectDeviceOK {
	return &UpdateProjectDeviceOK{}
}

/*UpdateProjectDeviceOK handles this case with default header values.

successful operation
*/
type UpdateProjectDeviceOK struct {
}

func (o *UpdateProjectDeviceOK) Error() string {
	return fmt.Sprintf("[PUT /project/{projectId}/devices/{deviceId}][%d] updateProjectDeviceOK ", 200)
}

func (o *UpdateProjectDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProjectDeviceBadRequest creates a UpdateProjectDeviceBadRequest with default headers values
func NewUpdateProjectDeviceBadRequest() *UpdateProjectDeviceBadRequest {
	return &UpdateProjectDeviceBadRequest{}
}

/*UpdateProjectDeviceBadRequest handles this case with default header values.

Bad request
*/
type UpdateProjectDeviceBadRequest struct {
}

func (o *UpdateProjectDeviceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /project/{projectId}/devices/{deviceId}][%d] updateProjectDeviceBadRequest ", 400)
}

func (o *UpdateProjectDeviceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProjectDeviceInternalServerError creates a UpdateProjectDeviceInternalServerError with default headers values
func NewUpdateProjectDeviceInternalServerError() *UpdateProjectDeviceInternalServerError {
	return &UpdateProjectDeviceInternalServerError{}
}

/*UpdateProjectDeviceInternalServerError handles this case with default header values.

server internal error
*/
type UpdateProjectDeviceInternalServerError struct {
}

func (o *UpdateProjectDeviceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /project/{projectId}/devices/{deviceId}][%d] updateProjectDeviceInternalServerError ", 500)
}

func (o *UpdateProjectDeviceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
