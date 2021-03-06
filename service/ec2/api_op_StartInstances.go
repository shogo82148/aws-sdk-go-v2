// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts an Amazon EBS-backed instance that you've previously stopped. Instances
// that use Amazon EBS volumes as their root devices can be quickly stopped and
// started. When an instance is stopped, the compute resources are released and you
// are not billed for instance usage. However, your root partition Amazon EBS
// volume remains and continues to persist your data, and you are charged for
// Amazon EBS volume usage. You can restart your instance at any time. Every time
// you start your Windows instance, Amazon EC2 charges you for a full instance
// hour. If you stop and restart your Windows instance, a new instance hour begins
// and Amazon EC2 charges you for another full instance hour even if you are still
// within the same 60-minute period when it was stopped. Every time you start your
// Linux instance, Amazon EC2 charges a one-minute minimum for instance usage, and
// thereafter charges per second for instance usage. Before stopping an instance,
// make sure it is in a state from which it can be restarted. Stopping an instance
// does not preserve data stored in RAM. Performing this operation on an instance
// that uses an instance store as its root device returns an error. For more
// information, see Stopping instances
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Stop_Start.html) in the
// Amazon Elastic Compute Cloud User Guide.
func (c *Client) StartInstances(ctx context.Context, params *StartInstancesInput, optFns ...func(*Options)) (*StartInstancesOutput, error) {
	stack := middleware.NewStack("StartInstances", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpStartInstancesMiddlewares(stack)
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
	addOpStartInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartInstances(options.Region), middleware.Before)
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
			OperationName: "StartInstances",
			Err:           err,
		}
	}
	out := result.(*StartInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartInstancesInput struct {

	// The IDs of the instances.
	//
	// This member is required.
	InstanceIds []*string

	// Reserved.
	AdditionalInfo *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type StartInstancesOutput struct {

	// Information about the started instances.
	StartingInstances []*types.InstanceStateChange

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpStartInstancesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpStartInstances{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpStartInstances{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "StartInstances",
	}
}
