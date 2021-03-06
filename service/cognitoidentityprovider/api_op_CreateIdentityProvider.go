// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an identity provider for a user pool.
func (c *Client) CreateIdentityProvider(ctx context.Context, params *CreateIdentityProviderInput, optFns ...func(*Options)) (*CreateIdentityProviderOutput, error) {
	stack := middleware.NewStack("CreateIdentityProvider", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateIdentityProviderMiddlewares(stack)
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
	addOpCreateIdentityProviderValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIdentityProvider(options.Region), middleware.Before)
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
			OperationName: "CreateIdentityProvider",
			Err:           err,
		}
	}
	out := result.(*CreateIdentityProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIdentityProviderInput struct {

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// The identity provider type.
	//
	// This member is required.
	ProviderType types.IdentityProviderTypeType

	// A mapping of identity provider attributes to standard and custom user pool
	// attributes.
	AttributeMapping map[string]*string

	// A list of identity provider identifiers.
	IdpIdentifiers []*string

	// The identity provider name.
	//
	// This member is required.
	ProviderName *string

	// The identity provider details. The following list describes the provider detail
	// keys for each identity provider type.
	//
	//     * For Google, Facebook and Login with
	// Amazon:
	//
	//         * client_id
	//
	//         * client_secret
	//
	//         *
	// authorize_scopes
	//
	//     * For Sign in with Apple:
	//
	//         * client_id
	//
	//         *
	// team_id
	//
	//         * key_id
	//
	//         * private_key
	//
	//         * authorize_scopes
	//
	//
	// * For OIDC providers:
	//
	//         * client_id
	//
	//         * client_secret
	//
	//         *
	// attributes_request_method
	//
	//         * oidc_issuer
	//
	//         * authorize_scopes
	//
	//
	// * authorize_url if not available from discovery URL specified by oidc_issuer
	// key
	//
	//         * token_url if not available from discovery URL specified by
	// oidc_issuer key
	//
	//         * attributes_url if not available from discovery URL
	// specified by oidc_issuer key
	//
	//         * jwks_uri if not available from discovery
	// URL specified by oidc_issuer key
	//
	//         * authorize_scopes
	//
	//     * For SAML
	// providers:
	//
	//         * MetadataFile OR MetadataURL
	//
	//         * IDPSignout optional
	//
	// This member is required.
	ProviderDetails map[string]*string
}

type CreateIdentityProviderOutput struct {

	// The newly created identity provider object.
	//
	// This member is required.
	IdentityProvider *types.IdentityProviderType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateIdentityProviderMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateIdentityProvider{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateIdentityProvider{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateIdentityProvider(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "CreateIdentityProvider",
	}
}
