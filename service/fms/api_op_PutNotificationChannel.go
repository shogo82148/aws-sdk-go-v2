// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Designates the IAM role and Amazon Simple Notification Service (SNS) topic that
// AWS Firewall Manager uses to record SNS logs.
func (c *Client) PutNotificationChannel(ctx context.Context, params *PutNotificationChannelInput, optFns ...func(*Options)) (*PutNotificationChannelOutput, error) {
	stack := middleware.NewStack("PutNotificationChannel", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutNotificationChannelMiddlewares(stack)
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
	addOpPutNotificationChannelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutNotificationChannel(options.Region), middleware.Before)
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
			OperationName: "PutNotificationChannel",
			Err:           err,
		}
	}
	out := result.(*PutNotificationChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutNotificationChannelInput struct {

	// The Amazon Resource Name (ARN) of the SNS topic that collects notifications from
	// AWS Firewall Manager.
	//
	// This member is required.
	SnsTopicArn *string

	// The Amazon Resource Name (ARN) of the IAM role that allows Amazon SNS to record
	// AWS Firewall Manager activity.
	//
	// This member is required.
	SnsRoleName *string
}

type PutNotificationChannelOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutNotificationChannelMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutNotificationChannel{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutNotificationChannel{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutNotificationChannel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "PutNotificationChannel",
	}
}
