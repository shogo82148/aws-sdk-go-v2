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

// Disassociates the specified subnets from the transit gateway multicast domain.
func (c *Client) DisassociateTransitGatewayMulticastDomain(ctx context.Context, params *DisassociateTransitGatewayMulticastDomainInput, optFns ...func(*Options)) (*DisassociateTransitGatewayMulticastDomainOutput, error) {
	stack := middleware.NewStack("DisassociateTransitGatewayMulticastDomain", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDisassociateTransitGatewayMulticastDomainMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateTransitGatewayMulticastDomain(options.Region), middleware.Before)
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
			OperationName: "DisassociateTransitGatewayMulticastDomain",
			Err:           err,
		}
	}
	out := result.(*DisassociateTransitGatewayMulticastDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateTransitGatewayMulticastDomainInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId *string

	// The ID of the attachment.
	TransitGatewayAttachmentId *string

	// The IDs of the subnets;
	SubnetIds []*string
}

type DisassociateTransitGatewayMulticastDomainOutput struct {

	// Information about the association.
	Associations *types.TransitGatewayMulticastDomainAssociations

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDisassociateTransitGatewayMulticastDomainMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDisassociateTransitGatewayMulticastDomain{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDisassociateTransitGatewayMulticastDomain{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateTransitGatewayMulticastDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DisassociateTransitGatewayMulticastDomain",
	}
}
