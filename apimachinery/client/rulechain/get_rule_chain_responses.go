// Code generated by go-swagger; DO NOT EDIT.

package rulechain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cloustone/pandas/models"
)

// GetRuleChainReader is a Reader for the GetRuleChain structure.
type GetRuleChainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuleChainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRuleChainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRuleChainBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRuleChainNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRuleChainOK creates a GetRuleChainOK with default headers values
func NewGetRuleChainOK() *GetRuleChainOK {
	return &GetRuleChainOK{}
}

/*GetRuleChainOK handles this case with default header values.

successfully operation
*/
type GetRuleChainOK struct {
	Payload *models.RuleChain
}

func (o *GetRuleChainOK) Error() string {
	return fmt.Sprintf("[GET /rulechains/{ruleChainId}][%d] getRuleChainOK  %+v", 200, o.Payload)
}

func (o *GetRuleChainOK) GetPayload() *models.RuleChain {
	return o.Payload
}

func (o *GetRuleChainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuleChain)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuleChainBadRequest creates a GetRuleChainBadRequest with default headers values
func NewGetRuleChainBadRequest() *GetRuleChainBadRequest {
	return &GetRuleChainBadRequest{}
}

/*GetRuleChainBadRequest handles this case with default header values.

Bad request
*/
type GetRuleChainBadRequest struct {
}

func (o *GetRuleChainBadRequest) Error() string {
	return fmt.Sprintf("[GET /rulechains/{ruleChainId}][%d] getRuleChainBadRequest ", 400)
}

func (o *GetRuleChainBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRuleChainNotFound creates a GetRuleChainNotFound with default headers values
func NewGetRuleChainNotFound() *GetRuleChainNotFound {
	return &GetRuleChainNotFound{}
}

/*GetRuleChainNotFound handles this case with default header values.

rule chain not found
*/
type GetRuleChainNotFound struct {
}

func (o *GetRuleChainNotFound) Error() string {
	return fmt.Sprintf("[GET /rulechains/{ruleChainId}][%d] getRuleChainNotFound ", 404)
}

func (o *GetRuleChainNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
