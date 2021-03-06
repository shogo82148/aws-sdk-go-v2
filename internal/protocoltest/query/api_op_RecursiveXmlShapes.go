// Code generated by smithy-go-codegen DO NOT EDIT.

package query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/query/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Recursive shapes
func (c *Client) RecursiveXmlShapes(ctx context.Context, params *RecursiveXmlShapesInput, optFns ...func(*Options)) (*RecursiveXmlShapesOutput, error) {
	stack := middleware.NewStack("RecursiveXmlShapes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRecursiveXmlShapesMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRecursiveXmlShapes(options.Region), middleware.Before)
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
			OperationName: "RecursiveXmlShapes",
			Err:           err,
		}
	}
	out := result.(*RecursiveXmlShapesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RecursiveXmlShapesInput struct {
}

type RecursiveXmlShapesOutput struct {
	Nested *types.RecursiveXmlShapesOutputNested1

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRecursiveXmlShapesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRecursiveXmlShapes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRecursiveXmlShapes{}, middleware.After)
}

func newServiceMetadataMiddleware_opRecursiveXmlShapes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RecursiveXmlShapes",
	}
}
