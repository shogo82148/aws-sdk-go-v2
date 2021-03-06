// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the resources associated with the specified TagOption.
func (c *Client) ListResourcesForTagOption(ctx context.Context, params *ListResourcesForTagOptionInput, optFns ...func(*Options)) (*ListResourcesForTagOptionOutput, error) {
	stack := middleware.NewStack("ListResourcesForTagOption", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListResourcesForTagOptionMiddlewares(stack)
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
	addOpListResourcesForTagOptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListResourcesForTagOption(options.Region), middleware.Before)
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
			OperationName: "ListResourcesForTagOption",
			Err:           err,
		}
	}
	out := result.(*ListResourcesForTagOptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourcesForTagOptionInput struct {

	// The TagOption identifier.
	//
	// This member is required.
	TagOptionId *string

	// The maximum number of items to return with this call.
	PageSize *int32

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string

	// The resource type.
	//
	//     * Portfolio
	//
	//     * Product
	ResourceType *string
}

type ListResourcesForTagOptionOutput struct {

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string

	// Information about the resources.
	ResourceDetails []*types.ResourceDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListResourcesForTagOptionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListResourcesForTagOption{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResourcesForTagOption{}, middleware.After)
}

func newServiceMetadataMiddleware_opListResourcesForTagOption(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ListResourcesForTagOption",
	}
}
