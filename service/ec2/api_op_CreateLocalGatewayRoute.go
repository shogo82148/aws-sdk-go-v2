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

// Creates a static route for the specified local gateway route table.
func (c *Client) CreateLocalGatewayRoute(ctx context.Context, params *CreateLocalGatewayRouteInput, optFns ...func(*Options)) (*CreateLocalGatewayRouteOutput, error) {
	stack := middleware.NewStack("CreateLocalGatewayRoute", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCreateLocalGatewayRouteMiddlewares(stack)
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
	addOpCreateLocalGatewayRouteValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocalGatewayRoute(options.Region), middleware.Before)
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
			OperationName: "CreateLocalGatewayRoute",
			Err:           err,
		}
	}
	out := result.(*CreateLocalGatewayRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLocalGatewayRouteInput struct {

	// The ID of the virtual interface group.
	//
	// This member is required.
	LocalGatewayVirtualInterfaceGroupId *string

	// The ID of the local gateway route table.
	//
	// This member is required.
	LocalGatewayRouteTableId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The CIDR range used for destination matches. Routing decisions are based on the
	// most specific match.
	//
	// This member is required.
	DestinationCidrBlock *string
}

type CreateLocalGatewayRouteOutput struct {

	// Information about the route.
	Route *types.LocalGatewayRoute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCreateLocalGatewayRouteMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCreateLocalGatewayRoute{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCreateLocalGatewayRoute{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateLocalGatewayRoute(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateLocalGatewayRoute",
	}
}
