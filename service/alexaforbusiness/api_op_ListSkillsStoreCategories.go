// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all categories in the Alexa skill store.
func (c *Client) ListSkillsStoreCategories(ctx context.Context, params *ListSkillsStoreCategoriesInput, optFns ...func(*Options)) (*ListSkillsStoreCategoriesOutput, error) {
	stack := middleware.NewStack("ListSkillsStoreCategories", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListSkillsStoreCategoriesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSkillsStoreCategories(options.Region), middleware.Before)
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
			OperationName: "ListSkillsStoreCategories",
			Err:           err,
		}
	}
	out := result.(*ListSkillsStoreCategoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSkillsStoreCategoriesInput struct {

	// The maximum number of categories returned, per paginated calls.
	MaxResults *int32

	// The tokens used for pagination.
	NextToken *string
}

type ListSkillsStoreCategoriesOutput struct {

	// The list of categories.
	CategoryList []*types.Category

	// The tokens used for pagination.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListSkillsStoreCategoriesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListSkillsStoreCategories{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSkillsStoreCategories{}, middleware.After)
}

func newServiceMetadataMiddleware_opListSkillsStoreCategories(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "ListSkillsStoreCategories",
	}
}
