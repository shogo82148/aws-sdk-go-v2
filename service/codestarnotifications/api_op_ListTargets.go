// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarnotifications

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of the notification rule targets for an AWS account.
func (c *Client) ListTargets(ctx context.Context, params *ListTargetsInput, optFns ...func(*Options)) (*ListTargetsOutput, error) {
	stack := middleware.NewStack("ListTargets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListTargetsMiddlewares(stack)
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
	addOpListTargetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTargets(options.Region), middleware.Before)
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
			OperationName: "ListTargets",
			Err:           err,
		}
	}
	out := result.(*ListTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTargetsInput struct {

	// The filters to use to return information by service or resource type. Valid
	// filters include target type, target address, and target status. A filter with
	// the same name can appear more than once when used with OR statements. Filters
	// with different names should be applied with AND statements.
	Filters []*types.ListTargetsFilter

	// An enumeration token that, when provided in a request, returns the next batch of
	// the results.
	NextToken *string

	// A non-negative integer used to limit the number of returned results. The maximum
	// number of results that can be returned is 100.
	MaxResults *int32
}

type ListTargetsOutput struct {

	// An enumeration token that can be used in a request to return the next batch of
	// results.
	NextToken *string

	// The list of notification rule targets.
	Targets []*types.TargetSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListTargetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListTargets{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListTargets{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTargets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar-notifications",
		OperationName: "ListTargets",
	}
}
