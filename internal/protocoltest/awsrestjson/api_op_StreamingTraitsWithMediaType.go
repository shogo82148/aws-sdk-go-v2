// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
)

// This examples serializes a streaming media-typed blob shape in the request body.
// This examples uses a @mediaType trait on the payload to force a custom
// content-type to be serialized.
func (c *Client) StreamingTraitsWithMediaType(ctx context.Context, params *StreamingTraitsWithMediaTypeInput, optFns ...func(*Options)) (*StreamingTraitsWithMediaTypeOutput, error) {
	stack := middleware.NewStack("StreamingTraitsWithMediaType", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStreamingTraitsWithMediaTypeMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStreamingTraitsWithMediaType(options.Region), middleware.Before)
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
			OperationName: "StreamingTraitsWithMediaType",
			Err:           err,
		}
	}
	out := result.(*StreamingTraitsWithMediaTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StreamingTraitsWithMediaTypeInput struct {
	Foo *string

	Blob io.Reader
}

type StreamingTraitsWithMediaTypeOutput struct {
	Foo *string

	Blob io.ReadCloser

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStreamingTraitsWithMediaTypeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStreamingTraitsWithMediaType{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStreamingTraitsWithMediaType{}, middleware.After)
}

func newServiceMetadataMiddleware_opStreamingTraitsWithMediaType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StreamingTraitsWithMediaType",
	}
}
