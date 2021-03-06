// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts a task to estimate the quality of the transform.  <p>When you provide
// label sets as examples of truth, AWS Glue machine learning uses some of those
// examples to learn from them. The rest of the labels are used as a test to
// estimate quality.</p> <p>Returns a unique identifier for the run. You can call
// <code>GetMLTaskRun</code> to get more information about the stats of the
// <code>EvaluationTaskRun</code>.</p>
func (c *Client) StartMLEvaluationTaskRun(ctx context.Context, params *StartMLEvaluationTaskRunInput, optFns ...func(*Options)) (*StartMLEvaluationTaskRunOutput, error) {
	stack := middleware.NewStack("StartMLEvaluationTaskRun", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartMLEvaluationTaskRunMiddlewares(stack)
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
	addOpStartMLEvaluationTaskRunValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartMLEvaluationTaskRun(options.Region), middleware.Before)
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
			OperationName: "StartMLEvaluationTaskRun",
			Err:           err,
		}
	}
	out := result.(*StartMLEvaluationTaskRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMLEvaluationTaskRunInput struct {

	// The unique identifier of the machine learning transform.
	//
	// This member is required.
	TransformId *string
}

type StartMLEvaluationTaskRunOutput struct {

	// The unique identifier associated with this run.
	TaskRunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartMLEvaluationTaskRunMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartMLEvaluationTaskRun{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartMLEvaluationTaskRun{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartMLEvaluationTaskRun(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "StartMLEvaluationTaskRun",
	}
}
