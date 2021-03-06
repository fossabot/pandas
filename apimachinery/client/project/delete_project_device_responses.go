// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteProjectDeviceReader is a Reader for the DeleteProjectDevice structure.
type DeleteProjectDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProjectDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteProjectDeviceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteProjectDeviceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteProjectDeviceOK creates a DeleteProjectDeviceOK with default headers values
func NewDeleteProjectDeviceOK() *DeleteProjectDeviceOK {
	return &DeleteProjectDeviceOK{}
}

/*DeleteProjectDeviceOK handles this case with default header values.

successful operation
*/
type DeleteProjectDeviceOK struct {
}

func (o *DeleteProjectDeviceOK) Error() string {
	return fmt.Sprintf("[DELETE /project/{projectId}/devices/{deviceId}][%d] deleteProjectDeviceOK ", 200)
}

func (o *DeleteProjectDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectDeviceBadRequest creates a DeleteProjectDeviceBadRequest with default headers values
func NewDeleteProjectDeviceBadRequest() *DeleteProjectDeviceBadRequest {
	return &DeleteProjectDeviceBadRequest{}
}

/*DeleteProjectDeviceBadRequest handles this case with default header values.

Bad request
*/
type DeleteProjectDeviceBadRequest struct {
}

func (o *DeleteProjectDeviceBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /project/{projectId}/devices/{deviceId}][%d] deleteProjectDeviceBadRequest ", 400)
}

func (o *DeleteProjectDeviceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectDeviceInternalServerError creates a DeleteProjectDeviceInternalServerError with default headers values
func NewDeleteProjectDeviceInternalServerError() *DeleteProjectDeviceInternalServerError {
	return &DeleteProjectDeviceInternalServerError{}
}

/*DeleteProjectDeviceInternalServerError handles this case with default header values.

server internal error
*/
type DeleteProjectDeviceInternalServerError struct {
}

func (o *DeleteProjectDeviceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /project/{projectId}/devices/{deviceId}][%d] deleteProjectDeviceInternalServerError ", 500)
}

func (o *DeleteProjectDeviceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
