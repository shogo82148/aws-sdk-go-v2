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

// Sets the certificate that terminates the specified listener's SSL connections.
// The specified certificate replaces any prior certificate that was used on the
// same load balancer and port.  <p>For more information about updating your SSL
// certificate, see <a
// href="https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-update-ssl-cert.html">Replace
// the SSL Certificate for Your Load Balancer</a> in the <i>Classic Load Balancers
// Guide</i>.</p>
func (c *Client) SetLoadBalancerListenerSSLCertificate(ctx context.Context, params *SetLoadBalancerListenerSSLCertificateInput, optFns ...func(*Options)) (*SetLoadBalancerListenerSSLCertificateOutput, error) {
	stack := middleware.NewStack("SetLoadBalancerListenerSSLCertificate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpSetLoadBalancerListenerSSLCertificateMiddlewares(stack)
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
	addOpSetLoadBalancerListenerSSLCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetLoadBalancerListenerSSLCertificate(options.Region), middleware.Before)
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
			OperationName: "SetLoadBalancerListenerSSLCertificate",
			Err:           err,
		}
	}
	out := result.(*SetLoadBalancerListenerSSLCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for SetLoadBalancerListenerSSLCertificate.
type SetLoadBalancerListenerSSLCertificateInput struct {

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	// The Amazon Resource Name (ARN) of the SSL certificate.
	//
	// This member is required.
	SSLCertificateId *string

	// The port that uses the specified SSL certificate.
	//
	// This member is required.
	LoadBalancerPort *int32
}

// Contains the output of SetLoadBalancerListenerSSLCertificate.
type SetLoadBalancerListenerSSLCertificateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpSetLoadBalancerListenerSSLCertificateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpSetLoadBalancerListenerSSLCertificate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpSetLoadBalancerListenerSSLCertificate{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetLoadBalancerListenerSSLCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "SetLoadBalancerListenerSSLCertificate",
	}
}
