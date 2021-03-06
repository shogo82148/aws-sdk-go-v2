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

// Lists training jobs.
func (c *Client) ListTrainingJobs(ctx context.Context, params *ListTrainingJobsInput, optFns ...func(*Options)) (*ListTrainingJobsOutput, error) {
	stack := middleware.NewStack("ListTrainingJobs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListTrainingJobsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTrainingJobs(options.Region), middleware.Before)
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
			OperationName: "ListTrainingJobs",
			Err:           err,
		}
	}
	out := result.(*ListTrainingJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTrainingJobsInput struct {

	// A filter that returns only training jobs created before the specified time
	// (timestamp).
	CreationTimeBefore *time.Time

	// The field to sort results by. The default is CreationTime.
	SortBy types.SortBy

	// A filter that returns only training jobs modified before the specified time
	// (timestamp).
	LastModifiedTimeBefore *time.Time

	// The maximum number of training jobs to return in the response.
	MaxResults *int32

	// A filter that returns only training jobs modified after the specified time
	// (timestamp).
	LastModifiedTimeAfter *time.Time

	// A filter that retrieves only training jobs with a specific status.
	StatusEquals types.TrainingJobStatus

	// If the result of the previous ListTrainingJobs request was truncated, the
	// response includes a NextToken. To retrieve the next set of training jobs, use
	// the token in the next request.
	NextToken *string

	// The sort order for results. The default is Ascending.
	SortOrder types.SortOrder

	// A filter that returns only training jobs created after the specified time
	// (timestamp).
	CreationTimeAfter *time.Time

	// A string in the training job name. This filter returns only training jobs whose
	// name contains the specified string.
	NameContains *string
}

type ListTrainingJobsOutput struct {

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of training jobs, use it in the subsequent request.
	NextToken *string

	// An array of TrainingJobSummary objects, each listing a training job.
	//
	// This member is required.
	TrainingJobSummaries []*types.TrainingJobSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListTrainingJobsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListTrainingJobs{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTrainingJobs{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTrainingJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListTrainingJobs",
	}
}
