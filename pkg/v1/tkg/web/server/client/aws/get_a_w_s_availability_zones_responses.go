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

// GetAWSAvailabilityZonesReader is a Reader for the GetAWSAvailabilityZones structure.
type GetAWSAvailabilityZonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAWSAvailabilityZonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAWSAvailabilityZonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAWSAvailabilityZonesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAWSAvailabilityZonesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAWSAvailabilityZonesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAWSAvailabilityZonesOK creates a GetAWSAvailabilityZonesOK with default headers values
func NewGetAWSAvailabilityZonesOK() *GetAWSAvailabilityZonesOK {
	return &GetAWSAvailabilityZonesOK{}
}

/*GetAWSAvailabilityZonesOK handles this case with default header values.

Successful retrieval of AWS availability zones
*/
type GetAWSAvailabilityZonesOK struct {
	Payload []*models.AWSAvailabilityZone
}

func (o *GetAWSAvailabilityZonesOK) Error() string {
	return fmt.Sprintf("[GET /api/providers/aws/AvailabilityZones][%d] getAWSAvailabilityZonesOK  %+v", 200, o.Payload)
}

func (o *GetAWSAvailabilityZonesOK) GetPayload() []*models.AWSAvailabilityZone {
	return o.Payload
}

func (o *GetAWSAvailabilityZonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAWSAvailabilityZonesBadRequest creates a GetAWSAvailabilityZonesBadRequest with default headers values
func NewGetAWSAvailabilityZonesBadRequest() *GetAWSAvailabilityZonesBadRequest {
	return &GetAWSAvailabilityZonesBadRequest{}
}

/*GetAWSAvailabilityZonesBadRequest handles this case with default header values.

Bad request
*/
type GetAWSAvailabilityZonesBadRequest struct {
	Payload *models.Error
}

func (o *GetAWSAvailabilityZonesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/providers/aws/AvailabilityZones][%d] getAWSAvailabilityZonesBadRequest  %+v", 400, o.Payload)
}

func (o *GetAWSAvailabilityZonesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAWSAvailabilityZonesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAWSAvailabilityZonesUnauthorized creates a GetAWSAvailabilityZonesUnauthorized with default headers values
func NewGetAWSAvailabilityZonesUnauthorized() *GetAWSAvailabilityZonesUnauthorized {
	return &GetAWSAvailabilityZonesUnauthorized{}
}

/*GetAWSAvailabilityZonesUnauthorized handles this case with default header values.

Incorrect credentials
*/
type GetAWSAvailabilityZonesUnauthorized struct {
	Payload *models.Error
}

func (o *GetAWSAvailabilityZonesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/providers/aws/AvailabilityZones][%d] getAWSAvailabilityZonesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAWSAvailabilityZonesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAWSAvailabilityZonesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAWSAvailabilityZonesInternalServerError creates a GetAWSAvailabilityZonesInternalServerError with default headers values
func NewGetAWSAvailabilityZonesInternalServerError() *GetAWSAvailabilityZonesInternalServerError {
	return &GetAWSAvailabilityZonesInternalServerError{}
}

/*GetAWSAvailabilityZonesInternalServerError handles this case with default header values.

Internal server error
*/
type GetAWSAvailabilityZonesInternalServerError struct {
	Payload *models.Error
}

func (o *GetAWSAvailabilityZonesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/providers/aws/AvailabilityZones][%d] getAWSAvailabilityZonesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAWSAvailabilityZonesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAWSAvailabilityZonesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
