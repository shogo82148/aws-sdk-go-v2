// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an environment.
func (c *Client) UpdateEnvironment(ctx context.Context, params *UpdateEnvironmentInput, optFns ...func(*Options)) (*UpdateEnvironmentOutput, error) {
	stack := middleware.NewStack("UpdateEnvironment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateEnvironmentMiddlewares(stack)
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
	addOpUpdateEnvironmentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEnvironment(options.Region), middleware.Before)
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
			OperationName: "UpdateEnvironment",
			Err:           err,
		}
	}
	out := result.(*UpdateEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEnvironmentInput struct {

	// A description of the environment.
	Description *string

	// Amazon CloudWatch alarms to monitor during the deployment process.
	Monitors []*types.Monitor

	// The environment ID.
	//
	// This member is required.
	EnvironmentId *string

	// The name of the environment.
	Name *string

	// The application ID.
	//
	// This member is required.
	ApplicationId *string
}

type UpdateEnvironmentOutput struct {

	// The name of the environment.
	Name *string

	// The application ID.
	ApplicationId *string

	// The description of the environment.
	Description *string

	// The state of the environment. An environment can be in one of the following
	// states: READY_FOR_DEPLOYMENT, DEPLOYING, ROLLING_BACK, or ROLLED_BACK
	State types.EnvironmentState

	// The environment ID.
	Id *string

	// Amazon CloudWatch alarms monitored during the deployment.
	Monitors []*types.Monitor

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateEnvironmentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateEnvironment{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateEnvironment{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateEnvironment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "UpdateEnvironment",
	}
}
