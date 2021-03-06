// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds the specified instances to the specified load balancer.  <p>The instance
// must be a running instance in the same network as the load balancer (EC2-Classic
// or the same VPC). If you have EC2-Classic instances and a load balancer in a VPC
// with ClassicLink enabled, you can link the EC2-Classic instances to that VPC and
// then register the linked EC2-Classic instances with the load balancer in the
// VPC.</p> <p>Note that <code>RegisterInstanceWithLoadBalancer</code> completes
// when the request has been registered. Instance registration takes a little time
// to complete. To check the state of the registered instances, use
// <a>DescribeLoadBalancers</a> or <a>DescribeInstanceHealth</a>.</p> <p>After the
// instance is registered, it starts receiving traffic and requests from the load
// balancer. Any instance that is not in one of the Availability Zones registered
// for the load balancer is moved to the <code>OutOfService</code> state. If an
// Availability Zone is added to the load balancer later, any instances registered
// with the load balancer move to the <code>InService</code> state.</p> <p>To
// deregister instances from a load balancer, use
// <a>DeregisterInstancesFromLoadBalancer</a>.</p> <p>For more information, see <a
// href="https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-deregister-register-instances.html">Register
// or De-Register EC2 Instances</a> in the <i>Classic Load Balancers Guide</i>.</p>
func (c *Client) RegisterInstancesWithLoadBalancer(ctx context.Context, params *RegisterInstancesWithLoadBalancerInput, optFns ...func(*Options)) (*RegisterInstancesWithLoadBalancerOutput, error) {
	stack := middleware.NewStack("RegisterInstancesWithLoadBalancer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRegisterInstancesWithLoadBalancerMiddlewares(stack)
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
	addOpRegisterInstancesWithLoadBalancerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterInstancesWithLoadBalancer(options.Region), middleware.Before)
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
			OperationName: "RegisterInstancesWithLoadBalancer",
			Err:           err,
		}
	}
	out := result.(*RegisterInstancesWithLoadBalancerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for RegisterInstancesWithLoadBalancer.
type RegisterInstancesWithLoadBalancerInput struct {

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	// The IDs of the instances.
	//
	// This member is required.
	Instances []*types.Instance
}

// Contains the output of RegisterInstancesWithLoadBalancer.
type RegisterInstancesWithLoadBalancerOutput struct {

	// The updated list of instances for the load balancer.
	Instances []*types.Instance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRegisterInstancesWithLoadBalancerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRegisterInstancesWithLoadBalancer{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRegisterInstancesWithLoadBalancer{}, middleware.After)
}

func newServiceMetadataMiddleware_opRegisterInstancesWithLoadBalancer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "RegisterInstancesWithLoadBalancer",
	}
}
