// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provisions an IP address range to use with your AWS resources through bring your
// own IP addresses (BYOIP) and creates a corresponding address pool. After the
// address range is provisioned, it is ready to be advertised using
// AdvertiseByoipCidr
// (https://docs.aws.amazon.com/global-accelerator/latest/api/AdvertiseByoipCidr.html).
// To see an AWS CLI example of provisioning an address range for BYOIP, scroll
// down to Example. For more information, see Bring Your Own IP Addresses (BYOIP)
// (https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html) in
// the AWS Global Accelerator Developer Guide.
func (c *Client) ProvisionByoipCidr(ctx context.Context, params *ProvisionByoipCidrInput, optFns ...func(*Options)) (*ProvisionByoipCidrOutput, error) {
	stack := middleware.NewStack("ProvisionByoipCidr", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpProvisionByoipCidrMiddlewares(stack)
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
	addOpProvisionByoipCidrValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opProvisionByoipCidr(options.Region), middleware.Before)
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
			OperationName: "ProvisionByoipCidr",
			Err:           err,
		}
	}
	out := result.(*ProvisionByoipCidrOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ProvisionByoipCidrInput struct {

	// The public IPv4 address range, in CIDR notation. The most specific IP prefix
	// that you can specify is /24. The address range cannot overlap with another
	// address range that you've brought to this or another Region.
	//
	// This member is required.
	Cidr *string

	// A signed document that proves that you are authorized to bring the specified IP
	// address range to Amazon using BYOIP.
	//
	// This member is required.
	CidrAuthorizationContext *types.CidrAuthorizationContext
}

type ProvisionByoipCidrOutput struct {

	// Information about the address range.
	ByoipCidr *types.ByoipCidr

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpProvisionByoipCidrMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpProvisionByoipCidr{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpProvisionByoipCidr{}, middleware.After)
}

func newServiceMetadataMiddleware_opProvisionByoipCidr(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "ProvisionByoipCidr",
	}
}
