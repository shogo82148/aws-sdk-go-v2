// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Searches for things associated with the specified entity. You can search by both
// device and device model. For example, if two different devices, camera1 and
// camera2, implement the camera device model, the user can associate thing1 to
// camera1 and thing2 to camera2. SearchThings(camera2) will return only thing2,
// but SearchThings(camera) will return both thing1 and thing2. This action
// searches for exact matches and doesn't perform partial text matching.
func (c *Client) SearchThings(ctx context.Context, params *SearchThingsInput, optFns ...func(*Options)) (*SearchThingsOutput, error) {
	stack := middleware.NewStack("SearchThings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSearchThingsMiddlewares(stack)
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
	addOpSearchThingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSearchThings(options.Region), middleware.Before)
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
			OperationName: "SearchThings",
			Err:           err,
		}
	}
	out := result.(*SearchThingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchThingsInput struct {

	// The ID of the entity to which the things are associated. The IDs should be in
	// the following format. urn:tdm:REGION/ACCOUNT ID/default:device:DEVICENAME
	//
	// This member is required.
	EntityId *string

	// The maximum number of results to return in the response.
	MaxResults *int32

	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string

	// The version of the user's namespace. Defaults to the latest version of the
	// user's namespace.
	NamespaceVersion *int64
}

type SearchThingsOutput struct {

	// An array of things in the result set.
	Things []*types.Thing

	// The string to specify as nextToken when you request the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSearchThingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSearchThings{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchThings{}, middleware.After)
}

func newServiceMetadataMiddleware_opSearchThings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "SearchThings",
	}
}
