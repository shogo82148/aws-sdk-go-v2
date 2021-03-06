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

// Updates information about the billing group.
func (c *Client) UpdateBillingGroup(ctx context.Context, params *UpdateBillingGroupInput, optFns ...func(*Options)) (*UpdateBillingGroupOutput, error) {
	stack := middleware.NewStack("UpdateBillingGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateBillingGroupMiddlewares(stack)
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
	addOpUpdateBillingGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBillingGroup(options.Region), middleware.Before)
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
			OperationName: "UpdateBillingGroup",
			Err:           err,
		}
	}
	out := result.(*UpdateBillingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateBillingGroupInput struct {

	// The properties of the billing group.
	//
	// This member is required.
	BillingGroupProperties *types.BillingGroupProperties

	// The expected version of the billing group. If the version of the billing group
	// does not match the expected version specified in the request, the
	// UpdateBillingGroup request is rejected with a VersionConflictException.
	ExpectedVersion *int64

	// The name of the billing group.
	//
	// This member is required.
	BillingGroupName *string
}

type UpdateBillingGroupOutput struct {

	// The latest version of the billing group.
	Version *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateBillingGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateBillingGroup{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateBillingGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateBillingGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "UpdateBillingGroup",
	}
}
