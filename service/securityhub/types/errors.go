// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// You don't have permission to perform the action specified in the request.
type AccessDeniedException struct {
	Message *string

	Code *string
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string             { return "AccessDeniedException" }
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *AccessDeniedException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *AccessDeniedException) HasCode() bool {
	return e.Code != nil
}
func (e *AccessDeniedException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *AccessDeniedException) HasMessage() bool {
	return e.Message != nil
}

// Internal server error.
type InternalException struct {
	Message *string

	Code *string
}

func (e *InternalException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalException) ErrorCode() string             { return "InternalException" }
func (e *InternalException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *InternalException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InternalException) HasMessage() bool {
	return e.Message != nil
}
func (e *InternalException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *InternalException) HasCode() bool {
	return e.Code != nil
}

// AWS Security Hub isn't enabled for the account used to make this request.
type InvalidAccessException struct {
	Message *string

	Code *string
}

func (e *InvalidAccessException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidAccessException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidAccessException) ErrorCode() string             { return "InvalidAccessException" }
func (e *InvalidAccessException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidAccessException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *InvalidAccessException) HasCode() bool {
	return e.Code != nil
}
func (e *InvalidAccessException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidAccessException) HasMessage() bool {
	return e.Message != nil
}

// The request was rejected because you supplied an invalid or out-of-range value
// for an input parameter.
type InvalidInputException struct {
	Message *string

	Code *string
}

func (e *InvalidInputException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidInputException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidInputException) ErrorCode() string             { return "InvalidInputException" }
func (e *InvalidInputException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidInputException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidInputException) HasMessage() bool {
	return e.Message != nil
}
func (e *InvalidInputException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *InvalidInputException) HasCode() bool {
	return e.Code != nil
}

// The request was rejected because it attempted to create resources beyond the
// current AWS account limits. The error code describes the limit exceeded.
type LimitExceededException struct {
	Message *string

	Code *string
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
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *LimitExceededException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *LimitExceededException) HasCode() bool {
	return e.Code != nil
}
func (e *LimitExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *LimitExceededException) HasMessage() bool {
	return e.Message != nil
}

// The resource specified in the request conflicts with an existing resource.
type ResourceConflictException struct {
	Message *string

	Code *string
}

func (e *ResourceConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceConflictException) ErrorCode() string             { return "ResourceConflictException" }
func (e *ResourceConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceConflictException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *ResourceConflictException) HasCode() bool {
	return e.Code != nil
}
func (e *ResourceConflictException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceConflictException) HasMessage() bool {
	return e.Message != nil
}

// The request was rejected because we can't find the specified resource.
type ResourceNotFoundException struct {
	Message *string

	Code *string
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
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceNotFoundException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *ResourceNotFoundException) HasCode() bool {
	return e.Code != nil
}
func (e *ResourceNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceNotFoundException) HasMessage() bool {
	return e.Message != nil
}