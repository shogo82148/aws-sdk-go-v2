// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the things in the specified group.
func (c *Client) ListThingsInThingGroup(ctx context.Context, params *ListThingsInThingGroupInput, optFns ...func(*Options)) (*ListThingsInThingGroupOutput, error) {
	stack := middleware.NewStack("ListThingsInThingGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListThingsInThingGroupMiddlewares(stack)
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
	addOpListThingsInThingGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListThingsInThingGroup(options.Region), middleware.Before)
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
			OperationName: "ListThingsInThingGroup",
			Err:           err,
		}
	}
	out := result.(*ListThingsInThingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListThingsInThingGroupInput struct {

	// The token to retrieve the next set of results.
	NextToken *string

	// When true, list things in this thing group and in all child groups as well.
	Recursive *bool

	// The thing group name.
	//
	// This member is required.
	ThingGroupName *string

	// The maximum number of results to return at one time.
	MaxResults *int32
}

type ListThingsInThingGroupOutput struct {

	// The things in the specified thing group.
	Things []*string

	// The token used to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListThingsInThingGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListThingsInThingGroup{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListThingsInThingGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opListThingsInThingGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListThingsInThingGroup",
	}
}
