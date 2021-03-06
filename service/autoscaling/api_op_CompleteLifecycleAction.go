// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Completes the lifecycle action for the specified token or instance with the
// specified result. This step is a part of the procedure for adding a lifecycle
// hook to an Auto Scaling group:
//
//     * (Optional) Create a Lambda function and a
// rule that allows CloudWatch Events to invoke your Lambda function when Amazon
// EC2 Auto Scaling launches or terminates instances.
//
//     * (Optional) Create a
// notification target and an IAM role. The target can be either an Amazon SQS
// queue or an Amazon SNS topic. The role allows Amazon EC2 Auto Scaling to publish
// lifecycle notifications to the target.
//
//     * Create the lifecycle hook. Specify
// whether the hook is used when the instances launch or terminate.
//
//     * If you
// need more time, record the lifecycle action heartbeat to keep the instance in a
// pending state.
//
//     * If you finish before the timeout period ends, complete the
// lifecycle action.
//
// For more information, see Amazon EC2 Auto Scaling Lifecycle
// Hooks
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/lifecycle-hooks.html) in
// the Amazon EC2 Auto Scaling User Guide.
func (c *Client) CompleteLifecycleAction(ctx context.Context, params *CompleteLifecycleActionInput, optFns ...func(*Options)) (*CompleteLifecycleActionOutput, error) {
	stack := middleware.NewStack("CompleteLifecycleAction", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCompleteLifecycleActionMiddlewares(stack)
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
	addOpCompleteLifecycleActionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCompleteLifecycleAction(options.Region), middleware.Before)
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
			OperationName: "CompleteLifecycleAction",
			Err:           err,
		}
	}
	out := result.(*CompleteLifecycleActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CompleteLifecycleActionInput struct {

	// A universally unique identifier (UUID) that identifies a specific lifecycle
	// action associated with an instance. Amazon EC2 Auto Scaling sends this token to
	// the notification target you specified when you created the lifecycle hook.
	LifecycleActionToken *string

	// The name of the lifecycle hook.
	//
	// This member is required.
	LifecycleHookName *string

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The action for the group to take. This parameter can be either CONTINUE or
	// ABANDON.
	//
	// This member is required.
	LifecycleActionResult *string

	// The ID of the instance.
	InstanceId *string
}

type CompleteLifecycleActionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCompleteLifecycleActionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCompleteLifecycleAction{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCompleteLifecycleAction{}, middleware.After)
}

func newServiceMetadataMiddleware_opCompleteLifecycleAction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "CompleteLifecycleAction",
	}
}
