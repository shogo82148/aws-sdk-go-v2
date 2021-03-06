// Code generated by smithy-go-codegen DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Unshares an application from an AWS Organization.This operation can be called
// only from the organization's master account.
func (c *Client) UnshareApplication(ctx context.Context, params *UnshareApplicationInput, optFns ...func(*Options)) (*UnshareApplicationOutput, error) {
	stack := middleware.NewStack("UnshareApplication", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUnshareApplicationMiddlewares(stack)
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
	addOpUnshareApplicationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUnshareApplication(options.Region), middleware.Before)
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
			OperationName: "UnshareApplication",
			Err:           err,
		}
	}
	out := result.(*UnshareApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UnshareApplicationInput struct {

	// The AWS Organization ID to unshare the application from.
	//
	// This member is required.
	OrganizationId *string

	// The Amazon Resource Name (ARN) of the application.
	//
	// This member is required.
	ApplicationId *string
}

type UnshareApplicationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUnshareApplicationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUnshareApplication{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUnshareApplication{}, middleware.After)
}

func newServiceMetadataMiddleware_opUnshareApplication(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "serverlessrepo",
		OperationName: "UnshareApplication",
	}
}
