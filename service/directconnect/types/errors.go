// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// One or more parameters are not valid.
type DirectConnectClientException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DirectConnectClientException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DirectConnectClientException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DirectConnectClientException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "DirectConnectClientException"
	}
	return *e.ErrorCodeOverride
}
func (e *DirectConnectClientException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A server-side error occurred.
type DirectConnectServerException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DirectConnectServerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DirectConnectServerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DirectConnectServerException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "DirectConnectServerException"
	}
	return *e.ErrorCodeOverride
}
func (e *DirectConnectServerException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// A tag key was specified more than once.
type DuplicateTagKeysException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DuplicateTagKeysException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicateTagKeysException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicateTagKeysException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "DuplicateTagKeysException"
	}
	return *e.ErrorCodeOverride
}
func (e *DuplicateTagKeysException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have reached the limit on the number of tags that can be assigned.
type TooManyTagsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *TooManyTagsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTagsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTagsException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "TooManyTagsException"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyTagsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
