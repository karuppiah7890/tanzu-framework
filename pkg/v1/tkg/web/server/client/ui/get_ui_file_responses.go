// Code generated by go-swagger; DO NOT EDIT.

package ui

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware-tanzu-private/core/pkg/v1/tkg/web/server/models"
)

// GetUIFileReader is a Reader for the GetUIFile structure.
type GetUIFileReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetUIFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUIFileOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUIFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUIFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUIFileOK creates a GetUIFileOK with default headers values
func NewGetUIFileOK(writer io.Writer) *GetUIFileOK {
	return &GetUIFileOK{
		Payload: writer,
	}
}

/*GetUIFileOK handles this case with default header values.

Successful operation
*/
type GetUIFileOK struct {
	Payload io.Writer
}

func (o *GetUIFileOK) Error() string {
	return fmt.Sprintf("[GET /{filename}][%d] getUiFileOK  %+v", 200, o.Payload)
}

func (o *GetUIFileOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetUIFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUIFileBadRequest creates a GetUIFileBadRequest with default headers values
func NewGetUIFileBadRequest() *GetUIFileBadRequest {
	return &GetUIFileBadRequest{}
}

/*GetUIFileBadRequest handles this case with default header values.

Bad request
*/
type GetUIFileBadRequest struct {
	Payload *models.Error
}

func (o *GetUIFileBadRequest) Error() string {
	return fmt.Sprintf("[GET /{filename}][%d] getUiFileBadRequest  %+v", 400, o.Payload)
}

func (o *GetUIFileBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUIFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUIFileInternalServerError creates a GetUIFileInternalServerError with default headers values
func NewGetUIFileInternalServerError() *GetUIFileInternalServerError {
	return &GetUIFileInternalServerError{}
}

/*GetUIFileInternalServerError handles this case with default header values.

Internal server error
*/
type GetUIFileInternalServerError struct {
	Payload *models.Error
}

func (o *GetUIFileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /{filename}][%d] getUiFileInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUIFileInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUIFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
