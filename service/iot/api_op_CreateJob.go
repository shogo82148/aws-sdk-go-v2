// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a job.
func (c *Client) CreateJob(ctx context.Context, params *CreateJobInput, optFns ...func(*Options)) (*CreateJobOutput, error) {
	stack := middleware.NewStack("CreateJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateJobMiddlewares(stack)
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
	addOpCreateJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateJob(options.Region), middleware.Before)
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
			OperationName: "CreateJob",
			Err:           err,
		}
	}
	out := result.(*CreateJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateJobInput struct {

	// The job document. If the job document resides in an S3 bucket, you must use a
	// placeholder link when specifying the document. The placeholder link is of the
	// following form: ${aws:iot:s3-presigned-url:https://s3.amazonaws.com/bucket/key}
	// where bucket is your bucket name and key is the object in the bucket to which
	// you are linking.
	Document *string

	// A job identifier which must be unique for your AWS account. We recommend using a
	// UUID. Alpha-numeric characters, "-" and "_" are valid for use here.
	//
	// This member is required.
	JobId *string

	// Specifies the amount of time each device has to finish its execution of the job.
	// The timer is started when the job execution status is set to IN_PROGRESS. If the
	// job execution status is not set to another terminal state before the time
	// expires, it will be automatically set to TIMED_OUT.
	TimeoutConfig *types.TimeoutConfig

	// A list of things and thing groups to which the job should be sent.
	//
	// This member is required.
	Targets []*string

	// Metadata which can be used to manage the job.
	Tags []*types.Tag

	// An S3 link to the job document.
	DocumentSource *string

	// Allows you to create criteria to abort a job.
	AbortConfig *types.AbortConfig

	// A short text description of the job.
	Description *string

	// Specifies whether the job will continue to run (CONTINUOUS), or will be complete
	// after all those things specified as targets have completed the job (SNAPSHOT).
	// If continuous, the job may also be run on a thing when a change is detected in a
	// target. For example, a job will run on a thing when the thing is added to a
	// target group, even after the job was completed by all things originally in the
	// group.
	TargetSelection types.TargetSelection

	// Allows you to create a staged rollout of the job.
	JobExecutionsRolloutConfig *types.JobExecutionsRolloutConfig

	// Configuration information for pre-signed S3 URLs.
	PresignedUrlConfig *types.PresignedUrlConfig
}

type CreateJobOutput struct {

	// The unique identifier you assigned to this job.
	JobId *string

	// The job description.
	Description *string

	// The job ARN.
	JobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateJob{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateJob",
	}
}
