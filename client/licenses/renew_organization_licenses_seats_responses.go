// Code generated by go-swagger; DO NOT EDIT.

package licenses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RenewOrganizationLicensesSeatsReader is a Reader for the RenewOrganizationLicensesSeats structure.
type RenewOrganizationLicensesSeatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenewOrganizationLicensesSeatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRenewOrganizationLicensesSeatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRenewOrganizationLicensesSeatsOK creates a RenewOrganizationLicensesSeatsOK with default headers values
func NewRenewOrganizationLicensesSeatsOK() *RenewOrganizationLicensesSeatsOK {
	return &RenewOrganizationLicensesSeatsOK{}
}

/*RenewOrganizationLicensesSeatsOK handles this case with default header values.

Successful operation
*/
type RenewOrganizationLicensesSeatsOK struct {
	Payload interface{}
}

func (o *RenewOrganizationLicensesSeatsOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/licenses/renewSeats][%d] renewOrganizationLicensesSeatsOK  %+v", 200, o.Payload)
}

func (o *RenewOrganizationLicensesSeatsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RenewOrganizationLicensesSeatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}