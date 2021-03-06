// Code generated by go-swagger; DO NOT EDIT.

package rulechain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUploadRuleChainParams creates a new UploadRuleChainParams object
// no default values defined in spec.
func NewUploadRuleChainParams() UploadRuleChainParams {

	return UploadRuleChainParams{}
}

// UploadRuleChainParams contains all the bound params for the upload rule chain operation
// typically these are obtained from a http.Request
//
// swagger:parameters uploadRuleChain
type UploadRuleChainParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*upload address
	  Required: true
	  In: formData
	*/
	Path string
	/*rule chain identifier
	  Required: true
	  In: path
	*/
	RuleChainID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUploadRuleChainParams() beforehand.
func (o *UploadRuleChainParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}
	fds := runtime.Values(r.Form)

	fdPath, fdhkPath, _ := fds.GetOK("path")
	if err := o.bindPath(fdPath, fdhkPath, route.Formats); err != nil {
		res = append(res, err)
	}

	rRuleChainID, rhkRuleChainID, _ := route.Params.GetOK("ruleChainId")
	if err := o.bindRuleChainID(rRuleChainID, rhkRuleChainID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindPath binds and validates parameter Path from formData.
func (o *UploadRuleChainParams) bindPath(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("path", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("path", "formData", raw); err != nil {
		return err
	}

	o.Path = raw

	return nil
}

// bindRuleChainID binds and validates parameter RuleChainID from path.
func (o *UploadRuleChainParams) bindRuleChainID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.RuleChainID = raw

	return nil
}
