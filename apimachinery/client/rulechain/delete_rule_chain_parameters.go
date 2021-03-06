// Code generated by go-swagger; DO NOT EDIT.

package rulechain

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
)

// NewDeleteRuleChainParams creates a new DeleteRuleChainParams object
// with the default values initialized.
func NewDeleteRuleChainParams() *DeleteRuleChainParams {
	var ()
	return &DeleteRuleChainParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRuleChainParamsWithTimeout creates a new DeleteRuleChainParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRuleChainParamsWithTimeout(timeout time.Duration) *DeleteRuleChainParams {
	var ()
	return &DeleteRuleChainParams{

		timeout: timeout,
	}
}

// NewDeleteRuleChainParamsWithContext creates a new DeleteRuleChainParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRuleChainParamsWithContext(ctx context.Context) *DeleteRuleChainParams {
	var ()
	return &DeleteRuleChainParams{

		Context: ctx,
	}
}

// NewDeleteRuleChainParamsWithHTTPClient creates a new DeleteRuleChainParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRuleChainParamsWithHTTPClient(client *http.Client) *DeleteRuleChainParams {
	var ()
	return &DeleteRuleChainParams{
		HTTPClient: client,
	}
}

/*DeleteRuleChainParams contains all the parameters to send to the API endpoint
for the delete rule chain operation typically these are written to a http.Request
*/
type DeleteRuleChainParams struct {

	/*RuleChainID
	  rule chain id

	*/
	RuleChainID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete rule chain params
func (o *DeleteRuleChainParams) WithTimeout(timeout time.Duration) *DeleteRuleChainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete rule chain params
func (o *DeleteRuleChainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete rule chain params
func (o *DeleteRuleChainParams) WithContext(ctx context.Context) *DeleteRuleChainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete rule chain params
func (o *DeleteRuleChainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete rule chain params
func (o *DeleteRuleChainParams) WithHTTPClient(client *http.Client) *DeleteRuleChainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete rule chain params
func (o *DeleteRuleChainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRuleChainID adds the ruleChainID to the delete rule chain params
func (o *DeleteRuleChainParams) WithRuleChainID(ruleChainID string) *DeleteRuleChainParams {
	o.SetRuleChainID(ruleChainID)
	return o
}

// SetRuleChainID adds the ruleChainId to the delete rule chain params
func (o *DeleteRuleChainParams) SetRuleChainID(ruleChainID string) {
	o.RuleChainID = ruleChainID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRuleChainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ruleChainId
	if err := r.SetPathParam("ruleChainId", o.RuleChainID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
