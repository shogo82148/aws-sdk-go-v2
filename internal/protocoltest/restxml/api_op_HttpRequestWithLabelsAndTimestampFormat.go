// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// The example tests how requests serialize different timestamp formats in the URI
// path.
func (c *Client) HttpRequestWithLabelsAndTimestampFormat(ctx context.Context, params *HttpRequestWithLabelsAndTimestampFormatInput, optFns ...func(*Options)) (*HttpRequestWithLabelsAndTimestampFormatOutput, error) {
	stack := middleware.NewStack("HttpRequestWithLabelsAndTimestampFormat", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpHttpRequestWithLabelsAndTimestampFormatMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpHttpRequestWithLabelsAndTimestampFormatValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opHttpRequestWithLabelsAndTimestampFormat(options.Region), middleware.Before)
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
			OperationName: "HttpRequestWithLabelsAndTimestampFormat",
			Err:           err,
		}
	}
	out := result.(*HttpRequestWithLabelsAndTimestampFormatOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpRequestWithLabelsAndTimestampFormatInput struct {
	MemberEpochSeconds *time.Time

	MemberHttpDate *time.Time

	MemberDateTime *time.Time

	DefaultFormat *time.Time

	TargetEpochSeconds *time.Time

	TargetHttpDate *time.Time

	TargetDateTime *time.Time
}

type HttpRequestWithLabelsAndTimestampFormatOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpHttpRequestWithLabelsAndTimestampFormatMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpHttpRequestWithLabelsAndTimestampFormat{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpHttpRequestWithLabelsAndTimestampFormat{}, middleware.After)
}

func newServiceMetadataMiddleware_opHttpRequestWithLabelsAndTimestampFormat(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "HttpRequestWithLabelsAndTimestampFormat",
	}
}
