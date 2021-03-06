// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetModelsParams creates a new GetModelsParams object
// with the default values initialized.
func NewGetModelsParams() *GetModelsParams {
	var ()
	return &GetModelsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetModelsParamsWithTimeout creates a new GetModelsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetModelsParamsWithTimeout(timeout time.Duration) *GetModelsParams {
	var ()
	return &GetModelsParams{

		timeout: timeout,
	}
}

// NewGetModelsParamsWithContext creates a new GetModelsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetModelsParamsWithContext(ctx context.Context) *GetModelsParams {
	var ()
	return &GetModelsParams{

		Context: ctx,
	}
}

// NewGetModelsParamsWithHTTPClient creates a new GetModelsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetModelsParamsWithHTTPClient(client *http.Client) *GetModelsParams {
	var ()
	return &GetModelsParams{
		HTTPClient: client,
	}
}

/*GetModelsParams contains all the parameters to send to the API endpoint
for the get models operation typically these are written to a http.Request
*/
type GetModelsParams struct {

	/*PageNumber
	  Page number

	*/
	PageNumber *int64
	/*PageSize
	  Number of persons returned

	*/
	PageSize *int64
	/*Q
	  query object for.You can get query key from rulechain object. This is a json string. For example:
	  * 模糊检索name,description,category
	  {"name": "product"}
	  {"description": "abcd"}
	  {"category": "abcd"}
	  * 多条件模糊检索(and)
	  {"name": "product", "description": "abcd"}
	  * (deprecated) 模糊检索created_at,updated_at
	  {"created_at": "2018-10-11T09:13:26Z"}
	  {"updated_at": "2018-10-11T09:13:26Z"}
	  * 精确检索user_id,id,template_id,key,secret,status,data_format
	  {"user_id": "bevh8dkvr53g2n6u9l70"}
	  {"id": "bevh8dkvr53g2n6u9l70"}
	  {"template_id": "bevh8dkvr53g2n6u9l70"}
	  {"key": "bevh8dkvr53g2n6u9l70"}
	  {"secret": "bevh8dkvr53g2n6u9l70"}
	  {"data_format": "JSON"}
	  {"data_format": "XML"}


	*/
	Q *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get models params
func (o *GetModelsParams) WithTimeout(timeout time.Duration) *GetModelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get models params
func (o *GetModelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get models params
func (o *GetModelsParams) WithContext(ctx context.Context) *GetModelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get models params
func (o *GetModelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get models params
func (o *GetModelsParams) WithHTTPClient(client *http.Client) *GetModelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get models params
func (o *GetModelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get models params
func (o *GetModelsParams) WithPageNumber(pageNumber *int64) *GetModelsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get models params
func (o *GetModelsParams) SetPageNumber(pageNumber *int64) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get models params
func (o *GetModelsParams) WithPageSize(pageSize *int64) *GetModelsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get models params
func (o *GetModelsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithQ adds the q to the get models params
func (o *GetModelsParams) WithQ(q *string) *GetModelsParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the get models params
func (o *GetModelsParams) SetQ(q *string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *GetModelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber int64
		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt64(qrPageNumber)
		if qPageNumber != "" {
			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
