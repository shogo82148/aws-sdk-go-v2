// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Tests the filter pattern of a metric filter against a sample of log event
// messages. You can use this operation to validate the correctness of a metric
// filter pattern.
func (c *Client) TestMetricFilter(ctx context.Context, params *TestMetricFilterInput, optFns ...func(*Options)) (*TestMetricFilterOutput, error) {
	stack := middleware.NewStack("TestMetricFilter", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpTestMetricFilterMiddlewares(stack)
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
	addOpTestMetricFilterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTestMetricFilter(options.Region), middleware.Before)
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
			OperationName: "TestMetricFilter",
			Err:           err,
		}
	}
	out := result.(*TestMetricFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TestMetricFilterInput struct {

	// The log event messages to test.
	//
	// This member is required.
	LogEventMessages []*string

	// A symbolic description of how CloudWatch Logs should interpret the data in each
	// log event. For example, a log event may contain timestamps, IP addresses,
	// strings, and so on. You use the filter pattern to specify what to look for in
	// the log event message.
	//
	// This member is required.
	FilterPattern *string
}

type TestMetricFilterOutput struct {

	// The matched events.
	Matches []*types.MetricFilterMatchRecord

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpTestMetricFilterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpTestMetricFilter{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpTestMetricFilter{}, middleware.After)
}

func newServiceMetadataMiddleware_opTestMetricFilter(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "TestMetricFilter",
	}
}
