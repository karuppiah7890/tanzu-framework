// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAzureVnetsParams creates a new GetAzureVnetsParams object
// no default values defined in spec.
func NewGetAzureVnetsParams() GetAzureVnetsParams {

	return GetAzureVnetsParams{}
}

// GetAzureVnetsParams contains all the bound params for the get azure vnets operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAzureVnets
type GetAzureVnetsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Azure region
	  Required: true
	  In: query
	*/
	Location string
	/*Name of the Azure resource group
	  Required: true
	  In: path
	*/
	ResourceGroupName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAzureVnetsParams() beforehand.
func (o *GetAzureVnetsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qLocation, qhkLocation, _ := qs.GetOK("location")
	if err := o.bindLocation(qLocation, qhkLocation, route.Formats); err != nil {
		res = append(res, err)
	}

	rResourceGroupName, rhkResourceGroupName, _ := route.Params.GetOK("resourceGroupName")
	if err := o.bindResourceGroupName(rResourceGroupName, rhkResourceGroupName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindLocation binds and validates parameter Location from query.
func (o *GetAzureVnetsParams) bindLocation(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("location", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("location", "query", raw); err != nil {
		return err
	}

	o.Location = raw

	return nil
}

// bindResourceGroupName binds and validates parameter ResourceGroupName from path.
func (o *GetAzureVnetsParams) bindResourceGroupName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ResourceGroupName = raw

	return nil
}
