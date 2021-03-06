// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a specified user profile. Required Permissions: To use this action, an
// IAM user must have an attached policy that explicitly grants permissions. For
// more information about user permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) UpdateUserProfile(ctx context.Context, params *UpdateUserProfileInput, optFns ...func(*Options)) (*UpdateUserProfileOutput, error) {
	stack := middleware.NewStack("UpdateUserProfile", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateUserProfileMiddlewares(stack)
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
	addOpUpdateUserProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUserProfile(options.Region), middleware.Before)
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
			OperationName: "UpdateUserProfile",
			Err:           err,
		}
	}
	out := result.(*UpdateUserProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateUserProfileInput struct {

	// The user's SSH user name. The allowable characters are [a-z], [A-Z], [0-9], '-',
	// and '_'. If the specified name includes other punctuation marks, AWS OpsWorks
	// Stacks removes them. For example, my.name will be changed to myname. If you do
	// not specify an SSH user name, AWS OpsWorks Stacks generates one from the IAM
	// user name.
	SshUsername *string

	// Whether users can specify their own SSH public key through the My Settings page.
	// For more information, see Managing User Permissions
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/security-settingsshkey.html).
	AllowSelfManagement *bool

	// The user IAM ARN. This can also be a federated user's ARN.
	//
	// This member is required.
	IamUserArn *string

	// The user's new SSH public key.
	SshPublicKey *string
}

type UpdateUserProfileOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateUserProfileMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateUserProfile{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateUserProfile{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateUserProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "UpdateUserProfile",
	}
}
