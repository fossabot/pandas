// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cloustone/pandas/models"
)

// GetProjectsProjectIDOKCode is the HTTP code returned for type GetProjectsProjectIDOK
const GetProjectsProjectIDOKCode int = 200

/*GetProjectsProjectIDOK successful operation

swagger:response getProjectsProjectIdOK
*/
type GetProjectsProjectIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Project `json:"body,omitempty"`
}

// NewGetProjectsProjectIDOK creates GetProjectsProjectIDOK with default headers values
func NewGetProjectsProjectIDOK() *GetProjectsProjectIDOK {

	return &GetProjectsProjectIDOK{}
}

// WithPayload adds the payload to the get projects project Id o k response
func (o *GetProjectsProjectIDOK) WithPayload(payload *models.Project) *GetProjectsProjectIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get projects project Id o k response
func (o *GetProjectsProjectIDOK) SetPayload(payload *models.Project) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectsProjectIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProjectsProjectIDNotFoundCode is the HTTP code returned for type GetProjectsProjectIDNotFound
const GetProjectsProjectIDNotFoundCode int = 404

/*GetProjectsProjectIDNotFound Bad request

swagger:response getProjectsProjectIdNotFound
*/
type GetProjectsProjectIDNotFound struct {
}

// NewGetProjectsProjectIDNotFound creates GetProjectsProjectIDNotFound with default headers values
func NewGetProjectsProjectIDNotFound() *GetProjectsProjectIDNotFound {

	return &GetProjectsProjectIDNotFound{}
}

// WriteResponse to the client
func (o *GetProjectsProjectIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
