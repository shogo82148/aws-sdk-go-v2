// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a batch of variables.
func (c *Client) BatchCreateVariable(ctx context.Context, params *BatchCreateVariableInput, optFns ...func(*Options)) (*BatchCreateVariableOutput, error) {
	stack := middleware.NewStack("BatchCreateVariable", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchCreateVariableMiddlewares(stack)
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
	addOpBatchCreateVariableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchCreateVariable(options.Region), middleware.Before)
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
			OperationName: "BatchCreateVariable",
			Err:           err,
		}
	}
	out := result.(*BatchCreateVariableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchCreateVariableInput struct {

	// A collection of key and value pairs.
	Tags []*types.Tag

	// The list of variables for the batch create variable request.
	//
	// This member is required.
	VariableEntries []*types.VariableEntry
}

type BatchCreateVariableOutput struct {

	// Provides the errors for the BatchCreateVariable request.
	Errors []*types.BatchCreateVariableError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchCreateVariableMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchCreateVariable{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchCreateVariable{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchCreateVariable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "BatchCreateVariable",
	}
}
