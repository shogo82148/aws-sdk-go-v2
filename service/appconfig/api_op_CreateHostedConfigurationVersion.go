// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Create a new configuration in the AppConfig configuration store.
func (c *Client) CreateHostedConfigurationVersion(ctx context.Context, params *CreateHostedConfigurationVersionInput, optFns ...func(*Options)) (*CreateHostedConfigurationVersionOutput, error) {
	stack := middleware.NewStack("CreateHostedConfigurationVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateHostedConfigurationVersionMiddlewares(stack)
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
	addOpCreateHostedConfigurationVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHostedConfigurationVersion(options.Region), middleware.Before)
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
			OperationName: "CreateHostedConfigurationVersion",
			Err:           err,
		}
	}
	out := result.(*CreateHostedConfigurationVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHostedConfigurationVersionInput struct {

	// An optional locking token used to prevent race conditions from overwriting
	// configuration updates when creating a new version. To ensure your data is not
	// overwritten when creating multiple hosted configuration versions in rapid
	// succession, specify the version of the latest hosted configuration version.
	LatestVersionNumber *int32

	// The content of the configuration or the configuration data.
	//
	// This member is required.
	Content []byte

	// A description of the configuration.
	Description *string

	// A standard MIME type describing the format of the configuration content. For
	// more information, see Content-Type
	// (https://docs.aws.amazon.com/https:/www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17).
	//
	// This member is required.
	ContentType *string

	// The configuration profile ID.
	//
	// This member is required.
	ConfigurationProfileId *string

	// The application ID.
	//
	// This member is required.
	ApplicationId *string
}

type CreateHostedConfigurationVersionOutput struct {

	// A standard MIME type describing the format of the configuration content. For
	// more information, see Content-Type
	// (https://docs.aws.amazon.com/https:/www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17).
	ContentType *string

	// A description of the configuration.
	Description *string

	// The configuration profile ID.
	ConfigurationProfileId *string

	// The configuration version.
	VersionNumber *int32

	// The application ID.
	ApplicationId *string

	// The content of the configuration or the configuration data.
	Content []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateHostedConfigurationVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateHostedConfigurationVersion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateHostedConfigurationVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateHostedConfigurationVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "CreateHostedConfigurationVersion",
	}
}
