// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a virtual MFA device. You must deactivate a user's virtual MFA device
// before you can delete it. For information about deactivating MFA devices, see
// DeactivateMFADevice ().
func (c *Client) DeleteVirtualMFADevice(ctx context.Context, params *DeleteVirtualMFADeviceInput, optFns ...func(*Options)) (*DeleteVirtualMFADeviceOutput, error) {
	stack := middleware.NewStack("DeleteVirtualMFADevice", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteVirtualMFADeviceMiddlewares(stack)
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
	addOpDeleteVirtualMFADeviceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVirtualMFADevice(options.Region), middleware.Before)
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
			OperationName: "DeleteVirtualMFADevice",
			Err:           err,
		}
	}
	out := result.(*DeleteVirtualMFADeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteVirtualMFADeviceInput struct {

	// The serial number that uniquely identifies the MFA device. For virtual MFA
	// devices, the serial number is the same as the ARN. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: =,.@:/-
	//
	// This member is required.
	SerialNumber *string
}

type DeleteVirtualMFADeviceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteVirtualMFADeviceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteVirtualMFADevice{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteVirtualMFADevice{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteVirtualMFADevice(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "DeleteVirtualMFADevice",
	}
}
