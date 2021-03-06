// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Given an identity (an email address or a domain), enables or disables whether
// Amazon SES forwards bounce and complaint notifications as email. Feedback
// forwarding can only be disabled when Amazon Simple Notification Service (Amazon
// SNS) topics are specified for both bounces and complaints. Feedback forwarding
// does not apply to delivery notifications. Delivery notifications are only
// available through Amazon SNS. You can execute this operation no more than once
// per second. For more information about using notifications with Amazon SES, see
// the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/notifications.html).
func (c *Client) SetIdentityFeedbackForwardingEnabled(ctx context.Context, params *SetIdentityFeedbackForwardingEnabledInput, optFns ...func(*Options)) (*SetIdentityFeedbackForwardingEnabledOutput, error) {
	stack := middleware.NewStack("SetIdentityFeedbackForwardingEnabled", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpSetIdentityFeedbackForwardingEnabledMiddlewares(stack)
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
	addOpSetIdentityFeedbackForwardingEnabledValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetIdentityFeedbackForwardingEnabled(options.Region), middleware.Before)
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
			OperationName: "SetIdentityFeedbackForwardingEnabled",
			Err:           err,
		}
	}
	out := result.(*SetIdentityFeedbackForwardingEnabledOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to enable or disable whether Amazon SES forwards you bounce
// and complaint notifications through email. For information about email feedback
// forwarding, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/notifications-via-email.html).
type SetIdentityFeedbackForwardingEnabledInput struct {

	// The identity for which to set bounce and complaint notification forwarding.
	// Examples: user@example.com, example.com.
	//
	// This member is required.
	Identity *string

	// Sets whether Amazon SES will forward bounce and complaint notifications as
	// email. true specifies that Amazon SES will forward bounce and complaint
	// notifications as email, in addition to any Amazon SNS topic publishing otherwise
	// specified. false specifies that Amazon SES will publish bounce and complaint
	// notifications only through Amazon SNS. This value can only be set to false when
	// Amazon SNS topics are set for both Bounce and Complaint notification types.
	//
	// This member is required.
	ForwardingEnabled *bool
}

// An empty element returned on a successful request.
type SetIdentityFeedbackForwardingEnabledOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpSetIdentityFeedbackForwardingEnabledMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpSetIdentityFeedbackForwardingEnabled{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpSetIdentityFeedbackForwardingEnabled{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetIdentityFeedbackForwardingEnabled(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "SetIdentityFeedbackForwardingEnabled",
	}
}
