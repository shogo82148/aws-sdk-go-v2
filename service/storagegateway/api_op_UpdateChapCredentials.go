// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the Challenge-Handshake Authentication Protocol (CHAP) credentials for a
// specified iSCSI target. By default, a gateway does not have CHAP enabled;
// however, for added security, you might use it. This operation is supported in
// the volume and tape gateway types.  <important> <p>When you update CHAP
// credentials, all existing connections on the target are closed and initiators
// must reconnect with the new credentials.</p> </important>
func (c *Client) UpdateChapCredentials(ctx context.Context, params *UpdateChapCredentialsInput, optFns ...func(*Options)) (*UpdateChapCredentialsOutput, error) {
	stack := middleware.NewStack("UpdateChapCredentials", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateChapCredentialsMiddlewares(stack)
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
	addOpUpdateChapCredentialsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateChapCredentials(options.Region), middleware.Before)
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
			OperationName: "UpdateChapCredentials",
			Err:           err,
		}
	}
	out := result.(*UpdateChapCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing one or more of the following fields:  <ul> <li> <p>
// <a>UpdateChapCredentialsInput$InitiatorName</a> </p> </li> <li> <p>
// <a>UpdateChapCredentialsInput$SecretToAuthenticateInitiator</a> </p> </li> <li>
// <p> <a>UpdateChapCredentialsInput$SecretToAuthenticateTarget</a> </p> </li> <li>
// <p> <a>UpdateChapCredentialsInput$TargetARN</a> </p> </li> </ul>
type UpdateChapCredentialsInput struct {

	// The secret key that the initiator (for example, the Windows client) must provide
	// to participate in mutual CHAP with the target.  <note> <p>The secret key must be
	// between 12 and 16 bytes when encoded in UTF-8.</p> </note>
	//
	// This member is required.
	SecretToAuthenticateInitiator *string

	// The secret key that the target must provide to participate in mutual CHAP with
	// the initiator (e.g. Windows client).  <p>Byte constraints: Minimum bytes of 12.
	// Maximum bytes of 16.</p> <note> <p>The secret key must be between 12 and 16
	// bytes when encoded in UTF-8.</p> </note>
	SecretToAuthenticateTarget *string

	// The iSCSI initiator that connects to the target.
	//
	// This member is required.
	InitiatorName *string

	// The Amazon Resource Name (ARN) of the iSCSI volume target. Use the
	// DescribeStorediSCSIVolumes () operation to return the TargetARN for specified
	// VolumeARN.
	//
	// This member is required.
	TargetARN *string
}

// A JSON object containing the following fields:
type UpdateChapCredentialsOutput struct {

	// The Amazon Resource Name (ARN) of the target. This is the same target specified
	// in the request.
	TargetARN *string

	// The iSCSI initiator that connects to the target. This is the same initiator name
	// specified in the request.
	InitiatorName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateChapCredentialsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateChapCredentials{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateChapCredentials{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateChapCredentials(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "UpdateChapCredentials",
	}
}
