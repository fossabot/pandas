// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cloustone/pandas/models"
)

// GetDeviceLogOKCode is the HTTP code returned for type GetDeviceLogOK
const GetDeviceLogOKCode int = 200

/*GetDeviceLogOK successful operation

swagger:response getDeviceLogOK
*/
type GetDeviceLogOK struct {

	/*
	  In: Body
	*/
	Payload []*models.DeviceLog `json:"body,omitempty"`
}

// NewGetDeviceLogOK creates GetDeviceLogOK with default headers values
func NewGetDeviceLogOK() *GetDeviceLogOK {

	return &GetDeviceLogOK{}
}

// WithPayload adds the payload to the get device log o k response
func (o *GetDeviceLogOK) WithPayload(payload []*models.DeviceLog) *GetDeviceLogOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get device log o k response
func (o *GetDeviceLogOK) SetPayload(payload []*models.DeviceLog) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDeviceLogOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.DeviceLog, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetDeviceLogNotFoundCode is the HTTP code returned for type GetDeviceLogNotFound
const GetDeviceLogNotFoundCode int = 404

/*GetDeviceLogNotFound device not found.

swagger:response getDeviceLogNotFound
*/
type GetDeviceLogNotFound struct {
}

// NewGetDeviceLogNotFound creates GetDeviceLogNotFound with default headers values
func NewGetDeviceLogNotFound() *GetDeviceLogNotFound {

	return &GetDeviceLogNotFound{}
}

// WriteResponse to the client
func (o *GetDeviceLogNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
