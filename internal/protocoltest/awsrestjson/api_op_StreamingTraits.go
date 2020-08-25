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

// This examples serializes a streaming blob shape in the request body. In this
// example, no JSON document is synthesized because the payload is not a structure
// or a union type.
func (c *Client) StreamingTraits(ctx context.Context, params *StreamingTraitsInput, optFns ...func(*Options)) (*StreamingTraitsOutput, error) {
	stack := middleware.NewStack("StreamingTraits", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStreamingTraitsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStreamingTraits(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "StreamingTraits",
			Err:           err,
		}
	}
	out := result.(*StreamingTraitsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StreamingTraitsInput struct {
	Foo  *string
	Blob io.Reader
}

type StreamingTraitsOutput struct {
	Foo  *string
	Blob io.ReadCloser

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStreamingTraitsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStreamingTraits{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStreamingTraits{}, middleware.After)
}

func newServiceMetadataMiddleware_opStreamingTraits(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Rest Json Protocol",
		ServiceID:      "restjsonprotocol",
		EndpointPrefix: "restjsonprotocol",
		OperationName:  "StreamingTraits",
	}
}