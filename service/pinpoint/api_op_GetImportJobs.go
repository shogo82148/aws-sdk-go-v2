// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about the status and settings of all the import jobs for
// an application.
func (c *Client) GetImportJobs(ctx context.Context, params *GetImportJobsInput, optFns ...func(*Options)) (*GetImportJobsOutput, error) {
	stack := middleware.NewStack("GetImportJobs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetImportJobsMiddlewares(stack)
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
	addOpGetImportJobsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetImportJobs(options.Region), middleware.Before)
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
			OperationName: "GetImportJobs",
			Err:           err,
		}
	}
	out := result.(*GetImportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetImportJobsInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string

	// The NextToken string that specifies which page of results to return in a
	// paginated response.
	Token *string

	// The maximum number of items to include in each page of a paginated response.
	// This parameter is not supported for application, campaign, and journey metrics.
	PageSize *string
}

type GetImportJobsOutput struct {

	// Provides information about the status and settings of all the import jobs that
	// are associated with an application or segment. An import job is a job that
	// imports endpoint definitions from one or more files.
	//
	// This member is required.
	ImportJobsResponse *types.ImportJobsResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetImportJobsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetImportJobs{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetImportJobs{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetImportJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "GetImportJobs",
	}
}
