// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware-tanzu-private/core/pkg/v1/tkg/web/server/models"
)

// ApplyTKGConfigForVsphereReader is a Reader for the ApplyTKGConfigForVsphere structure.
type ApplyTKGConfigForVsphereReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplyTKGConfigForVsphereReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplyTKGConfigForVsphereOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewApplyTKGConfigForVsphereBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewApplyTKGConfigForVsphereUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewApplyTKGConfigForVsphereInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewApplyTKGConfigForVsphereOK creates a ApplyTKGConfigForVsphereOK with default headers values
func NewApplyTKGConfigForVsphereOK() *ApplyTKGConfigForVsphereOK {
	return &ApplyTKGConfigForVsphereOK{}
}

/*ApplyTKGConfigForVsphereOK handles this case with default header values.

apply changes to TKG configuration file successfully
*/
type ApplyTKGConfigForVsphereOK struct {
	Payload *models.ConfigFileInfo
}

func (o *ApplyTKGConfigForVsphereOK) Error() string {
	return fmt.Sprintf("[POST /api/providers/vsphere/tkgconfig][%d] applyTKGConfigForVsphereOK  %+v", 200, o.Payload)
}

func (o *ApplyTKGConfigForVsphereOK) GetPayload() *models.ConfigFileInfo {
	return o.Payload
}

func (o *ApplyTKGConfigForVsphereOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ConfigFileInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplyTKGConfigForVsphereBadRequest creates a ApplyTKGConfigForVsphereBadRequest with default headers values
func NewApplyTKGConfigForVsphereBadRequest() *ApplyTKGConfigForVsphereBadRequest {
	return &ApplyTKGConfigForVsphereBadRequest{}
}

/*ApplyTKGConfigForVsphereBadRequest handles this case with default header values.

Bad request
*/
type ApplyTKGConfigForVsphereBadRequest struct {
	Payload *models.Error
}

func (o *ApplyTKGConfigForVsphereBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/providers/vsphere/tkgconfig][%d] applyTKGConfigForVsphereBadRequest  %+v", 400, o.Payload)
}

func (o *ApplyTKGConfigForVsphereBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ApplyTKGConfigForVsphereBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplyTKGConfigForVsphereUnauthorized creates a ApplyTKGConfigForVsphereUnauthorized with default headers values
func NewApplyTKGConfigForVsphereUnauthorized() *ApplyTKGConfigForVsphereUnauthorized {
	return &ApplyTKGConfigForVsphereUnauthorized{}
}

/*ApplyTKGConfigForVsphereUnauthorized handles this case with default header values.

Incorrect credentials
*/
type ApplyTKGConfigForVsphereUnauthorized struct {
	Payload *models.Error
}

func (o *ApplyTKGConfigForVsphereUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/providers/vsphere/tkgconfig][%d] applyTKGConfigForVsphereUnauthorized  %+v", 401, o.Payload)
}

func (o *ApplyTKGConfigForVsphereUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ApplyTKGConfigForVsphereUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplyTKGConfigForVsphereInternalServerError creates a ApplyTKGConfigForVsphereInternalServerError with default headers values
func NewApplyTKGConfigForVsphereInternalServerError() *ApplyTKGConfigForVsphereInternalServerError {
	return &ApplyTKGConfigForVsphereInternalServerError{}
}

/*ApplyTKGConfigForVsphereInternalServerError handles this case with default header values.

Internal server error
*/
type ApplyTKGConfigForVsphereInternalServerError struct {
	Payload *models.Error
}

func (o *ApplyTKGConfigForVsphereInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/providers/vsphere/tkgconfig][%d] applyTKGConfigForVsphereInternalServerError  %+v", 500, o.Payload)
}

func (o *ApplyTKGConfigForVsphereInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ApplyTKGConfigForVsphereInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
