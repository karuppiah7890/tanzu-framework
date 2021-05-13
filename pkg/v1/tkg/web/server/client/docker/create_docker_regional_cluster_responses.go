// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package docker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware-tanzu-private/core/pkg/v1/tkg/web/server/models"
)

// CreateDockerRegionalClusterReader is a Reader for the CreateDockerRegionalCluster structure.
type CreateDockerRegionalClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDockerRegionalClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDockerRegionalClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDockerRegionalClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateDockerRegionalClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDockerRegionalClusterOK creates a CreateDockerRegionalClusterOK with default headers values
func NewCreateDockerRegionalClusterOK() *CreateDockerRegionalClusterOK {
	return &CreateDockerRegionalClusterOK{}
}

/*CreateDockerRegionalClusterOK handles this case with default header values.

Creating regional cluster started successfully
*/
type CreateDockerRegionalClusterOK struct {
	Payload string
}

func (o *CreateDockerRegionalClusterOK) Error() string {
	return fmt.Sprintf("[POST /api/providers/docker/create][%d] createDockerRegionalClusterOK  %+v", 200, o.Payload)
}

func (o *CreateDockerRegionalClusterOK) GetPayload() string {
	return o.Payload
}

func (o *CreateDockerRegionalClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDockerRegionalClusterBadRequest creates a CreateDockerRegionalClusterBadRequest with default headers values
func NewCreateDockerRegionalClusterBadRequest() *CreateDockerRegionalClusterBadRequest {
	return &CreateDockerRegionalClusterBadRequest{}
}

/*CreateDockerRegionalClusterBadRequest handles this case with default header values.

Bad request
*/
type CreateDockerRegionalClusterBadRequest struct {
	Payload *models.Error
}

func (o *CreateDockerRegionalClusterBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/providers/docker/create][%d] createDockerRegionalClusterBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDockerRegionalClusterBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDockerRegionalClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDockerRegionalClusterInternalServerError creates a CreateDockerRegionalClusterInternalServerError with default headers values
func NewCreateDockerRegionalClusterInternalServerError() *CreateDockerRegionalClusterInternalServerError {
	return &CreateDockerRegionalClusterInternalServerError{}
}

/*CreateDockerRegionalClusterInternalServerError handles this case with default header values.

Internal server error
*/
type CreateDockerRegionalClusterInternalServerError struct {
	Payload *models.Error
}

func (o *CreateDockerRegionalClusterInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/providers/docker/create][%d] createDockerRegionalClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateDockerRegionalClusterInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDockerRegionalClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
