// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Sets the migration state of an application. For a given application identified
// by the value passed to ApplicationId, its status is set or updated by passing
// one of three values to Status: NOT_STARTED | IN_PROGRESS | COMPLETED.
func (c *Client) NotifyApplicationState(ctx context.Context, params *NotifyApplicationStateInput, optFns ...func(*Options)) (*NotifyApplicationStateOutput, error) {
	stack := middleware.NewStack("NotifyApplicationState", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpNotifyApplicationStateMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpNotifyApplicationStateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opNotifyApplicationState(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "NotifyApplicationState",
			Err:           err,
		}
	}
	out := result.(*NotifyApplicationStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NotifyApplicationStateInput struct {

	// Status of the application - Not Started, In-Progress, Complete.
	//
	// This member is required.
	Status types.ApplicationStatus

	// Optional boolean flag to indicate whether any effect should take place. Used to
	// test if the caller has permission to make the call.
	DryRun *bool

	// The configurationId in Application Discovery Service that uniquely identifies
	// the grouped application.
	//
	// This member is required.
	ApplicationId *string

	// The timestamp when the application state changed.
	UpdateDateTime *time.Time
}

type NotifyApplicationStateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpNotifyApplicationStateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpNotifyApplicationState{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpNotifyApplicationState{}, middleware.After)
}

func newServiceMetadataMiddleware_opNotifyApplicationState(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgh",
		OperationName: "NotifyApplicationState",
	}
}
