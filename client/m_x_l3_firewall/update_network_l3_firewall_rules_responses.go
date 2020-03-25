// Code generated by go-swagger; DO NOT EDIT.

package m_x_l3_firewall

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateNetworkL3FirewallRulesReader is a Reader for the UpdateNetworkL3FirewallRules structure.
type UpdateNetworkL3FirewallRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkL3FirewallRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkL3FirewallRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNetworkL3FirewallRulesOK creates a UpdateNetworkL3FirewallRulesOK with default headers values
func NewUpdateNetworkL3FirewallRulesOK() *UpdateNetworkL3FirewallRulesOK {
	return &UpdateNetworkL3FirewallRulesOK{}
}

/*UpdateNetworkL3FirewallRulesOK handles this case with default header values.

Successful operation
*/
type UpdateNetworkL3FirewallRulesOK struct {
	Payload interface{}
}

func (o *UpdateNetworkL3FirewallRulesOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/l3FirewallRules][%d] updateNetworkL3FirewallRulesOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkL3FirewallRulesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkL3FirewallRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}