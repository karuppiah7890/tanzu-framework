// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware-tanzu-private/core/pkg/v1/tkg/web/server/models"
)

// SetAWSEndpointReader is a Reader for the SetAWSEndpoint structure.
type SetAWSEndpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetAWSEndpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSetAWSEndpointCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetAWSEndpointBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSetAWSEndpointUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetAWSEndpointInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetAWSEndpointCreated creates a SetAWSEndpointCreated with default headers values
func NewSetAWSEndpointCreated() *SetAWSEndpointCreated {
	return &SetAWSEndpointCreated{}
}

/*SetAWSEndpointCreated handles this case with default header values.

Connection successful
*/
type SetAWSEndpointCreated struct{}

func (o *SetAWSEndpointCreated) Error() string {
	return fmt.Sprintf("[POST /api/providers/aws][%d] setAWSEndpointCreated ", 201)
}

func (o *SetAWSEndpointCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	return nil
}

// NewSetAWSEndpointBadRequest creates a SetAWSEndpointBadRequest with default headers values
func NewSetAWSEndpointBadRequest() *SetAWSEndpointBadRequest {
	return &SetAWSEndpointBadRequest{}
}

/*SetAWSEndpointBadRequest handles this case with default header values.

Bad request
*/
type SetAWSEndpointBadRequest struct {
	Payload *models.Error
}

func (o *SetAWSEndpointBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/providers/aws][%d] setAWSEndpointBadRequest  %+v", 400, o.Payload)
}

func (o *SetAWSEndpointBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetAWSEndpointBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAWSEndpointUnauthorized creates a SetAWSEndpointUnauthorized with default headers values
func NewSetAWSEndpointUnauthorized() *SetAWSEndpointUnauthorized {
	return &SetAWSEndpointUnauthorized{}
}

/*SetAWSEndpointUnauthorized handles this case with default header values.

Incorrect credentials
*/
type SetAWSEndpointUnauthorized struct {
	Payload *models.Error
}

func (o *SetAWSEndpointUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/providers/aws][%d] setAWSEndpointUnauthorized  %+v", 401, o.Payload)
}

func (o *SetAWSEndpointUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetAWSEndpointUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAWSEndpointInternalServerError creates a SetAWSEndpointInternalServerError with default headers values
func NewSetAWSEndpointInternalServerError() *SetAWSEndpointInternalServerError {
	return &SetAWSEndpointInternalServerError{}
}

/*SetAWSEndpointInternalServerError handles this case with default header values.

Internal server error
*/
type SetAWSEndpointInternalServerError struct {
	Payload *models.Error
}

func (o *SetAWSEndpointInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/providers/aws][%d] setAWSEndpointInternalServerError  %+v", 500, o.Payload)
}

func (o *SetAWSEndpointInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetAWSEndpointInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
