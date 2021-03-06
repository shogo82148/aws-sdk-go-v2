// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Replaces the current set of policies for the specified load balancer port with
// the specified set of policies. To enable back-end server authentication, use
// SetLoadBalancerPoliciesForBackendServer (). For more information about setting
// policies, see Update the SSL Negotiation Configuration
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/ssl-config-update.html),
// Duration-Based Session Stickiness
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-sticky-sessions.html#enable-sticky-sessions-duration),
// and Application-Controlled Session Stickiness
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-sticky-sessions.html#enable-sticky-sessions-application)
// in the Classic Load Balancers Guide.
func (c *Client) SetLoadBalancerPoliciesOfListener(ctx context.Context, params *SetLoadBalancerPoliciesOfListenerInput, optFns ...func(*Options)) (*SetLoadBalancerPoliciesOfListenerOutput, error) {
	stack := middleware.NewStack("SetLoadBalancerPoliciesOfListener", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpSetLoadBalancerPoliciesOfListenerMiddlewares(stack)
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
	addOpSetLoadBalancerPoliciesOfListenerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetLoadBalancerPoliciesOfListener(options.Region), middleware.Before)
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
			OperationName: "SetLoadBalancerPoliciesOfListener",
			Err:           err,
		}
	}
	out := result.(*SetLoadBalancerPoliciesOfListenerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for SetLoadBalancePoliciesOfListener.
type SetLoadBalancerPoliciesOfListenerInput struct {

	// The external port of the load balancer.
	//
	// This member is required.
	LoadBalancerPort *int32

	// The names of the policies. This list must include all policies to be enabled. If
	// you omit a policy that is currently enabled, it is disabled. If the list is
	// empty, all current policies are disabled.
	//
	// This member is required.
	PolicyNames []*string

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string
}

// Contains the output of SetLoadBalancePoliciesOfListener.
type SetLoadBalancerPoliciesOfListenerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpSetLoadBalancerPoliciesOfListenerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpSetLoadBalancerPoliciesOfListener{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpSetLoadBalancerPoliciesOfListener{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetLoadBalancerPoliciesOfListener(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "SetLoadBalancerPoliciesOfListener",
	}
}
