// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts an asynchronous job to detect medical conditions and link them to the
// ICD-10-CM ontology. Use the DescribeICD10CMInferenceJob operation to track the
// status of a job.
func (c *Client) StartICD10CMInferenceJob(ctx context.Context, params *StartICD10CMInferenceJobInput, optFns ...func(*Options)) (*StartICD10CMInferenceJobOutput, error) {
	stack := middleware.NewStack("StartICD10CMInferenceJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartICD10CMInferenceJobMiddlewares(stack)
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
	addIdempotencyToken_opStartICD10CMInferenceJobMiddleware(stack, options)
	addOpStartICD10CMInferenceJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartICD10CMInferenceJob(options.Region), middleware.Before)
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
			OperationName: "StartICD10CMInferenceJob",
			Err:           err,
		}
	}
	out := result.(*StartICD10CMInferenceJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartICD10CMInferenceJobInput struct {

	// Specifies where to send the output files.
	//
	// This member is required.
	OutputDataConfig *types.OutputDataConfig

	// The identifier of the job.
	JobName *string

	// The language of the input documents. All documents must be in the same language.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// A unique identifier for the request. If you don't set the client request token,
	// Amazon Comprehend Medical generates one.
	ClientRequestToken *string

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM)
	// role that grants Amazon Comprehend Medical read access to your input data. For
	// more information, see  Role-Based Permissions Required for Asynchronous
	// Operations
	// (https://docs.aws.amazon.com/comprehend/latest/dg/access-control-managing-permissions-med.html#auth-role-permissions-med).
	//
	// This member is required.
	DataAccessRoleArn *string

	// An AWS Key Management Service key to encrypt your output files. If you do not
	// specify a key, the files are written in plain text.
	KMSKey *string

	// Specifies the format and location of the input data for the job.
	//
	// This member is required.
	InputDataConfig *types.InputDataConfig
}

type StartICD10CMInferenceJobOutput struct {

	// The identifier generated for the job. To get the status of a job, use this
	// identifier with the StartICD10CMInferenceJob operation.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartICD10CMInferenceJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartICD10CMInferenceJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartICD10CMInferenceJob{}, middleware.After)
}

type idempotencyToken_initializeOpStartICD10CMInferenceJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartICD10CMInferenceJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartICD10CMInferenceJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartICD10CMInferenceJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartICD10CMInferenceJobInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartICD10CMInferenceJobMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpStartICD10CMInferenceJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartICD10CMInferenceJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehendmedical",
		OperationName: "StartICD10CMInferenceJob",
	}
}
