// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the major version families of each published schema. If a major version
// ARN is provided as SchemaArn, the minor version revisions in that family are
// listed instead.
func (c *Client) ListPublishedSchemaArns(ctx context.Context, params *ListPublishedSchemaArnsInput, optFns ...func(*Options)) (*ListPublishedSchemaArnsOutput, error) {
	stack := middleware.NewStack("ListPublishedSchemaArns", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListPublishedSchemaArnsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPublishedSchemaArns(options.Region), middleware.Before)
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
			OperationName: "ListPublishedSchemaArns",
			Err:           err,
		}
	}
	out := result.(*ListPublishedSchemaArnsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPublishedSchemaArnsInput struct {

	// The pagination token.
	NextToken *string

	// The maximum number of results to retrieve.
	MaxResults *int32

	// The response for ListPublishedSchemaArns when this parameter is used will list
	// all minor version ARNs for a major version.
	SchemaArn *string
}

type ListPublishedSchemaArnsOutput struct {

	// The pagination token.
	NextToken *string

	// The ARNs of published schemas.
	SchemaArns []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListPublishedSchemaArnsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListPublishedSchemaArns{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListPublishedSchemaArns{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPublishedSchemaArns(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListPublishedSchemaArns",
	}
}
