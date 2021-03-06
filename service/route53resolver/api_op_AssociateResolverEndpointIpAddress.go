// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds IP addresses to an inbound or an outbound resolver endpoint. If you want to
// adding more than one IP address, submit one AssociateResolverEndpointIpAddress
// request for each IP address. To remove an IP address from an endpoint, see
// DisassociateResolverEndpointIpAddress ().
func (c *Client) AssociateResolverEndpointIpAddress(ctx context.Context, params *AssociateResolverEndpointIpAddressInput, optFns ...func(*Options)) (*AssociateResolverEndpointIpAddressOutput, error) {
	stack := middleware.NewStack("AssociateResolverEndpointIpAddress", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAssociateResolverEndpointIpAddressMiddlewares(stack)
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
	addOpAssociateResolverEndpointIpAddressValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateResolverEndpointIpAddress(options.Region), middleware.Before)
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
			OperationName: "AssociateResolverEndpointIpAddress",
			Err:           err,
		}
	}
	out := result.(*AssociateResolverEndpointIpAddressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateResolverEndpointIpAddressInput struct {

	// Either the IPv4 address that you want to add to a resolver endpoint or a subnet
	// ID. If you specify a subnet ID, Resolver chooses an IP address for you from the
	// available IPs in the specified subnet.
	//
	// This member is required.
	IpAddress *types.IpAddressUpdate

	// The ID of the resolver endpoint that you want to associate IP addresses with.
	//
	// This member is required.
	ResolverEndpointId *string
}

type AssociateResolverEndpointIpAddressOutput struct {

	// The response to an AssociateResolverEndpointIpAddress request.
	ResolverEndpoint *types.ResolverEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAssociateResolverEndpointIpAddressMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateResolverEndpointIpAddress{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateResolverEndpointIpAddress{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateResolverEndpointIpAddress(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "AssociateResolverEndpointIpAddress",
	}
}
