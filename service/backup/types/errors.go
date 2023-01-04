// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The required resource already exists.
type AlreadyExistsException struct {
	Message *string

	ErrorCodeOverride *string

	Code             *string
	CreatorRequestId *string
	Arn              *string
	Type             *string
	Context          *string

	noSmithyDocumentSerde
}

func (e *AlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AlreadyExistsException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "AlreadyExistsException"
	}
	return *e.ErrorCodeOverride
}
func (e *AlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Backup can't perform the action that you requested until it finishes performing
// a previous action. Try again later.
type ConflictException struct {
	Message *string

	ErrorCodeOverride *string

	Code    *string
	Type    *string
	Context *string

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

// A dependent Amazon Web Services service or resource returned an error to the
// Backup service, and the action cannot be completed.
type DependencyFailureException struct {
	Message *string

	ErrorCodeOverride *string

	Code    *string
	Type    *string
	Context *string

	noSmithyDocumentSerde
}

func (e *DependencyFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DependencyFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DependencyFailureException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "DependencyFailureException"
	}
	return *e.ErrorCodeOverride
}
func (e *DependencyFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Indicates that something is wrong with a parameter's value. For example, the
// value is out of range.
type InvalidParameterValueException struct {
	Message *string

	ErrorCodeOverride *string

	Code    *string
	Type    *string
	Context *string

	noSmithyDocumentSerde
}

func (e *InvalidParameterValueException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterValueException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterValueException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidParameterValueException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidParameterValueException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that something is wrong with the input to the request. For example, a
// parameter is of the wrong type.
type InvalidRequestException struct {
	Message *string

	ErrorCodeOverride *string

	Code    *string
	Type    *string
	Context *string

	noSmithyDocumentSerde
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidRequestException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Backup is already performing an action on this recovery point. It can't perform
// the action you requested until the first action finishes. Try again later.
type InvalidResourceStateException struct {
	Message *string

	ErrorCodeOverride *string

	Code    *string
	Type    *string
	Context *string

	noSmithyDocumentSerde
}

func (e *InvalidResourceStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidResourceStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidResourceStateException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidResourceStateException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidResourceStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A limit in the request has been exceeded; for example, a maximum number of items
// allowed in a request.
type LimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	Code    *string
	Type    *string
	Context *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "LimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that a required parameter is missing.
type MissingParameterValueException struct {
	Message *string

	ErrorCodeOverride *string

	Code    *string
	Type    *string
	Context *string

	noSmithyDocumentSerde
}

func (e *MissingParameterValueException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MissingParameterValueException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MissingParameterValueException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "MissingParameterValueException"
	}
	return *e.ErrorCodeOverride
}
func (e *MissingParameterValueException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A resource that is required for the action doesn't exist.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	Code    *string
	Type    *string
	Context *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ResourceNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request failed due to a temporary failure of the server.
type ServiceUnavailableException struct {
	Message *string

	ErrorCodeOverride *string

	Code    *string
	Type    *string
	Context *string

	noSmithyDocumentSerde
}

func (e *ServiceUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUnavailableException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ServiceUnavailableException"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
