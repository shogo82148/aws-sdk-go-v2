// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified IP address filter. For information about managing IP
// address filters, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-managing-ip-filters.html).
// You can execute this operation no more than once per second.
func (c *Client) DeleteReceiptFilter(ctx context.Context, params *DeleteReceiptFilterInput, optFns ...func(*Options)) (*DeleteReceiptFilterOutput, error) {
	stack := middleware.NewStack("DeleteReceiptFilter", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteReceiptFilterMiddlewares(stack)
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
	addOpDeleteReceiptFilterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteReceiptFilter(options.Region), middleware.Before)
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
			OperationName: "DeleteReceiptFilter",
			Err:           err,
		}
	}
	out := result.(*DeleteReceiptFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to delete an IP address filter. You use IP address filters
// when you receive email with Amazon SES. For more information, see the Amazon SES
// Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
type DeleteReceiptFilterInput struct {

	// The name of the IP address filter to delete.
	//
	// This member is required.
	FilterName *string
}

// An empty element returned on a successful request.
type DeleteReceiptFilterOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteReceiptFilterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteReceiptFilter{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteReceiptFilter{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteReceiptFilter(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "DeleteReceiptFilter",
	}
}
