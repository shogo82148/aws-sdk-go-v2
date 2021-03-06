// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the link associations for a device or a link. Either the device ID or the
// link ID must be specified.
func (c *Client) GetLinkAssociations(ctx context.Context, params *GetLinkAssociationsInput, optFns ...func(*Options)) (*GetLinkAssociationsOutput, error) {
	stack := middleware.NewStack("GetLinkAssociations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetLinkAssociationsMiddlewares(stack)
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
	addOpGetLinkAssociationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetLinkAssociations(options.Region), middleware.Before)
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
			OperationName: "GetLinkAssociations",
			Err:           err,
		}
	}
	out := result.(*GetLinkAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLinkAssociationsInput struct {

	// The token for the next page of results.
	NextToken *string

	// The maximum number of results to return.
	MaxResults *int32

	// The ID of the device.
	DeviceId *string

	// The ID of the link.
	LinkId *string

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string
}

type GetLinkAssociationsOutput struct {

	// The link associations.
	LinkAssociations []*types.LinkAssociation

	// The token for the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetLinkAssociationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetLinkAssociations{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetLinkAssociations{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetLinkAssociations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "GetLinkAssociations",
	}
}
