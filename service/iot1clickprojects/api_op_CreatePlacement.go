// Code generated by smithy-go-codegen DO NOT EDIT.

package iot1clickprojects

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an empty placement.
func (c *Client) CreatePlacement(ctx context.Context, params *CreatePlacementInput, optFns ...func(*Options)) (*CreatePlacementOutput, error) {
	stack := middleware.NewStack("CreatePlacement", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreatePlacementMiddlewares(stack)
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
	addOpCreatePlacementValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePlacement(options.Region), middleware.Before)
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
			OperationName: "CreatePlacement",
			Err:           err,
		}
	}
	out := result.(*CreatePlacementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePlacementInput struct {

	// The name of the project in which to create the placement.
	//
	// This member is required.
	ProjectName *string

	// The name of the placement to be created.
	//
	// This member is required.
	PlacementName *string

	// Optional user-defined key/value pairs providing contextual data (such as
	// location or function) for the placement.
	Attributes map[string]*string
}

type CreatePlacementOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreatePlacementMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreatePlacement{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePlacement{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreatePlacement(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot1click",
		OperationName: "CreatePlacement",
	}
}
