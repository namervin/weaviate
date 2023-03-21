//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// SchemaClassesIndexesGetReader is a Reader for the SchemaClassesIndexesGet structure.
type SchemaClassesIndexesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchemaClassesIndexesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSchemaClassesIndexesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSchemaClassesIndexesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSchemaClassesIndexesGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSchemaClassesIndexesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSchemaClassesIndexesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSchemaClassesIndexesGetOK creates a SchemaClassesIndexesGetOK with default headers values
func NewSchemaClassesIndexesGetOK() *SchemaClassesIndexesGetOK {
	return &SchemaClassesIndexesGetOK{}
}

/*
SchemaClassesIndexesGetOK describes a response with status code 200, with default header values.

OK. Indexes are listed in the response
*/
type SchemaClassesIndexesGetOK struct {
	Payload *models.IndexStatusList
}

// IsSuccess returns true when this schema classes indexes get o k response has a 2xx status code
func (o *SchemaClassesIndexesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this schema classes indexes get o k response has a 3xx status code
func (o *SchemaClassesIndexesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schema classes indexes get o k response has a 4xx status code
func (o *SchemaClassesIndexesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this schema classes indexes get o k response has a 5xx status code
func (o *SchemaClassesIndexesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this schema classes indexes get o k response a status code equal to that given
func (o *SchemaClassesIndexesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the schema classes indexes get o k response
func (o *SchemaClassesIndexesGetOK) Code() int {
	return 200
}

func (o *SchemaClassesIndexesGetOK) Error() string {
	return fmt.Sprintf("[GET /schema/{className}/indexes][%d] schemaClassesIndexesGetOK  %+v", 200, o.Payload)
}

func (o *SchemaClassesIndexesGetOK) String() string {
	return fmt.Sprintf("[GET /schema/{className}/indexes][%d] schemaClassesIndexesGetOK  %+v", 200, o.Payload)
}

func (o *SchemaClassesIndexesGetOK) GetPayload() *models.IndexStatusList {
	return o.Payload
}

func (o *SchemaClassesIndexesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IndexStatusList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaClassesIndexesGetUnauthorized creates a SchemaClassesIndexesGetUnauthorized with default headers values
func NewSchemaClassesIndexesGetUnauthorized() *SchemaClassesIndexesGetUnauthorized {
	return &SchemaClassesIndexesGetUnauthorized{}
}

/*
SchemaClassesIndexesGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type SchemaClassesIndexesGetUnauthorized struct {
}

// IsSuccess returns true when this schema classes indexes get unauthorized response has a 2xx status code
func (o *SchemaClassesIndexesGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schema classes indexes get unauthorized response has a 3xx status code
func (o *SchemaClassesIndexesGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schema classes indexes get unauthorized response has a 4xx status code
func (o *SchemaClassesIndexesGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this schema classes indexes get unauthorized response has a 5xx status code
func (o *SchemaClassesIndexesGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this schema classes indexes get unauthorized response a status code equal to that given
func (o *SchemaClassesIndexesGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the schema classes indexes get unauthorized response
func (o *SchemaClassesIndexesGetUnauthorized) Code() int {
	return 401
}

func (o *SchemaClassesIndexesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /schema/{className}/indexes][%d] schemaClassesIndexesGetUnauthorized ", 401)
}

func (o *SchemaClassesIndexesGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /schema/{className}/indexes][%d] schemaClassesIndexesGetUnauthorized ", 401)
}

func (o *SchemaClassesIndexesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSchemaClassesIndexesGetForbidden creates a SchemaClassesIndexesGetForbidden with default headers values
func NewSchemaClassesIndexesGetForbidden() *SchemaClassesIndexesGetForbidden {
	return &SchemaClassesIndexesGetForbidden{}
}

/*
SchemaClassesIndexesGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SchemaClassesIndexesGetForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this schema classes indexes get forbidden response has a 2xx status code
func (o *SchemaClassesIndexesGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schema classes indexes get forbidden response has a 3xx status code
func (o *SchemaClassesIndexesGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schema classes indexes get forbidden response has a 4xx status code
func (o *SchemaClassesIndexesGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this schema classes indexes get forbidden response has a 5xx status code
func (o *SchemaClassesIndexesGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this schema classes indexes get forbidden response a status code equal to that given
func (o *SchemaClassesIndexesGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the schema classes indexes get forbidden response
func (o *SchemaClassesIndexesGetForbidden) Code() int {
	return 403
}

func (o *SchemaClassesIndexesGetForbidden) Error() string {
	return fmt.Sprintf("[GET /schema/{className}/indexes][%d] schemaClassesIndexesGetForbidden  %+v", 403, o.Payload)
}

func (o *SchemaClassesIndexesGetForbidden) String() string {
	return fmt.Sprintf("[GET /schema/{className}/indexes][%d] schemaClassesIndexesGetForbidden  %+v", 403, o.Payload)
}

func (o *SchemaClassesIndexesGetForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaClassesIndexesGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaClassesIndexesGetNotFound creates a SchemaClassesIndexesGetNotFound with default headers values
func NewSchemaClassesIndexesGetNotFound() *SchemaClassesIndexesGetNotFound {
	return &SchemaClassesIndexesGetNotFound{}
}

/*
SchemaClassesIndexesGetNotFound describes a response with status code 404, with default header values.

This class does not exist
*/
type SchemaClassesIndexesGetNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this schema classes indexes get not found response has a 2xx status code
func (o *SchemaClassesIndexesGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schema classes indexes get not found response has a 3xx status code
func (o *SchemaClassesIndexesGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schema classes indexes get not found response has a 4xx status code
func (o *SchemaClassesIndexesGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this schema classes indexes get not found response has a 5xx status code
func (o *SchemaClassesIndexesGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this schema classes indexes get not found response a status code equal to that given
func (o *SchemaClassesIndexesGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the schema classes indexes get not found response
func (o *SchemaClassesIndexesGetNotFound) Code() int {
	return 404
}

func (o *SchemaClassesIndexesGetNotFound) Error() string {
	return fmt.Sprintf("[GET /schema/{className}/indexes][%d] schemaClassesIndexesGetNotFound  %+v", 404, o.Payload)
}

func (o *SchemaClassesIndexesGetNotFound) String() string {
	return fmt.Sprintf("[GET /schema/{className}/indexes][%d] schemaClassesIndexesGetNotFound  %+v", 404, o.Payload)
}

func (o *SchemaClassesIndexesGetNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaClassesIndexesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaClassesIndexesGetInternalServerError creates a SchemaClassesIndexesGetInternalServerError with default headers values
func NewSchemaClassesIndexesGetInternalServerError() *SchemaClassesIndexesGetInternalServerError {
	return &SchemaClassesIndexesGetInternalServerError{}
}

/*
SchemaClassesIndexesGetInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type SchemaClassesIndexesGetInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this schema classes indexes get internal server error response has a 2xx status code
func (o *SchemaClassesIndexesGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schema classes indexes get internal server error response has a 3xx status code
func (o *SchemaClassesIndexesGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schema classes indexes get internal server error response has a 4xx status code
func (o *SchemaClassesIndexesGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this schema classes indexes get internal server error response has a 5xx status code
func (o *SchemaClassesIndexesGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this schema classes indexes get internal server error response a status code equal to that given
func (o *SchemaClassesIndexesGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the schema classes indexes get internal server error response
func (o *SchemaClassesIndexesGetInternalServerError) Code() int {
	return 500
}

func (o *SchemaClassesIndexesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /schema/{className}/indexes][%d] schemaClassesIndexesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *SchemaClassesIndexesGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /schema/{className}/indexes][%d] schemaClassesIndexesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *SchemaClassesIndexesGetInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaClassesIndexesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}