// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the permissions for a specified stack. Required Permissions: To use
// this action, an IAM user must have a Manage permissions level for the stack, or
// an attached policy that explicitly grants permissions. For more information on
// user permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) DescribePermissions(ctx context.Context, params *DescribePermissionsInput, optFns ...func(*Options)) (*DescribePermissionsOutput, error) {
	stack := middleware.NewStack("DescribePermissions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribePermissionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePermissions(options.Region), middleware.Before)
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
			OperationName: "DescribePermissions",
			Err:           err,
		}
	}
	out := result.(*DescribePermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePermissionsInput struct {

	// The user's IAM ARN. This can also be a federated user's ARN. For more
	// information about IAM ARNs, see Using Identifiers
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html).
	IamUserArn *string

	// The stack ID.
	StackId *string
}

// Contains the response to a DescribePermissions request.
type DescribePermissionsOutput struct {

	// An array of Permission objects that describe the stack permissions.
	//
	//     * If
	// the request object contains only a stack ID, the array contains a Permission
	// object with permissions for each of the stack IAM ARNs.
	//
	//     * If the request
	// object contains only an IAM ARN, the array contains a Permission object with
	// permissions for each of the user's stack IDs.
	//
	//     * If the request contains a
	// stack ID and an IAM ARN, the array contains a single Permission object with
	// permissions for the specified stack and IAM ARN.
	Permissions []*types.Permission

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribePermissionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePermissions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePermissions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribePermissions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "DescribePermissions",
	}
}
