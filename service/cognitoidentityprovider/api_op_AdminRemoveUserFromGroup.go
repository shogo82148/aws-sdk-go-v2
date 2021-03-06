// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the specified user from the specified group. Calling this action
// requires developer credentials.
func (c *Client) AdminRemoveUserFromGroup(ctx context.Context, params *AdminRemoveUserFromGroupInput, optFns ...func(*Options)) (*AdminRemoveUserFromGroupOutput, error) {
	stack := middleware.NewStack("AdminRemoveUserFromGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAdminRemoveUserFromGroupMiddlewares(stack)
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
	addOpAdminRemoveUserFromGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAdminRemoveUserFromGroup(options.Region), middleware.Before)
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
			OperationName: "AdminRemoveUserFromGroup",
			Err:           err,
		}
	}
	out := result.(*AdminRemoveUserFromGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AdminRemoveUserFromGroupInput struct {

	// The username for the user.
	//
	// This member is required.
	Username *string

	// The user pool ID for the user pool.
	//
	// This member is required.
	UserPoolId *string

	// The group name.
	//
	// This member is required.
	GroupName *string
}

type AdminRemoveUserFromGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAdminRemoveUserFromGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAdminRemoveUserFromGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminRemoveUserFromGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opAdminRemoveUserFromGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminRemoveUserFromGroup",
	}
}
