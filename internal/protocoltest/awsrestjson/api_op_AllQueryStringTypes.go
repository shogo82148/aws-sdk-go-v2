// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This example uses all query string types.
func (c *Client) AllQueryStringTypes(ctx context.Context, params *AllQueryStringTypesInput, optFns ...func(*Options)) (*AllQueryStringTypesOutput, error) {
	stack := middleware.NewStack("AllQueryStringTypes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpAllQueryStringTypesMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAllQueryStringTypes(options.Region), middleware.Before)
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
			OperationName: "AllQueryStringTypes",
			Err:           err,
		}
	}
	out := result.(*AllQueryStringTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AllQueryStringTypesInput struct {
	QueryString *string

	QueryStringList []*string

	QueryStringSet []*string

	QueryByte *int8

	QueryShort *int16

	QueryInteger *int32

	QueryIntegerList []*int32

	QueryIntegerSet []*int32

	QueryLong *int64

	QueryFloat *float32

	QueryDouble *float64

	QueryDoubleList []*float64

	QueryBoolean *bool

	QueryBooleanList []*bool

	QueryTimestamp *time.Time

	QueryTimestampList []*time.Time

	QueryEnum types.FooEnum

	QueryEnumList []types.FooEnum
}

type AllQueryStringTypesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpAllQueryStringTypesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpAllQueryStringTypes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAllQueryStringTypes{}, middleware.After)
}

func newServiceMetadataMiddleware_opAllQueryStringTypes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AllQueryStringTypes",
	}
}
