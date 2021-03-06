// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of the batch translation jobs that you have submitted.
func (c *Client) ListTextTranslationJobs(ctx context.Context, params *ListTextTranslationJobsInput, optFns ...func(*Options)) (*ListTextTranslationJobsOutput, error) {
	stack := middleware.NewStack("ListTextTranslationJobs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListTextTranslationJobsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTextTranslationJobs(options.Region), middleware.Before)
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
			OperationName: "ListTextTranslationJobs",
			Err:           err,
		}
	}
	out := result.(*ListTextTranslationJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTextTranslationJobsInput struct {

	// The token to request the next page of results.
	NextToken *string

	// The maximum number of results to return in each page. The default value is 100.
	MaxResults *int32

	// The parameters that specify which batch translation jobs to retrieve. Filters
	// include job name, job status, and submission time. You can only set one filter
	// at a time.
	Filter *types.TextTranslationJobFilter
}

type ListTextTranslationJobsOutput struct {

	// The token to use to retreive the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// A list containing the properties of each job that is returned.
	TextTranslationJobPropertiesList []*types.TextTranslationJobProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListTextTranslationJobsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListTextTranslationJobs{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTextTranslationJobs{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTextTranslationJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "translate",
		OperationName: "ListTextTranslationJobs",
	}
}
