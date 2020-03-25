// Code generated by go-swagger; DO NOT EDIT.

package switch_ports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cisco-sso/meraki-cli/models"
)

// NewUpdateDeviceSwitchPortParams creates a new UpdateDeviceSwitchPortParams object
// with the default values initialized.
func NewUpdateDeviceSwitchPortParams() *UpdateDeviceSwitchPortParams {
	var ()
	return &UpdateDeviceSwitchPortParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeviceSwitchPortParamsWithTimeout creates a new UpdateDeviceSwitchPortParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateDeviceSwitchPortParamsWithTimeout(timeout time.Duration) *UpdateDeviceSwitchPortParams {
	var ()
	return &UpdateDeviceSwitchPortParams{

		timeout: timeout,
	}
}

// NewUpdateDeviceSwitchPortParamsWithContext creates a new UpdateDeviceSwitchPortParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateDeviceSwitchPortParamsWithContext(ctx context.Context) *UpdateDeviceSwitchPortParams {
	var ()
	return &UpdateDeviceSwitchPortParams{

		Context: ctx,
	}
}

// NewUpdateDeviceSwitchPortParamsWithHTTPClient creates a new UpdateDeviceSwitchPortParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateDeviceSwitchPortParamsWithHTTPClient(client *http.Client) *UpdateDeviceSwitchPortParams {
	var ()
	return &UpdateDeviceSwitchPortParams{
		HTTPClient: client,
	}
}

/*UpdateDeviceSwitchPortParams contains all the parameters to send to the API endpoint
for the update device switch port operation typically these are written to a http.Request
*/
type UpdateDeviceSwitchPortParams struct {

	/*Number*/
	Number string
	/*Serial*/
	Serial string
	/*UpdateDeviceSwitchPort*/
	UpdateDeviceSwitchPort *models.UpdateDeviceSwitchPort

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update device switch port params
func (o *UpdateDeviceSwitchPortParams) WithTimeout(timeout time.Duration) *UpdateDeviceSwitchPortParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device switch port params
func (o *UpdateDeviceSwitchPortParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device switch port params
func (o *UpdateDeviceSwitchPortParams) WithContext(ctx context.Context) *UpdateDeviceSwitchPortParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device switch port params
func (o *UpdateDeviceSwitchPortParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device switch port params
func (o *UpdateDeviceSwitchPortParams) WithHTTPClient(client *http.Client) *UpdateDeviceSwitchPortParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device switch port params
func (o *UpdateDeviceSwitchPortParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNumber adds the number to the update device switch port params
func (o *UpdateDeviceSwitchPortParams) WithNumber(number string) *UpdateDeviceSwitchPortParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the update device switch port params
func (o *UpdateDeviceSwitchPortParams) SetNumber(number string) {
	o.Number = number
}

// WithSerial adds the serial to the update device switch port params
func (o *UpdateDeviceSwitchPortParams) WithSerial(serial string) *UpdateDeviceSwitchPortParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the update device switch port params
func (o *UpdateDeviceSwitchPortParams) SetSerial(serial string) {
	o.Serial = serial
}

// WithUpdateDeviceSwitchPort adds the updateDeviceSwitchPort to the update device switch port params
func (o *UpdateDeviceSwitchPortParams) WithUpdateDeviceSwitchPort(updateDeviceSwitchPort *models.UpdateDeviceSwitchPort) *UpdateDeviceSwitchPortParams {
	o.SetUpdateDeviceSwitchPort(updateDeviceSwitchPort)
	return o
}

// SetUpdateDeviceSwitchPort adds the updateDeviceSwitchPort to the update device switch port params
func (o *UpdateDeviceSwitchPortParams) SetUpdateDeviceSwitchPort(updateDeviceSwitchPort *models.UpdateDeviceSwitchPort) {
	o.UpdateDeviceSwitchPort = updateDeviceSwitchPort
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeviceSwitchPortParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param number
	if err := r.SetPathParam("number", o.Number); err != nil {
		return err
	}

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}

	if o.UpdateDeviceSwitchPort != nil {
		if err := r.SetBodyParam(o.UpdateDeviceSwitchPort); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}