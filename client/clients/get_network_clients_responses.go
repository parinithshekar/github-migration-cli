// Code generated by go-swagger; DO NOT EDIT.

package clients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkClientsReader is a Reader for the GetNetworkClients structure.
type GetNetworkClientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkClientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkClientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkClientsOK creates a GetNetworkClientsOK with default headers values
func NewGetNetworkClientsOK() *GetNetworkClientsOK {
	return &GetNetworkClientsOK{}
}

/*GetNetworkClientsOK handles this case with default header values.

Successful operation
*/
type GetNetworkClientsOK struct {
	Payload interface{}
}

func (o *GetNetworkClientsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/clients][%d] getNetworkClientsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkClientsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkClientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}