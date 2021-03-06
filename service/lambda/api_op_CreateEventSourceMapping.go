// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a mapping between an event source and an AWS Lambda function. Lambda
// reads items from the event source and triggers the function. For details about
// each event source type, see the following topics.
//
//     * Using AWS Lambda with
// Amazon DynamoDB (https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html)
//
//
// * Using AWS Lambda with Amazon Kinesis
// (https://docs.aws.amazon.com/lambda/latest/dg/with-kinesis.html)
//
//     * Using
// AWS Lambda with Amazon SQS
// (https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html)
//
// The following
// error handling options are only available for stream sources (DynamoDB and
// Kinesis):
//
//     * BisectBatchOnFunctionError - If the function returns an error,
// split the batch in two and retry.
//
//     * DestinationConfig - Send discarded
// records to an Amazon SQS queue or Amazon SNS topic.
//
//     *
// MaximumRecordAgeInSeconds - Discard records older than the specified age.
//
//     *
// MaximumRetryAttempts - Discard records after the specified number of retries.
//
//
// * ParallelizationFactor - Process multiple batches from each shard concurrently.
func (c *Client) CreateEventSourceMapping(ctx context.Context, params *CreateEventSourceMappingInput, optFns ...func(*Options)) (*CreateEventSourceMappingOutput, error) {
	stack := middleware.NewStack("CreateEventSourceMapping", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateEventSourceMappingMiddlewares(stack)
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
	addOpCreateEventSourceMappingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEventSourceMapping(options.Region), middleware.Before)
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
			OperationName: "CreateEventSourceMapping",
			Err:           err,
		}
	}
	out := result.(*CreateEventSourceMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEventSourceMappingInput struct {

	// Disables the event source mapping to pause polling and invocation.
	Enabled *bool

	// (Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded
	// records.
	DestinationConfig *types.DestinationConfig

	// (Streams) If the function returns an error, split the batch in two and retry.
	BisectBatchOnFunctionError *bool

	// With StartingPosition set to AT_TIMESTAMP, the time from which to start reading.
	StartingPositionTimestamp *time.Time

	// (Streams) The number of batches to process from each shard concurrently.
	ParallelizationFactor *int32

	// The maximum number of items to retrieve in a single batch.
	//
	//     * Amazon Kinesis
	// - Default 100. Max 10,000.
	//
	//     * Amazon DynamoDB Streams - Default 100. Max
	// 1,000.
	//
	//     * Amazon Simple Queue Service - Default 10. Max 10.
	BatchSize *int32

	// (Streams) The maximum number of times to retry when the function returns an
	// error.
	MaximumRetryAttempts *int32

	// The name of the Lambda function. Name formats
	//
	//     * Function name -
	// MyFunction.
	//
	//     * Function ARN -
	// arn:aws:lambda:us-west-2:123456789012:function:MyFunction.
	//
	//     * Version or
	// Alias ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction:PROD.
	//
	//
	// * Partial ARN - 123456789012:function:MyFunction.
	//
	// The length constraint applies
	// only to the full ARN. If you specify only the function name, it's limited to 64
	// characters in length.
	//
	// This member is required.
	FunctionName *string

	// (Streams) The maximum age of a record that Lambda sends to a function for
	// processing.
	MaximumRecordAgeInSeconds *int32

	// The position in a stream from which to start reading. Required for Amazon
	// Kinesis and Amazon DynamoDB Streams sources. AT_TIMESTAMP is only supported for
	// Amazon Kinesis streams.
	StartingPosition types.EventSourcePosition

	// (Streams) The maximum amount of time to gather records before invoking the
	// function, in seconds.
	MaximumBatchingWindowInSeconds *int32

	// The Amazon Resource Name (ARN) of the event source.
	//
	//     * Amazon Kinesis - The
	// ARN of the data stream or a stream consumer.
	//
	//     * Amazon DynamoDB Streams -
	// The ARN of the stream.
	//
	//     * Amazon Simple Queue Service - The ARN of the
	// queue.
	//
	// This member is required.
	EventSourceArn *string
}

// A mapping between an AWS resource and an AWS Lambda function. See
// CreateEventSourceMapping () for details.
type CreateEventSourceMappingOutput struct {

	// The state of the event source mapping. It can be one of the following: Creating,
	// Enabling, Enabled, Disabling, Disabled, Updating, or Deleting.
	State *string

	// (Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded
	// records.
	DestinationConfig *types.DestinationConfig

	// (Streams) The maximum amount of time to gather records before invoking the
	// function, in seconds.
	MaximumBatchingWindowInSeconds *int32

	// (Streams) The maximum number of times to retry when the function returns an
	// error.
	MaximumRetryAttempts *int32

	// The maximum number of items to retrieve in a single batch.
	BatchSize *int32

	// The result of the last AWS Lambda invocation of your Lambda function.
	LastProcessingResult *string

	// (Streams) The number of batches to process from each shard concurrently.
	ParallelizationFactor *int32

	// (Streams) If the function returns an error, split the batch in two and retry.
	BisectBatchOnFunctionError *bool

	// Indicates whether the last change to the event source mapping was made by a
	// user, or by the Lambda service.
	StateTransitionReason *string

	// (Streams) The maximum age of a record that Lambda sends to a function for
	// processing.
	MaximumRecordAgeInSeconds *int32

	// The date that the event source mapping was last updated, or its state changed.
	LastModified *time.Time

	// The ARN of the Lambda function.
	FunctionArn *string

	// The Amazon Resource Name (ARN) of the event source.
	EventSourceArn *string

	// The identifier of the event source mapping.
	UUID *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateEventSourceMappingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateEventSourceMapping{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateEventSourceMapping{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateEventSourceMapping(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "CreateEventSourceMapping",
	}
}
