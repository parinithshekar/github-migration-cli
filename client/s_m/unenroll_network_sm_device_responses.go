// Code generated by go-swagger; DO NOT EDIT.

package s_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UnenrollNetworkSmDeviceReader is a Reader for the UnenrollNetworkSmDevice structure.
type UnenrollNetworkSmDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnenrollNetworkSmDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnenrollNetworkSmDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUnenrollNetworkSmDeviceOK creates a UnenrollNetworkSmDeviceOK with default headers values
func NewUnenrollNetworkSmDeviceOK() *UnenrollNetworkSmDeviceOK {
	return &UnenrollNetworkSmDeviceOK{}
}

/*UnenrollNetworkSmDeviceOK handles this case with default header values.

Successful operation
*/
type UnenrollNetworkSmDeviceOK struct {
	Payload interface{}
}

func (o *UnenrollNetworkSmDeviceOK) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/sm/devices/{deviceId}/unenroll][%d] unenrollNetworkSmDeviceOK  %+v", 200, o.Payload)
}

func (o *UnenrollNetworkSmDeviceOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UnenrollNetworkSmDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}