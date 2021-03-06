// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation lists jobs for a vault, including jobs that are in-progress and
// jobs that have recently finished. The List Job operation returns a list of these
// jobs sorted by job initiation time.  <note> <p>Amazon Glacier retains recently
// completed jobs for a period before deleting them; however, it eventually removes
// completed jobs. The output of completed jobs can be retrieved. Retaining
// completed jobs for a period of time after they have completed enables you to get
// a job output in the event you miss the job completion notification or your first
// attempt to download it fails. For example, suppose you start an archive
// retrieval job to download an archive. After the job completes, you start to
// download the archive but encounter a network error. In this scenario, you can
// retry and download the archive while the job exists.</p> </note> <p>The List
// Jobs operation supports pagination. You should always check the response
// <code>Marker</code> field. If there are no more jobs to list, the
// <code>Marker</code> field is set to <code>null</code>. If there are more jobs to
// list, the <code>Marker</code> field is set to a non-null value, which you can
// use to continue the pagination of the list. To return a list of jobs that begins
// at a specific job, set the marker request parameter to the <code>Marker</code>
// value for that job that you obtained from a previous List Jobs request.</p>
// <p>You can set a maximum limit for the number of jobs returned in the response
// by specifying the <code>limit</code> parameter in the request. The default limit
// is 50. The number of jobs returned might be fewer than the limit, but the number
// of returned jobs never exceeds the limit.</p> <p>Additionally, you can filter
// the jobs list returned by specifying the optional <code>statuscode</code>
// parameter or <code>completed</code> parameter, or both. Using the
// <code>statuscode</code> parameter, you can specify to return only jobs that
// match either the <code>InProgress</code>, <code>Succeeded</code>, or
// <code>Failed</code> status. Using the <code>completed</code> parameter, you can
// specify to return only jobs that were completed (<code>true</code>) or jobs that
// were not completed (<code>false</code>).</p> <p>For more information about using
// this operation, see the documentation for the underlying REST API <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/api-jobs-get.html">List
// Jobs</a>. </p>
func (c *Client) ListJobs(ctx context.Context, params *ListJobsInput, optFns ...func(*Options)) (*ListJobsOutput, error) {
	stack := middleware.NewStack("ListJobs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListJobsMiddlewares(stack)
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
	addOpListJobsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListJobs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	glaciercust.AddTreeHashMiddleware(stack)
	glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion)
	glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID)

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
			OperationName: "ListJobs",
			Err:           err,
		}
	}
	out := result.(*ListJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options for retrieving a job list for an Amazon S3 Glacier vault.
type ListJobsInput struct {

	// The state of the jobs to return. You can specify true or false.
	Completed *string

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// An opaque string used for pagination. This value specifies the job at which the
	// listing of jobs should begin. Get the marker value from a previous List Jobs
	// response. You only need to include the marker if you are continuing the
	// pagination of results started in a previous List Jobs request.
	Marker *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string

	// The type of job status to return. You can specify the following values:
	// InProgress, Succeeded, or Failed.
	Statuscode *string

	// The maximum number of jobs to be returned. The default limit is 50. The number
	// of jobs returned might be fewer than the specified limit, but the number of
	// returned jobs never exceeds the limit.
	Limit *string
}

// Contains the Amazon S3 Glacier response to your request.
type ListJobsOutput struct {

	// An opaque string used for pagination that specifies the job at which the listing
	// of jobs should begin. You get the marker value from a previous List Jobs
	// response. You only need to include the marker if you are continuing the
	// pagination of the results started in a previous List Jobs request.
	Marker *string

	// A list of job objects. Each job object contains metadata describing the job.
	JobList []*types.GlacierJobDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListJobsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListJobs{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListJobs{}, middleware.After)
}

func newServiceMetadataMiddleware_opListJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "ListJobs",
	}
}
