// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables the specified user in the user pool. Users can't sign in to AppStream
// 2.0 until they are re-enabled. This action does not delete the user.
func (c *Client) DisableUser(ctx context.Context, params *DisableUserInput, optFns ...func(*Options)) (*DisableUserOutput, error) {
	stack := middleware.NewStack("DisableUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDisableUserMiddlewares(stack)
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
	addOpDisableUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableUser(options.Region), middleware.Before)
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
			OperationName: "DisableUser",
			Err:           err,
		}
	}
	out := result.(*DisableUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableUserInput struct {

	// The authentication type for the user. You must specify USERPOOL.
	//
	// This member is required.
	AuthenticationType types.AuthenticationType

	// The email address of the user.  <note> <p>Users' email addresses are
	// case-sensitive.</p> </note>
	//
	// This member is required.
	UserName *string
}

type DisableUserOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDisableUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDisableUser{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisableUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisableUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "DisableUser",
	}
}
