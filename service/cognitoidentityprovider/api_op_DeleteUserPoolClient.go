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

// Allows the developer to delete the user pool client.
func (c *Client) DeleteUserPoolClient(ctx context.Context, params *DeleteUserPoolClientInput, optFns ...func(*Options)) (*DeleteUserPoolClientOutput, error) {
	stack := middleware.NewStack("DeleteUserPoolClient", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteUserPoolClientMiddlewares(stack)
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
	addOpDeleteUserPoolClientValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteUserPoolClient(options.Region), middleware.Before)
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
			OperationName: "DeleteUserPoolClient",
			Err:           err,
		}
	}
	out := result.(*DeleteUserPoolClientOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to delete a user pool client.
type DeleteUserPoolClientInput struct {

	// The app client ID of the app associated with the user pool.
	//
	// This member is required.
	ClientId *string

	// The user pool ID for the user pool where you want to delete the client.
	//
	// This member is required.
	UserPoolId *string
}

type DeleteUserPoolClientOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteUserPoolClientMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteUserPoolClient{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteUserPoolClient{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteUserPoolClient(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "DeleteUserPoolClient",
	}
}
