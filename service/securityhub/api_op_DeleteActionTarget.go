// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a custom action target from Security Hub. Deleting a custom action
// target does not affect any findings or insights that were already sent to Amazon
// CloudWatch Events using the custom action.
func (c *Client) DeleteActionTarget(ctx context.Context, params *DeleteActionTargetInput, optFns ...func(*Options)) (*DeleteActionTargetOutput, error) {
	stack := middleware.NewStack("DeleteActionTarget", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteActionTargetMiddlewares(stack)
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
	addOpDeleteActionTargetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteActionTarget(options.Region), middleware.Before)
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
			OperationName: "DeleteActionTarget",
			Err:           err,
		}
	}
	out := result.(*DeleteActionTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteActionTargetInput struct {

	// The ARN of the custom action target to delete.
	//
	// This member is required.
	ActionTargetArn *string
}

type DeleteActionTargetOutput struct {

	// The ARN of the custom action target that was deleted.
	//
	// This member is required.
	ActionTargetArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteActionTargetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteActionTarget{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteActionTarget{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteActionTarget(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "DeleteActionTarget",
	}
}
