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
	"time"
)

// Gets information about a mitigation action.
func (c *Client) DescribeMitigationAction(ctx context.Context, params *DescribeMitigationActionInput, optFns ...func(*Options)) (*DescribeMitigationActionOutput, error) {
	stack := middleware.NewStack("DescribeMitigationAction", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeMitigationActionMiddlewares(stack)
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
	addOpDescribeMitigationActionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMitigationAction(options.Region), middleware.Before)
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
			OperationName: "DescribeMitigationAction",
			Err:           err,
		}
	}
	out := result.(*DescribeMitigationActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMitigationActionInput struct {

	// The friendly name that uniquely identifies the mitigation action.
	//
	// This member is required.
	ActionName *string
}

type DescribeMitigationActionOutput struct {

	// The ARN of the IAM role used to apply this action.
	RoleArn *string

	// The ARN that identifies this migration action.
	ActionArn *string

	// The date and time when the mitigation action was last changed.
	LastModifiedDate *time.Time

	// Parameters that control how the mitigation action is applied, specific to the
	// type of mitigation action.
	ActionParams *types.MitigationActionParams

	// The friendly name that uniquely identifies the mitigation action.
	ActionName *string

	// The type of mitigation action.
	ActionType types.MitigationActionType

	// The date and time when the mitigation action was added to your AWS account.
	CreationDate *time.Time

	// A unique identifier for this action.
	ActionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeMitigationActionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeMitigationAction{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeMitigationAction{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeMitigationAction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeMitigationAction",
	}
}
