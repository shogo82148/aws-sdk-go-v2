// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the user belonging to a file transfer protocol-enabled server you
// specify.  <p>No response returns from this operation.</p> <note> <p>When you
// delete a user from a server, the user's information is lost.</p> </note>
func (c *Client) DeleteUser(ctx context.Context, params *DeleteUserInput, optFns ...func(*Options)) (*DeleteUserOutput, error) {
	stack := middleware.NewStack("DeleteUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteUserMiddlewares(stack)
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
	addOpDeleteUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteUser(options.Region), middleware.Before)
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
			OperationName: "DeleteUser",
			Err:           err,
		}
	}
	out := result.(*DeleteUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteUserInput struct {

	// A system-assigned unique identifier for a file transfer protocol-enabled server
	// instance that has the user assigned to it.
	//
	// This member is required.
	ServerId *string

	// A unique string that identifies a user that is being deleted from a file
	// transfer protocol-enabled server.
	//
	// This member is required.
	UserName *string
}

type DeleteUserOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteUser{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "DeleteUser",
	}
}
