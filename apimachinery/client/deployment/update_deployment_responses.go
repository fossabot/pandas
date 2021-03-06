// Code generated by go-swagger; DO NOT EDIT.

package deployment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateDeploymentReader is a Reader for the UpdateDeployment structure.
type UpdateDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateDeploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateDeploymentOK creates a UpdateDeploymentOK with default headers values
func NewUpdateDeploymentOK() *UpdateDeploymentOK {
	return &UpdateDeploymentOK{}
}

/*UpdateDeploymentOK handles this case with default header values.

successfully operation
*/
type UpdateDeploymentOK struct {
}

func (o *UpdateDeploymentOK) Error() string {
	return fmt.Sprintf("[PUT /deployments/{deploymentId}][%d] updateDeploymentOK ", 200)
}

func (o *UpdateDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDeploymentNotFound creates a UpdateDeploymentNotFound with default headers values
func NewUpdateDeploymentNotFound() *UpdateDeploymentNotFound {
	return &UpdateDeploymentNotFound{}
}

/*UpdateDeploymentNotFound handles this case with default header values.

deployment not found
*/
type UpdateDeploymentNotFound struct {
}

func (o *UpdateDeploymentNotFound) Error() string {
	return fmt.Sprintf("[PUT /deployments/{deploymentId}][%d] updateDeploymentNotFound ", 404)
}

func (o *UpdateDeploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
