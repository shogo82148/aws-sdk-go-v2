// Code generated by smithy-go-codegen DO NOT EDIT.

package iot1clickprojects

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot1clickprojects/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes a placement in a project.
func (c *Client) DescribePlacement(ctx context.Context, params *DescribePlacementInput, optFns ...func(*Options)) (*DescribePlacementOutput, error) {
	stack := middleware.NewStack("DescribePlacement", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribePlacementMiddlewares(stack)
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
	addOpDescribePlacementValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePlacement(options.Region), middleware.Before)
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
			OperationName: "DescribePlacement",
			Err:           err,
		}
	}
	out := result.(*DescribePlacementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePlacementInput struct {

	// The name of the placement within a project.
	//
	// This member is required.
	PlacementName *string

	// The project containing the placement to be described.
	//
	// This member is required.
	ProjectName *string
}

type DescribePlacementOutput struct {

	// An object describing the placement.
	//
	// This member is required.
	Placement *types.PlacementDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribePlacementMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribePlacement{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribePlacement{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribePlacement(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot1click",
		OperationName: "DescribePlacement",
	}
}
