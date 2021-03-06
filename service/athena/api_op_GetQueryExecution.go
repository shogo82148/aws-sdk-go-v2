// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about a single execution of a query if you have access to
// the workgroup in which the query ran. Each time a query executes, information
// about the query execution is saved with a unique ID.
func (c *Client) GetQueryExecution(ctx context.Context, params *GetQueryExecutionInput, optFns ...func(*Options)) (*GetQueryExecutionOutput, error) {
	stack := middleware.NewStack("GetQueryExecution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetQueryExecutionMiddlewares(stack)
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
	addOpGetQueryExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetQueryExecution(options.Region), middleware.Before)
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
			OperationName: "GetQueryExecution",
			Err:           err,
		}
	}
	out := result.(*GetQueryExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetQueryExecutionInput struct {

	// The unique ID of the query execution.
	//
	// This member is required.
	QueryExecutionId *string
}

type GetQueryExecutionOutput struct {

	// Information about the query execution.
	QueryExecution *types.QueryExecution

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetQueryExecutionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetQueryExecution{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetQueryExecution{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetQueryExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "GetQueryExecution",
	}
}
