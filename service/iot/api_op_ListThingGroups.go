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

// List the thing groups in your account.
func (c *Client) ListThingGroups(ctx context.Context, params *ListThingGroupsInput, optFns ...func(*Options)) (*ListThingGroupsOutput, error) {
	stack := middleware.NewStack("ListThingGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListThingGroupsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListThingGroups(options.Region), middleware.Before)
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
			OperationName: "ListThingGroups",
			Err:           err,
		}
	}
	out := result.(*ListThingGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListThingGroupsInput struct {

	// The maximum number of results to return at one time.
	MaxResults *int32

	// The token to retrieve the next set of results.
	NextToken *string

	// A filter that limits the results to those with the specified name prefix.
	NamePrefixFilter *string

	// A filter that limits the results to those with the specified parent group.
	ParentGroup *string

	// If true, return child groups as well.
	Recursive *bool
}

type ListThingGroupsOutput struct {

	// The token used to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// The thing groups.
	ThingGroups []*types.GroupNameAndArn

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListThingGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListThingGroups{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListThingGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opListThingGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListThingGroups",
	}
}
