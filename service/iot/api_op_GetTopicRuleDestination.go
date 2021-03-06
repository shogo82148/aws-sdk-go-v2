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

// Gets information about a topic rule destination.
func (c *Client) GetTopicRuleDestination(ctx context.Context, params *GetTopicRuleDestinationInput, optFns ...func(*Options)) (*GetTopicRuleDestinationOutput, error) {
	stack := middleware.NewStack("GetTopicRuleDestination", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetTopicRuleDestinationMiddlewares(stack)
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
	addOpGetTopicRuleDestinationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTopicRuleDestination(options.Region), middleware.Before)
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
			OperationName: "GetTopicRuleDestination",
			Err:           err,
		}
	}
	out := result.(*GetTopicRuleDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTopicRuleDestinationInput struct {

	// The ARN of the topic rule destination.
	//
	// This member is required.
	Arn *string
}

type GetTopicRuleDestinationOutput struct {

	// The topic rule destination.
	TopicRuleDestination *types.TopicRuleDestination

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetTopicRuleDestinationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetTopicRuleDestination{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTopicRuleDestination{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetTopicRuleDestination(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "GetTopicRuleDestination",
	}
}
