// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// CreateModelOKCode is the HTTP code returned for type CreateModelOK
const CreateModelOKCode int = 200

/*CreateModelOK successful operation

swagger:response createModelOK
*/
type CreateModelOK struct {
}

// NewCreateModelOK creates CreateModelOK with default headers values
func NewCreateModelOK() *CreateModelOK {

	return &CreateModelOK{}
}

// WriteResponse to the client
func (o *CreateModelOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// CreateModelBadRequestCode is the HTTP code returned for type CreateModelBadRequest
const CreateModelBadRequestCode int = 400

/*CreateModelBadRequest bad request

swagger:response createModelBadRequest
*/
type CreateModelBadRequest struct {
}

// NewCreateModelBadRequest creates CreateModelBadRequest with default headers values
func NewCreateModelBadRequest() *CreateModelBadRequest {

	return &CreateModelBadRequest{}
}

// WriteResponse to the client
func (o *CreateModelBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// CreateModelNotFoundCode is the HTTP code returned for type CreateModelNotFound
const CreateModelNotFoundCode int = 404

/*CreateModelNotFound device not found

swagger:response createModelNotFound
*/
type CreateModelNotFound struct {
}

// NewCreateModelNotFound creates CreateModelNotFound with default headers values
func NewCreateModelNotFound() *CreateModelNotFound {

	return &CreateModelNotFound{}
}

// WriteResponse to the client
func (o *CreateModelNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// CreateModelInternalServerErrorCode is the HTTP code returned for type CreateModelInternalServerError
const CreateModelInternalServerErrorCode int = 500

/*CreateModelInternalServerError server internal error

swagger:response createModelInternalServerError
*/
type CreateModelInternalServerError struct {
}

// NewCreateModelInternalServerError creates CreateModelInternalServerError with default headers values
func NewCreateModelInternalServerError() *CreateModelInternalServerError {

	return &CreateModelInternalServerError{}
}

// WriteResponse to the client
func (o *CreateModelInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
