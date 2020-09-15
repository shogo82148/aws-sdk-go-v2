// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// The service failed in an unexpected way.
type InternalServerException struct {
	Message *string

	Code *string
}

func (e *InternalServerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerException) ErrorCode() string             { return "InternalServerException" }
func (e *InternalServerException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *InternalServerException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *InternalServerException) HasCode() bool {
	return e.Code != nil
}
func (e *InternalServerException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InternalServerException) HasMessage() bool {
	return e.Message != nil
}

// Bad request. The request is missing required parameters or has invalid
// parameters.
type InvalidRequestException struct {
	Message *string

	MutuallyExclusiveParameters []*string
	Code                        *string
	RequiredParameters          []*string
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
func (e *InvalidRequestException) ErrorCode() string             { return "InvalidRequestException" }
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidRequestException) GetMutuallyExclusiveParameters() []*string {
	return e.MutuallyExclusiveParameters
}
func (e *InvalidRequestException) HasMutuallyExclusiveParameters() bool {
	return e.MutuallyExclusiveParameters != nil
}
func (e *InvalidRequestException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *InvalidRequestException) HasCode() bool {
	return e.Code != nil
}
func (e *InvalidRequestException) GetRequiredParameters() []*string {
	return e.RequiredParameters
}
func (e *InvalidRequestException) HasRequiredParameters() bool {
	return e.RequiredParameters != nil
}
func (e *InvalidRequestException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidRequestException) HasMessage() bool {
	return e.Message != nil
}

// The request failed because a limit was exceeded.
type LimitExceededException struct {
	Message *string

	Code         *string
	ResourceType *string
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
func (e *LimitExceededException) GetResourceType() string {
	return ptr.ToString(e.ResourceType)
}
func (e *LimitExceededException) HasResourceType() bool {
	return e.ResourceType != nil
}
func (e *LimitExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *LimitExceededException) HasMessage() bool {
	return e.Message != nil
}

// A requested resource was not found.
type ResourceNotFoundException struct {
	Message *string

	Code         *string
	ResourceIds  []*string
	ResourceType *string
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
func (e *ResourceNotFoundException) GetResourceIds() []*string {
	return e.ResourceIds
}
func (e *ResourceNotFoundException) HasResourceIds() bool {
	return e.ResourceIds != nil
}
func (e *ResourceNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceNotFoundException) HasMessage() bool {
	return e.Message != nil
}
func (e *ResourceNotFoundException) GetResourceType() string {
	return ptr.ToString(e.ResourceType)
}
func (e *ResourceNotFoundException) HasResourceType() bool {
	return e.ResourceType != nil
}