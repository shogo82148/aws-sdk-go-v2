// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Remove a field-level encryption configuration.
func (c *Client) DeleteFieldLevelEncryptionConfig(ctx context.Context, params *DeleteFieldLevelEncryptionConfigInput, optFns ...func(*Options)) (*DeleteFieldLevelEncryptionConfigOutput, error) {
	stack := middleware.NewStack("DeleteFieldLevelEncryptionConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpDeleteFieldLevelEncryptionConfigMiddlewares(stack)
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
	addOpDeleteFieldLevelEncryptionConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFieldLevelEncryptionConfig(options.Region), middleware.Before)
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
			OperationName: "DeleteFieldLevelEncryptionConfig",
			Err:           err,
		}
	}
	out := result.(*DeleteFieldLevelEncryptionConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFieldLevelEncryptionConfigInput struct {

	// The ID of the configuration you want to delete from CloudFront.
	//
	// This member is required.
	Id *string

	// The value of the ETag header that you received when retrieving the configuration
	// identity to delete. For example: E2QWRUHAPOMQZL.
	IfMatch *string
}

type DeleteFieldLevelEncryptionConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpDeleteFieldLevelEncryptionConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpDeleteFieldLevelEncryptionConfig{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteFieldLevelEncryptionConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteFieldLevelEncryptionConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "DeleteFieldLevelEncryptionConfig",
	}
}
