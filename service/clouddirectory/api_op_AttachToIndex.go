// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Attaches the specified object to the specified index.
func (c *Client) AttachToIndex(ctx context.Context, params *AttachToIndexInput, optFns ...func(*Options)) (*AttachToIndexOutput, error) {
	stack := middleware.NewStack("AttachToIndex", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpAttachToIndexMiddlewares(stack)
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
	addOpAttachToIndexValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAttachToIndex(options.Region), middleware.Before)
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
			OperationName: "AttachToIndex",
			Err:           err,
		}
	}
	out := result.(*AttachToIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachToIndexInput struct {

	// The Amazon Resource Name (ARN) of the directory where the object and index
	// exist.
	//
	// This member is required.
	DirectoryArn *string

	// A reference to the object that you are attaching to the index.
	//
	// This member is required.
	TargetReference *types.ObjectReference

	// A reference to the index that you are attaching the object to.
	//
	// This member is required.
	IndexReference *types.ObjectReference
}

type AttachToIndexOutput struct {

	// The ObjectIdentifier of the object that was attached to the index.
	AttachedObjectIdentifier *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpAttachToIndexMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpAttachToIndex{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAttachToIndex{}, middleware.After)
}

func newServiceMetadataMiddleware_opAttachToIndex(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "AttachToIndex",
	}
}
