// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// If the IdentityProviderType of a file transfer protocol-enabled server is
// API_Gateway, tests whether your API Gateway is set up successfully. We highly
// recommend that you call this operation to test your authentication method as
// soon as you create your server. By doing so, you can troubleshoot issues with
// the API Gateway integration to ensure that your users can successfully use the
// service.
func (c *Client) TestIdentityProvider(ctx context.Context, params *TestIdentityProviderInput, optFns ...func(*Options)) (*TestIdentityProviderOutput, error) {
	stack := middleware.NewStack("TestIdentityProvider", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpTestIdentityProviderMiddlewares(stack)
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
	addOpTestIdentityProviderValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTestIdentityProvider(options.Region), middleware.Before)
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
			OperationName: "TestIdentityProvider",
			Err:           err,
		}
	}
	out := result.(*TestIdentityProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TestIdentityProviderInput struct {

	// The password of the user account to be tested.
	UserPassword *string

	// The source IP address of the user account to be tested.
	SourceIp *string

	// The type of file transfer protocol to be tested.  <p>The available protocols
	// are:</p> <ul> <li> <p>Secure Shell (SSH) File Transfer Protocol (SFTP)</p> </li>
	// <li> <p>File Transfer Protocol Secure (FTPS)</p> </li> <li> <p>File Transfer
	// Protocol (FTP)</p> </li> </ul>
	ServerProtocol types.Protocol

	// A system-assigned identifier for a specific file transfer protocol-enabled
	// server. That server's user authentication method is tested with a user name and
	// password.
	//
	// This member is required.
	ServerId *string

	// The name of the user account to be tested.
	//
	// This member is required.
	UserName *string
}

type TestIdentityProviderOutput struct {

	// The response that is returned from your API Gateway.
	Response *string

	// A message that indicates whether the test was successful or not.
	Message *string

	// The HTTP status code that is the response from your API Gateway.
	//
	// This member is required.
	StatusCode *int32

	// The endpoint of the service used to authenticate a user.
	//
	// This member is required.
	Url *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpTestIdentityProviderMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpTestIdentityProvider{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpTestIdentityProvider{}, middleware.After)
}

func newServiceMetadataMiddleware_opTestIdentityProvider(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "TestIdentityProvider",
	}
}
