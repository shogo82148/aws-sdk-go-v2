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

// Updates the description or maximum session duration setting of a role.
func (c *Client) UpdateRole(ctx context.Context, params *UpdateRoleInput, optFns ...func(*Options)) (*UpdateRoleOutput, error) {
	stack := middleware.NewStack("UpdateRole", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpUpdateRoleMiddlewares(stack)
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
	addOpUpdateRoleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRole(options.Region), middleware.Before)
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
			OperationName: "UpdateRole",
			Err:           err,
		}
	}
	out := result.(*UpdateRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRoleInput struct {

	// The new description that you want to apply to the specified role.
	Description *string

	// The maximum session duration (in seconds) that you want to set for the specified
	// role. If you do not specify a value for this setting, the default maximum of one
	// hour is applied. This setting can have a value from 1 hour to 12 hours. Anyone
	// who assumes the role from the AWS CLI or API can use the DurationSeconds API
	// parameter or the duration-seconds CLI parameter to request a longer session. The
	// MaxSessionDuration setting determines the maximum duration that can be requested
	// using the DurationSeconds parameter. If users don't specify a value for the
	// DurationSeconds parameter, their security credentials are valid for one hour by
	// default. This applies when you use the AssumeRole* API operations or the
	// assume-role* CLI operations but does not apply when you use those operations to
	// create a console URL. For more information, see Using IAM Roles
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html) in the IAM
	// User Guide.
	MaxSessionDuration *int32

	// The name of the role that you want to modify.
	//
	// This member is required.
	RoleName *string
}

type UpdateRoleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpUpdateRoleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpUpdateRole{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateRole{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateRole(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "UpdateRole",
	}
}
