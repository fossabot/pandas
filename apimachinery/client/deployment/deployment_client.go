// Code generated by go-swagger; DO NOT EDIT.

package deployment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new deployment API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for deployment API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateDeployment creates a new deployment

create a new deployment
*/
func (a *Client) CreateDeployment(params *CreateDeploymentParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDeploymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDeploymentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createDeployment",
		Method:             "POST",
		PathPattern:        "/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateDeploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDeploymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createDeployment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteDeployment deletes specified deployment
*/
func (a *Client) DeleteDeployment(params *DeleteDeploymentParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteDeploymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDeploymentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDeployment",
		Method:             "DELETE",
		PathPattern:        "/deployments/{deploymentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteDeploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteDeploymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteDeployment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDeployment gets specified deployment
*/
func (a *Client) GetDeployment(params *GetDeploymentParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeployment",
		Method:             "GET",
		PathPattern:        "/deployments/{deploymentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDeploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeploymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeployment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDeployments gets all deployments
*/
func (a *Client) GetDeployments(params *GetDeploymentsParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeployments",
		Method:             "GET",
		PathPattern:        "/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDeploymentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeploymentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDeploymentsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SetDeploymentStatus updates specified deployment
*/
func (a *Client) SetDeploymentStatus(params *SetDeploymentStatusParams, authInfo runtime.ClientAuthInfoWriter) (*SetDeploymentStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetDeploymentStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setDeploymentStatus",
		Method:             "PATCH",
		PathPattern:        "/deployments/{deploymentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetDeploymentStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetDeploymentStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setDeploymentStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateDeployment updates specified deployment
*/
func (a *Client) UpdateDeployment(params *UpdateDeploymentParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateDeploymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDeploymentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateDeployment",
		Method:             "PUT",
		PathPattern:        "/deployments/{deploymentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateDeploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateDeploymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateDeployment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
