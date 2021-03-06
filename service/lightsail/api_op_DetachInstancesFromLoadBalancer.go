// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Detaches the specified instances from a Lightsail load balancer. This operation
// waits until the instances are no longer needed before they are detached from the
// load balancer. The detach instances from load balancer operation supports
// tag-based access control via resource tags applied to the resource identified by
// load balancer name. For more information, see the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) DetachInstancesFromLoadBalancer(ctx context.Context, params *DetachInstancesFromLoadBalancerInput, optFns ...func(*Options)) (*DetachInstancesFromLoadBalancerOutput, error) {
	stack := middleware.NewStack("DetachInstancesFromLoadBalancer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDetachInstancesFromLoadBalancerMiddlewares(stack)
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
	addOpDetachInstancesFromLoadBalancerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetachInstancesFromLoadBalancer(options.Region), middleware.Before)
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
			OperationName: "DetachInstancesFromLoadBalancer",
			Err:           err,
		}
	}
	out := result.(*DetachInstancesFromLoadBalancerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachInstancesFromLoadBalancerInput struct {

	// An array of strings containing the names of the instances you want to detach
	// from the load balancer.
	//
	// This member is required.
	InstanceNames []*string

	// The name of the Lightsail load balancer.
	//
	// This member is required.
	LoadBalancerName *string
}

type DetachInstancesFromLoadBalancerOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDetachInstancesFromLoadBalancerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDetachInstancesFromLoadBalancer{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetachInstancesFromLoadBalancer{}, middleware.After)
}

func newServiceMetadataMiddleware_opDetachInstancesFromLoadBalancer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "DetachInstancesFromLoadBalancer",
	}
}
