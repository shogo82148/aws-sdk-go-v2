// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops a previously started monitoring schedule.
func (c *Client) StopMonitoringSchedule(ctx context.Context, params *StopMonitoringScheduleInput, optFns ...func(*Options)) (*StopMonitoringScheduleOutput, error) {
	stack := middleware.NewStack("StopMonitoringSchedule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStopMonitoringScheduleMiddlewares(stack)
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
	addOpStopMonitoringScheduleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopMonitoringSchedule(options.Region), middleware.Before)
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
			OperationName: "StopMonitoringSchedule",
			Err:           err,
		}
	}
	out := result.(*StopMonitoringScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopMonitoringScheduleInput struct {

	// The name of the schedule to stop.
	//
	// This member is required.
	MonitoringScheduleName *string
}

type StopMonitoringScheduleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStopMonitoringScheduleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStopMonitoringSchedule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopMonitoringSchedule{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopMonitoringSchedule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "StopMonitoringSchedule",
	}
}
