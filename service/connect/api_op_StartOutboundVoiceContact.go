// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This API places an outbound call to a contact, and then initiates the contact
// flow. It performs the actions in the contact flow that's specified (in
// ContactFlowId).  <p>Agents are not involved in initiating the outbound API (that
// is, dialing the contact). If  the contact flow places an outbound call to a
// contact, and then puts the contact in queue, that's when the call is routed to
// the agent, like any other inbound case.  <p>There is a 60 second dialing timeout
// for this operation. If the call is not connected after  60 seconds, it fails.
func (c *Client) StartOutboundVoiceContact(ctx context.Context, params *StartOutboundVoiceContactInput, optFns ...func(*Options)) (*StartOutboundVoiceContactOutput, error) {
	stack := middleware.NewStack("StartOutboundVoiceContact", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStartOutboundVoiceContactMiddlewares(stack)
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
	addIdempotencyToken_opStartOutboundVoiceContactMiddleware(stack, options)
	addOpStartOutboundVoiceContactValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartOutboundVoiceContact(options.Region), middleware.Before)
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
			OperationName: "StartOutboundVoiceContact",
			Err:           err,
		}
	}
	out := result.(*StartOutboundVoiceContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartOutboundVoiceContactInput struct {

	// The queue for the call. If you specify a queue, the phone displayed for caller
	// ID is the phone number specified in the queue. If you do not specify a queue,
	// the queue defined in the contact flow is used. If you do not specify a queue,
	// you must specify a source phone number.
	QueueId *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. The token is valid for 7 days after creation. If a contact is
	// already started, the contact ID is returned. If the contact is disconnected, a
	// new contact is started.
	ClientToken *string

	// The phone number of the customer, in E.164 format.
	//
	// This member is required.
	DestinationPhoneNumber *string

	// The phone number associated with the Amazon Connect instance, in E.164 format.
	// If you do not specify a source phone number, you must specify a queue.
	SourcePhoneNumber *string

	// The identifier of the contact flow for the outbound call.
	//
	// This member is required.
	ContactFlowId *string

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// A custom key-value pair using an attribute map. The attributes are standard
	// Amazon Connect attributes, and can be accessed in contact flows just like any
	// other contact attributes. There can be up to 32,768 UTF-8 bytes across all
	// key-value pairs per contact. Attribute keys can include only alphanumeric, dash,
	// and underscore characters.
	Attributes map[string]*string
}

type StartOutboundVoiceContactOutput struct {

	// The identifier of this contact within the Amazon Connect instance.
	ContactId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStartOutboundVoiceContactMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStartOutboundVoiceContact{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStartOutboundVoiceContact{}, middleware.After)
}

type idempotencyToken_initializeOpStartOutboundVoiceContact struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartOutboundVoiceContact) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartOutboundVoiceContact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartOutboundVoiceContactInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartOutboundVoiceContactInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartOutboundVoiceContactMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpStartOutboundVoiceContact{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartOutboundVoiceContact(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "StartOutboundVoiceContact",
	}
}
