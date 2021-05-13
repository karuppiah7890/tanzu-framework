// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package tmc

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

// NewRetrieveTMCInstallYmlParams creates a new RetrieveTMCInstallYmlParams object
// with the default values initialized.
func NewRetrieveTMCInstallYmlParams() *RetrieveTMCInstallYmlParams {
	var ()
	return &RetrieveTMCInstallYmlParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveTMCInstallYmlParamsWithTimeout creates a new RetrieveTMCInstallYmlParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrieveTMCInstallYmlParamsWithTimeout(timeout time.Duration) *RetrieveTMCInstallYmlParams {
	var ()
	return &RetrieveTMCInstallYmlParams{

		timeout: timeout,
	}
}

// NewRetrieveTMCInstallYmlParamsWithContext creates a new RetrieveTMCInstallYmlParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetrieveTMCInstallYmlParamsWithContext(ctx context.Context) *RetrieveTMCInstallYmlParams {
	var ()
	return &RetrieveTMCInstallYmlParams{

		Context: ctx,
	}
}

// NewRetrieveTMCInstallYmlParamsWithHTTPClient creates a new RetrieveTMCInstallYmlParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrieveTMCInstallYmlParamsWithHTTPClient(client *http.Client) *RetrieveTMCInstallYmlParams {
	var ()
	return &RetrieveTMCInstallYmlParams{
		HTTPClient: client,
	}
}

/*RetrieveTMCInstallYmlParams contains all the parameters to send to the API endpoint
for the retrieve t m c install yml operation typically these are written to a http.Request
*/
type RetrieveTMCInstallYmlParams struct {

	/*URL
	  The URL to retrieve the TMC install yml from.

	*/
	URL string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retrieve t m c install yml params
func (o *RetrieveTMCInstallYmlParams) WithTimeout(timeout time.Duration) *RetrieveTMCInstallYmlParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve t m c install yml params
func (o *RetrieveTMCInstallYmlParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve t m c install yml params
func (o *RetrieveTMCInstallYmlParams) WithContext(ctx context.Context) *RetrieveTMCInstallYmlParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve t m c install yml params
func (o *RetrieveTMCInstallYmlParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve t m c install yml params
func (o *RetrieveTMCInstallYmlParams) WithHTTPClient(client *http.Client) *RetrieveTMCInstallYmlParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve t m c install yml params
func (o *RetrieveTMCInstallYmlParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithURL adds the url to the retrieve t m c install yml params
func (o *RetrieveTMCInstallYmlParams) WithURL(url string) *RetrieveTMCInstallYmlParams {
	o.SetURL(url)
	return o
}

// SetURL adds the url to the retrieve t m c install yml params
func (o *RetrieveTMCInstallYmlParams) SetURL(url string) {
	o.URL = url
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveTMCInstallYmlParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param url
	qrURL := o.URL
	qURL := qrURL
	if qURL != "" {
		if err := r.SetQueryParam("url", qURL); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
