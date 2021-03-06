// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists the machine learning algorithms that have been created.
func (c *Client) ListAlgorithms(ctx context.Context, params *ListAlgorithmsInput, optFns ...func(*Options)) (*ListAlgorithmsOutput, error) {
	stack := middleware.NewStack("ListAlgorithms", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListAlgorithmsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAlgorithms(options.Region), middleware.Before)
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
			OperationName: "ListAlgorithms",
			Err:           err,
		}
	}
	out := result.(*ListAlgorithmsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAlgorithmsInput struct {

	// The sort order for the results. The default is Ascending.
	SortOrder types.SortOrder

	// A string in the algorithm name. This filter returns only algorithms whose name
	// contains the specified string.
	NameContains *string

	// A filter that returns only algorithms created before the specified time
	// (timestamp).
	CreationTimeBefore *time.Time

	// A filter that returns only algorithms created after the specified time
	// (timestamp).
	CreationTimeAfter *time.Time

	// The maximum number of algorithms to return in the response.
	MaxResults *int32

	// If the response to a previous ListAlgorithms request was truncated, the response
	// includes a NextToken. To retrieve the next set of algorithms, use the token in
	// the next request.
	NextToken *string

	// The parameter by which to sort the results. The default is CreationTime.
	SortBy types.AlgorithmSortBy
}

type ListAlgorithmsOutput struct {

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of algorithms, use it in the subsequent request.
	NextToken *string

	// >An array of AlgorithmSummary objects, each of which lists an algorithm.
	//
	// This member is required.
	AlgorithmSummaryList []*types.AlgorithmSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListAlgorithmsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListAlgorithms{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAlgorithms{}, middleware.After)
}

func newServiceMetadataMiddleware_opListAlgorithms(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListAlgorithms",
	}
}
