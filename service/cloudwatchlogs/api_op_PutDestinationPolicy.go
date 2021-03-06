// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates or updates an access policy associated with an existing destination. An
// access policy is an IAM policy document
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies_overview.html) that
// is used to authorize claims to register a subscription filter against a given
// destination.
func (c *Client) PutDestinationPolicy(ctx context.Context, params *PutDestinationPolicyInput, optFns ...func(*Options)) (*PutDestinationPolicyOutput, error) {
	stack := middleware.NewStack("PutDestinationPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutDestinationPolicyMiddlewares(stack)
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
	addOpPutDestinationPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutDestinationPolicy(options.Region), middleware.Before)
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
			OperationName: "PutDestinationPolicy",
			Err:           err,
		}
	}
	out := result.(*PutDestinationPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutDestinationPolicyInput struct {

	// A name for an existing destination.
	//
	// This member is required.
	DestinationName *string

	// An IAM policy document that authorizes cross-account users to deliver their log
	// events to the associated destination.
	//
	// This member is required.
	AccessPolicy *string
}

type PutDestinationPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutDestinationPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutDestinationPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutDestinationPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutDestinationPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "PutDestinationPolicy",
	}
}
