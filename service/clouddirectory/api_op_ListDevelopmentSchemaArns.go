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

// Retrieves each Amazon Resource Name (ARN) of schemas in the development state.
func (c *Client) ListDevelopmentSchemaArns(ctx context.Context, params *ListDevelopmentSchemaArnsInput, optFns ...func(*Options)) (*ListDevelopmentSchemaArnsOutput, error) {
	stack := middleware.NewStack("ListDevelopmentSchemaArns", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListDevelopmentSchemaArnsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDevelopmentSchemaArns(options.Region), middleware.Before)
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
			OperationName: "ListDevelopmentSchemaArns",
			Err:           err,
		}
	}
	out := result.(*ListDevelopmentSchemaArnsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDevelopmentSchemaArnsInput struct {

	// The pagination token.
	NextToken *string

	// The maximum number of results to retrieve.
	MaxResults *int32
}

type ListDevelopmentSchemaArnsOutput struct {

	// The pagination token.
	NextToken *string

	// The ARNs of retrieved development schemas.
	SchemaArns []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListDevelopmentSchemaArnsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListDevelopmentSchemaArns{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListDevelopmentSchemaArns{}, middleware.After)
}

func newServiceMetadataMiddleware_opListDevelopmentSchemaArns(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListDevelopmentSchemaArns",
	}
}
