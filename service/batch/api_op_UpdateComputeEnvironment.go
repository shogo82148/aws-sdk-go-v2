// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an AWS Batch compute environment.
func (c *Client) UpdateComputeEnvironment(ctx context.Context, params *UpdateComputeEnvironmentInput, optFns ...func(*Options)) (*UpdateComputeEnvironmentOutput, error) {
	stack := middleware.NewStack("UpdateComputeEnvironment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateComputeEnvironmentMiddlewares(stack)
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
	addOpUpdateComputeEnvironmentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateComputeEnvironment(options.Region), middleware.Before)
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
			OperationName: "UpdateComputeEnvironment",
			Err:           err,
		}
	}
	out := result.(*UpdateComputeEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateComputeEnvironmentInput struct {

	// Details of the compute resources managed by the compute environment. Required
	// for a managed compute environment.
	ComputeResources *types.ComputeResourceUpdate

	// The name or full Amazon Resource Name (ARN) of the compute environment to
	// update.
	//
	// This member is required.
	ComputeEnvironment *string

	// The full Amazon Resource Name (ARN) of the IAM role that allows AWS Batch to
	// make calls to other AWS services on your behalf. If your specified role has a
	// path other than /, then you must either specify the full role ARN (this is
	// recommended) or prefix the role name with the path. Depending on how you created
	// your AWS Batch service role, its ARN may contain the service-role path prefix.
	// When you only specify the name of the service role, AWS Batch assumes that your
	// ARN does not use the service-role path prefix. Because of this, we recommend
	// that you specify the full ARN of your service role when you create compute
	// environments.
	ServiceRole *string

	// The state of the compute environment. Compute environments in the ENABLED state
	// can accept jobs from a queue and scale in or out automatically based on the
	// workload demand of its associated queues.
	State types.CEState
}

type UpdateComputeEnvironmentOutput struct {

	// The name of the compute environment.
	ComputeEnvironmentName *string

	// The Amazon Resource Name (ARN) of the compute environment.
	ComputeEnvironmentArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateComputeEnvironmentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateComputeEnvironment{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateComputeEnvironment{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateComputeEnvironment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "UpdateComputeEnvironment",
	}
}
