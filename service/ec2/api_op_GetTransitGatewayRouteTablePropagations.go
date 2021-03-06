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

// Gets information about the route table propagations for the specified transit
// gateway route table.
func (c *Client) GetTransitGatewayRouteTablePropagations(ctx context.Context, params *GetTransitGatewayRouteTablePropagationsInput, optFns ...func(*Options)) (*GetTransitGatewayRouteTablePropagationsOutput, error) {
	stack := middleware.NewStack("GetTransitGatewayRouteTablePropagations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpGetTransitGatewayRouteTablePropagationsMiddlewares(stack)
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
	addOpGetTransitGatewayRouteTablePropagationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTransitGatewayRouteTablePropagations(options.Region), middleware.Before)
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
			OperationName: "GetTransitGatewayRouteTablePropagations",
			Err:           err,
		}
	}
	out := result.(*GetTransitGatewayRouteTablePropagationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTransitGatewayRouteTablePropagationsInput struct {

	// The ID of the transit gateway route table.
	//
	// This member is required.
	TransitGatewayRouteTableId *string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The token for the next page of results.
	NextToken *string

	// One or more filters. The possible values are:
	//
	//     * resource-id - The ID of the
	// resource.
	//
	//     * resource-type - The resource type (vpc | vpn).
	//
	//     *
	// transit-gateway-attachment-id - The ID of the attachment.
	Filters []*types.Filter
}

type GetTransitGatewayRouteTablePropagationsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the route table propagations.
	TransitGatewayRouteTablePropagations []*types.TransitGatewayRouteTablePropagation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpGetTransitGatewayRouteTablePropagationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpGetTransitGatewayRouteTablePropagations{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpGetTransitGatewayRouteTablePropagations{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetTransitGatewayRouteTablePropagations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetTransitGatewayRouteTablePropagations",
	}
}
