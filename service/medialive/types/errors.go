// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Placeholder documentation for BadGatewayException
type BadGatewayException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *BadGatewayException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BadGatewayException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BadGatewayException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "BadGatewayException"
	}
	return *e.ErrorCodeOverride
}
func (e *BadGatewayException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Placeholder documentation for BadRequestException
type BadRequestException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *BadRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BadRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BadRequestException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "BadRequestException"
	}
	return *e.ErrorCodeOverride
}
func (e *BadRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Placeholder documentation for ConflictException
type ConflictException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ConflictException"
	}
	return *e.ErrorCodeOverride
}
func (e *ConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Placeholder documentation for ForbiddenException
type ForbiddenException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ForbiddenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ForbiddenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ForbiddenException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ForbiddenException"
	}
	return *e.ErrorCodeOverride
}
func (e *ForbiddenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Placeholder documentation for GatewayTimeoutException
type GatewayTimeoutException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *GatewayTimeoutException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *GatewayTimeoutException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *GatewayTimeoutException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "GatewayTimeoutException"
	}
	return *e.ErrorCodeOverride
}
func (e *GatewayTimeoutException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Placeholder documentation for InternalServerErrorException
type InternalServerErrorException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalServerErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerErrorException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InternalServerErrorException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServerErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Placeholder documentation for NotFoundException
type NotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *NotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotFoundException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "NotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *NotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Placeholder documentation for TooManyRequestsException
type TooManyRequestsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *TooManyRequestsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyRequestsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyRequestsException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "TooManyRequestsException"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyRequestsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Placeholder documentation for UnprocessableEntityException
type UnprocessableEntityException struct {
	Message *string

	ErrorCodeOverride *string

	ValidationErrors []ValidationError

	noSmithyDocumentSerde
}

func (e *UnprocessableEntityException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnprocessableEntityException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnprocessableEntityException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "UnprocessableEntityException"
	}
	return *e.ErrorCodeOverride
}
func (e *UnprocessableEntityException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
