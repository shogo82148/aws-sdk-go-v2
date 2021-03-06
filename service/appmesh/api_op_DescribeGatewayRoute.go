// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes an existing gateway route.
func (c *Client) DescribeGatewayRoute(ctx context.Context, params *DescribeGatewayRouteInput, optFns ...func(*Options)) (*DescribeGatewayRouteOutput, error) {
	stack := middleware.NewStack("DescribeGatewayRoute", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeGatewayRouteMiddlewares(stack)
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
	addOpDescribeGatewayRouteValidationMiddleware(stack)
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
			OperationName: "DescribeGatewayRoute",
			Err:           err,
		}
	}
	out := result.(*DescribeGatewayRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGatewayRouteInput struct {

	// The name of the gateway route to describe.
	//
	// This member is required.
	GatewayRouteName *string

	// The name of the service mesh that the gateway route resides in.
	//
	// This member is required.
	MeshName *string

	// The name of the virtual gateway that the gateway route is associated with.
	//
	// This member is required.
	VirtualGatewayName *string

	// The AWS IAM account ID of the service mesh owner. If the account ID is not your
	// own, then it's the ID of the account that shared the mesh with your account. For
	// more information about mesh sharing, see Working with shared meshes
	// (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).
	MeshOwner *string
}

type DescribeGatewayRouteOutput struct {

	// The full description of your gateway route.
	//
	// This member is required.
	GatewayRoute *types.GatewayRouteData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeGatewayRouteMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeGatewayRoute{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeGatewayRoute{}, middleware.After)
}
