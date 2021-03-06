// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cloustone/pandas/models"
)

// NewUpdateProjectDeviceParams creates a new UpdateProjectDeviceParams object
// with the default values initialized.
func NewUpdateProjectDeviceParams() *UpdateProjectDeviceParams {
	var ()
	return &UpdateProjectDeviceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProjectDeviceParamsWithTimeout creates a new UpdateProjectDeviceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateProjectDeviceParamsWithTimeout(timeout time.Duration) *UpdateProjectDeviceParams {
	var ()
	return &UpdateProjectDeviceParams{

		timeout: timeout,
	}
}

// NewUpdateProjectDeviceParamsWithContext creates a new UpdateProjectDeviceParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateProjectDeviceParamsWithContext(ctx context.Context) *UpdateProjectDeviceParams {
	var ()
	return &UpdateProjectDeviceParams{

		Context: ctx,
	}
}

// NewUpdateProjectDeviceParamsWithHTTPClient creates a new UpdateProjectDeviceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateProjectDeviceParamsWithHTTPClient(client *http.Client) *UpdateProjectDeviceParams {
	var ()
	return &UpdateProjectDeviceParams{
		HTTPClient: client,
	}
}

/*UpdateProjectDeviceParams contains all the parameters to send to the API endpoint
for the update project device operation typically these are written to a http.Request
*/
type UpdateProjectDeviceParams struct {

	/*Device*/
	Device *models.Device
	/*DeviceID
	  device id

	*/
	DeviceID string
	/*ProjectID
	  project id

	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update project device params
func (o *UpdateProjectDeviceParams) WithTimeout(timeout time.Duration) *UpdateProjectDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update project device params
func (o *UpdateProjectDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update project device params
func (o *UpdateProjectDeviceParams) WithContext(ctx context.Context) *UpdateProjectDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update project device params
func (o *UpdateProjectDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update project device params
func (o *UpdateProjectDeviceParams) WithHTTPClient(client *http.Client) *UpdateProjectDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update project device params
func (o *UpdateProjectDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevice adds the device to the update project device params
func (o *UpdateProjectDeviceParams) WithDevice(device *models.Device) *UpdateProjectDeviceParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the update project device params
func (o *UpdateProjectDeviceParams) SetDevice(device *models.Device) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the update project device params
func (o *UpdateProjectDeviceParams) WithDeviceID(deviceID string) *UpdateProjectDeviceParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the update project device params
func (o *UpdateProjectDeviceParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithProjectID adds the projectID to the update project device params
func (o *UpdateProjectDeviceParams) WithProjectID(projectID string) *UpdateProjectDeviceParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the update project device params
func (o *UpdateProjectDeviceParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProjectDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Device != nil {
		if err := r.SetBodyParam(o.Device); err != nil {
			return err
		}
	}

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	// path param projectId
	if err := r.SetPathParam("projectId", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
