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

// Accepts one or more interface VPC endpoint connection requests to your VPC
// endpoint service.
func (c *Client) AcceptVpcEndpointConnections(ctx context.Context, params *AcceptVpcEndpointConnectionsInput, optFns ...func(*Options)) (*AcceptVpcEndpointConnectionsOutput, error) {
	stack := middleware.NewStack("AcceptVpcEndpointConnections", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpAcceptVpcEndpointConnectionsMiddlewares(stack)
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
	addOpAcceptVpcEndpointConnectionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptVpcEndpointConnections(options.Region), middleware.Before)
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
			OperationName: "AcceptVpcEndpointConnections",
			Err:           err,
		}
	}
	out := result.(*AcceptVpcEndpointConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AcceptVpcEndpointConnectionsInput struct {

	// The ID of the VPC endpoint service.
	//
	// This member is required.
	ServiceId *string

	// The IDs of one or more interface VPC endpoints.
	//
	// This member is required.
	VpcEndpointIds []*string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type AcceptVpcEndpointConnectionsOutput struct {

	// Information about the interface endpoints that were not accepted, if applicable.
	Unsuccessful []*types.UnsuccessfulItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpAcceptVpcEndpointConnectionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpAcceptVpcEndpointConnections{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpAcceptVpcEndpointConnections{}, middleware.After)
}

func newServiceMetadataMiddleware_opAcceptVpcEndpointConnections(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AcceptVpcEndpointConnections",
	}
}
