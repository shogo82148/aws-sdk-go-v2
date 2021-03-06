// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sends email to a maximum of 50 users, inviting them to the specified Amazon
// Chime Team account. Only Team account types are currently supported for this
// action.
func (c *Client) InviteUsers(ctx context.Context, params *InviteUsersInput, optFns ...func(*Options)) (*InviteUsersOutput, error) {
	stack := middleware.NewStack("InviteUsers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpInviteUsersMiddlewares(stack)
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
	addOpInviteUsersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opInviteUsers(options.Region), middleware.Before)
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
			OperationName: "InviteUsers",
			Err:           err,
		}
	}
	out := result.(*InviteUsersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InviteUsersInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The user type.
	UserType types.UserType

	// The user email addresses to which to send the email invitation.
	//
	// This member is required.
	UserEmailList []*string
}

type InviteUsersOutput struct {

	// The email invitation details.
	Invites []*types.Invite

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpInviteUsersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpInviteUsers{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpInviteUsers{}, middleware.After)
}

func newServiceMetadataMiddleware_opInviteUsers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "InviteUsers",
	}
}
