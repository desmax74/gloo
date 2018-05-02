// Code generated by go-swagger; DO NOT EDIT.

package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/solo-io/gloo/internal/function-discovery/updater/projectfn/imported/models"
)

// PostAppsAppRoutesReader is a Reader for the PostAppsAppRoutes structure.
type PostAppsAppRoutesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAppsAppRoutesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostAppsAppRoutesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostAppsAppRoutesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostAppsAppRoutesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostAppsAppRoutesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostAppsAppRoutesOK creates a PostAppsAppRoutesOK with default headers values
func NewPostAppsAppRoutesOK() *PostAppsAppRoutesOK {
	return &PostAppsAppRoutesOK{}
}

/*PostAppsAppRoutesOK handles this case with default header values.

Route created
*/
type PostAppsAppRoutesOK struct {
	Payload *models.RouteWrapper
}

func (o *PostAppsAppRoutesOK) Error() string {
	return fmt.Sprintf("[POST /apps/{app}/routes][%d] postAppsAppRoutesOK  %+v", 200, o.Payload)
}

func (o *PostAppsAppRoutesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RouteWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAppsAppRoutesBadRequest creates a PostAppsAppRoutesBadRequest with default headers values
func NewPostAppsAppRoutesBadRequest() *PostAppsAppRoutesBadRequest {
	return &PostAppsAppRoutesBadRequest{}
}

/*PostAppsAppRoutesBadRequest handles this case with default header values.

Invalid route due to parameters being missing or invalid.
*/
type PostAppsAppRoutesBadRequest struct {
	Payload *models.Error
}

func (o *PostAppsAppRoutesBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{app}/routes][%d] postAppsAppRoutesBadRequest  %+v", 400, o.Payload)
}

func (o *PostAppsAppRoutesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAppsAppRoutesConflict creates a PostAppsAppRoutesConflict with default headers values
func NewPostAppsAppRoutesConflict() *PostAppsAppRoutesConflict {
	return &PostAppsAppRoutesConflict{}
}

/*PostAppsAppRoutesConflict handles this case with default header values.

Route already exists.
*/
type PostAppsAppRoutesConflict struct {
	Payload *models.Error
}

func (o *PostAppsAppRoutesConflict) Error() string {
	return fmt.Sprintf("[POST /apps/{app}/routes][%d] postAppsAppRoutesConflict  %+v", 409, o.Payload)
}

func (o *PostAppsAppRoutesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAppsAppRoutesDefault creates a PostAppsAppRoutesDefault with default headers values
func NewPostAppsAppRoutesDefault(code int) *PostAppsAppRoutesDefault {
	return &PostAppsAppRoutesDefault{
		_statusCode: code,
	}
}

/*PostAppsAppRoutesDefault handles this case with default header values.

Unexpected error
*/
type PostAppsAppRoutesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post apps app routes default response
func (o *PostAppsAppRoutesDefault) Code() int {
	return o._statusCode
}

func (o *PostAppsAppRoutesDefault) Error() string {
	return fmt.Sprintf("[POST /apps/{app}/routes][%d] PostAppsAppRoutes default  %+v", o._statusCode, o.Payload)
}

func (o *PostAppsAppRoutesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}