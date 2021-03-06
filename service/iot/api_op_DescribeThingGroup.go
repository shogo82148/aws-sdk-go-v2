// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describe a thing group.
func (c *Client) DescribeThingGroup(ctx context.Context, params *DescribeThingGroupInput, optFns ...func(*Options)) (*DescribeThingGroupOutput, error) {
	stack := middleware.NewStack("DescribeThingGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeThingGroupMiddlewares(stack)
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
	addOpDescribeThingGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeThingGroup(options.Region), middleware.Before)
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
			OperationName: "DescribeThingGroup",
			Err:           err,
		}
	}
	out := result.(*DescribeThingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeThingGroupInput struct {

	// The name of the thing group.
	//
	// This member is required.
	ThingGroupName *string
}

type DescribeThingGroupOutput struct {

	// The dynamic thing group query version.
	QueryVersion *string

	// The thing group ID.
	ThingGroupId *string

	// The dynamic thing group index name.
	IndexName *string

	// The thing group ARN.
	ThingGroupArn *string

	// The name of the thing group.
	ThingGroupName *string

	// Thing group metadata.
	ThingGroupMetadata *types.ThingGroupMetadata

	// The thing group properties.
	ThingGroupProperties *types.ThingGroupProperties

	// The dynamic thing group status.
	Status types.DynamicGroupStatus

	// The dynamic thing group search query string.
	QueryString *string

	// The version of the thing group.
	Version *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeThingGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeThingGroup{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeThingGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeThingGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeThingGroup",
	}
}
