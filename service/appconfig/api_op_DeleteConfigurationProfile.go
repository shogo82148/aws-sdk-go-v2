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

// Delete a configuration profile. Deleting a configuration profile does not delete
// a configuration from a host.
func (c *Client) DeleteConfigurationProfile(ctx context.Context, params *DeleteConfigurationProfileInput, optFns ...func(*Options)) (*DeleteConfigurationProfileOutput, error) {
	stack := middleware.NewStack("DeleteConfigurationProfile", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteConfigurationProfileMiddlewares(stack)
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
	addOpDeleteConfigurationProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteConfigurationProfile(options.Region), middleware.Before)
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
			OperationName: "DeleteConfigurationProfile",
			Err:           err,
		}
	}
	out := result.(*DeleteConfigurationProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteConfigurationProfileInput struct {

	// The application ID that includes the configuration profile you want to delete.
	//
	// This member is required.
	ApplicationId *string

	// The ID of the configuration profile you want to delete.
	//
	// This member is required.
	ConfigurationProfileId *string
}

type DeleteConfigurationProfileOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteConfigurationProfileMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteConfigurationProfile{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteConfigurationProfile{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteConfigurationProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "DeleteConfigurationProfile",
	}
}
