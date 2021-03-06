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

// Returns the requested sending authorization policies for the given identity (an
// email address or a domain). The policies are returned as a map of policy names
// to policy contents. You can retrieve a maximum of 20 policies at a time. This
// API is for the identity owner only. If you have not verified the identity, this
// API will return an error. Sending authorization is a feature that enables an
// identity owner to authorize other senders to use its identities. For information
// about using sending authorization, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
// You can execute this operation no more than once per second.
func (c *Client) GetIdentityPolicies(ctx context.Context, params *GetIdentityPoliciesInput, optFns ...func(*Options)) (*GetIdentityPoliciesOutput, error) {
	stack := middleware.NewStack("GetIdentityPolicies", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGetIdentityPoliciesMiddlewares(stack)
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
	addOpGetIdentityPoliciesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetIdentityPolicies(options.Region), middleware.Before)
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
			OperationName: "GetIdentityPolicies",
			Err:           err,
		}
	}
	out := result.(*GetIdentityPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to return the requested sending authorization policies for
// an identity. Sending authorization is an Amazon SES feature that enables you to
// authorize other senders to use your identities. For information, see the Amazon
// SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
type GetIdentityPoliciesInput struct {

	// A list of the names of policies to be retrieved. You can retrieve a maximum of
	// 20 policies at a time. If you do not know the names of the policies that are
	// attached to the identity, you can use ListIdentityPolicies.
	//
	// This member is required.
	PolicyNames []*string

	// The identity for which the policies will be retrieved. You can specify an
	// identity by using its name or by using its Amazon Resource Name (ARN). Examples:
	// user@example.com, example.com,
	// arn:aws:ses:us-east-1:123456789012:identity/example.com. To successfully call
	// this API, you must own the identity.
	//
	// This member is required.
	Identity *string
}

// Represents the requested sending authorization policies.
type GetIdentityPoliciesOutput struct {

	// A map of policy names to policies.
	//
	// This member is required.
	Policies map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGetIdentityPoliciesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGetIdentityPolicies{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGetIdentityPolicies{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetIdentityPolicies(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetIdentityPolicies",
	}
}
