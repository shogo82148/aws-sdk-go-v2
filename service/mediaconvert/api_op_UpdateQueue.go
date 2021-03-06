// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconvert

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modify one of your existing queues.
func (c *Client) UpdateQueue(ctx context.Context, params *UpdateQueueInput, optFns ...func(*Options)) (*UpdateQueueOutput, error) {
	stack := middleware.NewStack("UpdateQueue", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateQueueMiddlewares(stack)
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
	addOpUpdateQueueValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateQueue(options.Region), middleware.Before)
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
			OperationName: "UpdateQueue",
			Err:           err,
		}
	}
	out := result.(*UpdateQueueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateQueueInput struct {

	// The new description for the queue, if you are changing it.
	Description *string

	// Pause or activate a queue by changing its status between ACTIVE and PAUSED. If
	// you pause a queue, jobs in that queue won't begin. Jobs that are running when
	// you pause the queue continue to run until they finish or result in an error.
	Status types.QueueStatus

	// The new details of your pricing plan for your reserved queue. When you set up a
	// new pricing plan to replace an expired one, you enter into another 12-month
	// commitment. When you add capacity to your queue by increasing the number of RTS,
	// you extend the term of your commitment to 12 months from when you add capacity.
	// After you make these commitments, you can't cancel them.
	ReservationPlanSettings *types.ReservationPlanSettings

	// The name of the queue that you are modifying.
	//
	// This member is required.
	Name *string
}

type UpdateQueueOutput struct {

	// You can use queues to manage the resources that are available to your AWS
	// account for running multiple transcoding jobs at the same time. If you don't
	// specify a queue, the service sends all jobs through the default queue. For more
	// information, see
	// https://docs.aws.amazon.com/mediaconvert/latest/ug/working-with-queues.html.
	Queue *types.Queue

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateQueueMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateQueue{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateQueue{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateQueue(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconvert",
		OperationName: "UpdateQueue",
	}
}
