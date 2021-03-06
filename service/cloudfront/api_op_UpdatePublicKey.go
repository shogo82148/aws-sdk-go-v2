// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Update public key information. Note that the only value you can change is the
// comment.
func (c *Client) UpdatePublicKey(ctx context.Context, params *UpdatePublicKeyInput, optFns ...func(*Options)) (*UpdatePublicKeyOutput, error) {
	stack := middleware.NewStack("UpdatePublicKey", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpUpdatePublicKeyMiddlewares(stack)
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
	addOpUpdatePublicKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePublicKey(options.Region), middleware.Before)
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
			OperationName: "UpdatePublicKey",
			Err:           err,
		}
	}
	out := result.(*UpdatePublicKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePublicKeyInput struct {

	// Request to update public key information.
	//
	// This member is required.
	PublicKeyConfig *types.PublicKeyConfig

	// ID of the public key to be updated.
	//
	// This member is required.
	Id *string

	// The value of the ETag header that you received when retrieving the public key to
	// update. For example: E2QWRUHAPOMQZL.
	IfMatch *string
}

type UpdatePublicKeyOutput struct {

	// The current version of the update public key result. For example:
	// E2QWRUHAPOMQZL.
	ETag *string

	// Return the results of updating the public key.
	PublicKey *types.PublicKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpUpdatePublicKeyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpUpdatePublicKey{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpUpdatePublicKey{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdatePublicKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "UpdatePublicKey",
	}
}
