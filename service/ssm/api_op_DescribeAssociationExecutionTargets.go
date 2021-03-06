// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Use this API action to view information about a specific execution of a specific
// association.
func (c *Client) DescribeAssociationExecutionTargets(ctx context.Context, params *DescribeAssociationExecutionTargetsInput, optFns ...func(*Options)) (*DescribeAssociationExecutionTargetsOutput, error) {
	stack := middleware.NewStack("DescribeAssociationExecutionTargets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeAssociationExecutionTargetsMiddlewares(stack)
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
	addOpDescribeAssociationExecutionTargetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAssociationExecutionTargets(options.Region), middleware.Before)
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
			OperationName: "DescribeAssociationExecutionTargets",
			Err:           err,
		}
	}
	out := result.(*DescribeAssociationExecutionTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAssociationExecutionTargetsInput struct {

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// Filters for the request. You can specify the following filters and values.
	// Status (EQUAL) ResourceId (EQUAL) ResourceType (EQUAL)
	Filters []*types.AssociationExecutionTargetsFilter

	// The execution ID for which you want to view details.
	//
	// This member is required.
	ExecutionId *string

	// The association ID that includes the execution for which you want to view
	// details.
	//
	// This member is required.
	AssociationId *string

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string
}

type DescribeAssociationExecutionTargetsOutput struct {

	// Information about the execution.
	AssociationExecutionTargets []*types.AssociationExecutionTarget

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeAssociationExecutionTargetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAssociationExecutionTargets{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAssociationExecutionTargets{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAssociationExecutionTargets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeAssociationExecutionTargets",
	}
}
