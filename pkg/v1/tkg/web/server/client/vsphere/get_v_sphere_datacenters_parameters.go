// Code generated by go-swagger; DO NOT EDIT.

package vsphere

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

// NewGetVSphereDatacentersParams creates a new GetVSphereDatacentersParams object
// with the default values initialized.
func NewGetVSphereDatacentersParams() *GetVSphereDatacentersParams {

	return &GetVSphereDatacentersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVSphereDatacentersParamsWithTimeout creates a new GetVSphereDatacentersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVSphereDatacentersParamsWithTimeout(timeout time.Duration) *GetVSphereDatacentersParams {

	return &GetVSphereDatacentersParams{

		timeout: timeout,
	}
}

// NewGetVSphereDatacentersParamsWithContext creates a new GetVSphereDatacentersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVSphereDatacentersParamsWithContext(ctx context.Context) *GetVSphereDatacentersParams {

	return &GetVSphereDatacentersParams{

		Context: ctx,
	}
}

// NewGetVSphereDatacentersParamsWithHTTPClient creates a new GetVSphereDatacentersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVSphereDatacentersParamsWithHTTPClient(client *http.Client) *GetVSphereDatacentersParams {

	return &GetVSphereDatacentersParams{
		HTTPClient: client,
	}
}

/*GetVSphereDatacentersParams contains all the parameters to send to the API endpoint
for the get v sphere datacenters operation typically these are written to a http.Request
*/
type GetVSphereDatacentersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v sphere datacenters params
func (o *GetVSphereDatacentersParams) WithTimeout(timeout time.Duration) *GetVSphereDatacentersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v sphere datacenters params
func (o *GetVSphereDatacentersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v sphere datacenters params
func (o *GetVSphereDatacentersParams) WithContext(ctx context.Context) *GetVSphereDatacentersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v sphere datacenters params
func (o *GetVSphereDatacentersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v sphere datacenters params
func (o *GetVSphereDatacentersParams) WithHTTPClient(client *http.Client) *GetVSphereDatacentersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v sphere datacenters params
func (o *GetVSphereDatacentersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetVSphereDatacentersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
