// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a file gateway to an Active Directory domain. This operation is only
// supported for file gateways that support the SMB file protocol.
func (c *Client) JoinDomain(ctx context.Context, params *JoinDomainInput, optFns ...func(*Options)) (*JoinDomainOutput, error) {
	stack := middleware.NewStack("JoinDomain", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpJoinDomainMiddlewares(stack)
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
	addOpJoinDomainValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opJoinDomain(options.Region), middleware.Before)
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
			OperationName: "JoinDomain",
			Err:           err,
		}
	}
	out := result.(*JoinDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// JoinDomainInput
type JoinDomainInput struct {

	// Sets the user name of user who has permission to add the gateway to the Active
	// Directory domain. The domain user account should be enabled to join computers to
	// the domain. For example, you can use the domain administrator account or an
	// account with delegated permissions to join computers to the domain.
	//
	// This member is required.
	UserName *string

	// The name of the domain that you want the gateway to join.
	//
	// This member is required.
	DomainName *string

	// List of IPv4 addresses, NetBIOS names, or host names of your domain server. If
	// you need to specify the port number include it after the colon (“:”). For
	// example, mydc.mydomain.com:389.
	DomainControllers []*string

	// Sets the password of the user who has permission to add the gateway to the
	// Active Directory domain.
	//
	// This member is required.
	Password *string

	// The organizational unit (OU) is a container in an Active Directory that can hold
	// users, groups, computers, and other OUs and this parameter specifies the OU that
	// the gateway will join within the AD domain.
	OrganizationalUnit *string

	// Specifies the time in seconds, in which the JoinDomain operation must complete.
	// The default is 20 seconds.
	TimeoutInSeconds *int32

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string
}

// JoinDomainOutput
type JoinDomainOutput struct {

	// The unique Amazon Resource Name (ARN) of the gateway that joined the domain.
	GatewayARN *string

	// Indicates the status of the gateway as a member of the Active Directory domain.
	// <ul> <li> <p> <code>ACCESS_DENIED</code>: Indicates that the
	// <code>JoinDomain</code> operation failed due to an authentication error.</p>
	// </li> <li> <p> <code>DETACHED</code>: Indicates that gateway is not joined to a
	// domain.</p> </li> <li> <p> <code>JOINED</code>: Indicates that the gateway has
	// successfully joined a domain.</p> </li> <li> <p> <code>JOINING</code>: Indicates
	// that a <code>JoinDomain</code> operation is in progress.</p> </li> <li> <p>
	// <code>NETWORK_ERROR</code>: Indicates that <code>JoinDomain</code> operation
	// failed due to a network or connectivity error.</p> </li> <li> <p>
	// <code>TIMEOUT</code>: Indicates that the <code>JoinDomain</code> operation
	// failed because the operation didn't complete within the allotted time.</p> </li>
	// <li> <p> <code>UNKNOWN_ERROR</code>: Indicates that the <code>JoinDomain</code>
	// operation failed due to another type of error.</p> </li> </ul>
	ActiveDirectoryStatus types.ActiveDirectoryStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpJoinDomainMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpJoinDomain{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpJoinDomain{}, middleware.After)
}

func newServiceMetadataMiddleware_opJoinDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "JoinDomain",
	}
}
