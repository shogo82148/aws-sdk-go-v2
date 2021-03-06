// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a network profile with the specified details.
func (c *Client) CreateNetworkProfile(ctx context.Context, params *CreateNetworkProfileInput, optFns ...func(*Options)) (*CreateNetworkProfileOutput, error) {
	stack := middleware.NewStack("CreateNetworkProfile", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateNetworkProfileMiddlewares(stack)
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
	addIdempotencyToken_opCreateNetworkProfileMiddleware(stack, options)
	addOpCreateNetworkProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNetworkProfile(options.Region), middleware.Before)
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
			OperationName: "CreateNetworkProfile",
			Err:           err,
		}
	}
	out := result.(*CreateNetworkProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNetworkProfileInput struct {

	// The authentication standard that is used in the EAP framework. Currently,
	// EAP_TLS is supported.
	EapMethod types.NetworkEapMethod

	// Detailed information about a device's network profile.
	Description *string

	// The root certificates of your authentication server that is installed on your
	// devices and used to trust your authentication server during EAP negotiation.
	TrustAnchors []*string

	// The next, or subsequent, password of the Wi-Fi network. This password is
	// asynchronously transmitted to the device and is used when the password of the
	// network changes to NextPassword.
	NextPassword *string

	// The SSID of the Wi-Fi network.
	//
	// This member is required.
	Ssid *string

	// The security type of the Wi-Fi network. This can be WPA2_ENTERPRISE, WPA2_PSK,
	// WPA_PSK, WEP, or OPEN.
	//
	// This member is required.
	SecurityType types.NetworkSecurityType

	// The ARN of the Private Certificate Authority (PCA) created in AWS Certificate
	// Manager (ACM). This is used to issue certificates to the devices.
	CertificateAuthorityArn *string

	// The current password of the Wi-Fi network.
	CurrentPassword *string

	// The name of the network profile associated with a device.
	//
	// This member is required.
	NetworkProfileName *string

	// A unique, user-specified identifier for the request that ensures idempotency.
	//
	// This member is required.
	ClientRequestToken *string
}

type CreateNetworkProfileOutput struct {

	// The ARN of the network profile associated with a device.
	NetworkProfileArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateNetworkProfileMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateNetworkProfile{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateNetworkProfile{}, middleware.After)
}

type idempotencyToken_initializeOpCreateNetworkProfile struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateNetworkProfile) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateNetworkProfile) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateNetworkProfileInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateNetworkProfileInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateNetworkProfileMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateNetworkProfile{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateNetworkProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "CreateNetworkProfile",
	}
}
