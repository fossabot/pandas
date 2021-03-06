// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdateProjectDeviceOKCode is the HTTP code returned for type UpdateProjectDeviceOK
const UpdateProjectDeviceOKCode int = 200

/*UpdateProjectDeviceOK successful operation

swagger:response updateProjectDeviceOK
*/
type UpdateProjectDeviceOK struct {
}

// NewUpdateProjectDeviceOK creates UpdateProjectDeviceOK with default headers values
func NewUpdateProjectDeviceOK() *UpdateProjectDeviceOK {

	return &UpdateProjectDeviceOK{}
}

// WriteResponse to the client
func (o *UpdateProjectDeviceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdateProjectDeviceBadRequestCode is the HTTP code returned for type UpdateProjectDeviceBadRequest
const UpdateProjectDeviceBadRequestCode int = 400

/*UpdateProjectDeviceBadRequest Bad request

swagger:response updateProjectDeviceBadRequest
*/
type UpdateProjectDeviceBadRequest struct {
}

// NewUpdateProjectDeviceBadRequest creates UpdateProjectDeviceBadRequest with default headers values
func NewUpdateProjectDeviceBadRequest() *UpdateProjectDeviceBadRequest {

	return &UpdateProjectDeviceBadRequest{}
}

// WriteResponse to the client
func (o *UpdateProjectDeviceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// UpdateProjectDeviceInternalServerErrorCode is the HTTP code returned for type UpdateProjectDeviceInternalServerError
const UpdateProjectDeviceInternalServerErrorCode int = 500

/*UpdateProjectDeviceInternalServerError server internal error

swagger:response updateProjectDeviceInternalServerError
*/
type UpdateProjectDeviceInternalServerError struct {
}

// NewUpdateProjectDeviceInternalServerError creates UpdateProjectDeviceInternalServerError with default headers values
func NewUpdateProjectDeviceInternalServerError() *UpdateProjectDeviceInternalServerError {

	return &UpdateProjectDeviceInternalServerError{}
}

// WriteResponse to the client
func (o *UpdateProjectDeviceInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
