// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Create a new voice message and send it to a recipient's phone number.
func (c *Client) SendVoiceMessage(ctx context.Context, params *SendVoiceMessageInput, optFns ...func(*Options)) (*SendVoiceMessageOutput, error) {
	stack := middleware.NewStack("SendVoiceMessage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpSendVoiceMessageMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opSendVoiceMessage(options.Region), middleware.Before)
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
			OperationName: "SendVoiceMessage",
			Err:           err,
		}
	}
	out := result.(*SendVoiceMessageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// SendVoiceMessageRequest
type SendVoiceMessageInput struct {

	// The phone number that appears on recipients' devices when they receive the
	// message.
	CallerId *string

	// An object that contains a voice message and information about the recipient that
	// you want to send it to.
	Content *types.VoiceMessageContent

	// The name of the configuration set that you want to use to send the message.
	ConfigurationSetName *string

	// The phone number that Amazon Pinpoint should use to send the voice message. This
	// isn't necessarily the phone number that appears on recipients' devices when they
	// receive the message, because you can specify a CallerId parameter in the
	// request.
	OriginationPhoneNumber *string

	// The phone number that you want to send the voice message to.
	DestinationPhoneNumber *string
}

// An object that that contains the Message ID of a Voice message that was sent
// successfully.
type SendVoiceMessageOutput struct {

	// A unique identifier for the voice message.
	MessageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpSendVoiceMessageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpSendVoiceMessage{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpSendVoiceMessage{}, middleware.After)
}

func newServiceMetadataMiddleware_opSendVoiceMessage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "SendVoiceMessage",
	}
}
