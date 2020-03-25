// Code generated by go-swagger; DO NOT EDIT.

package p_i_i

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkPiiRequestReader is a Reader for the GetNetworkPiiRequest structure.
type GetNetworkPiiRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkPiiRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkPiiRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkPiiRequestOK creates a GetNetworkPiiRequestOK with default headers values
func NewGetNetworkPiiRequestOK() *GetNetworkPiiRequestOK {
	return &GetNetworkPiiRequestOK{}
}

/*GetNetworkPiiRequestOK handles this case with default header values.

Successful operation
*/
type GetNetworkPiiRequestOK struct {
	Payload interface{}
}

func (o *GetNetworkPiiRequestOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/pii/requests/{requestId}][%d] getNetworkPiiRequestOK  %+v", 200, o.Payload)
}

func (o *GetNetworkPiiRequestOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkPiiRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}