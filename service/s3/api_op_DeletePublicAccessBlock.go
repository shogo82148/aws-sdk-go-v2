// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the PublicAccessBlock configuration for an Amazon S3 bucket. To use this
// operation, you must have the s3:PutBucketPublicAccessBlock permission. For more
// information about permissions, see Permissions Related to Bucket Subresource
// Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html).
// <p>The following operations are related to
// <code>DeletePublicAccessBlock</code>:</p> <ul> <li> <p> <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html">Using
// Amazon S3 Block Public Access</a> </p> </li> <li> <p>
// <a>GetPublicAccessBlock</a> </p> </li> <li> <p> <a>PutPublicAccessBlock</a> </p>
// </li> <li> <p> <a>GetBucketPolicyStatus</a> </p> </li> </ul>
func (c *Client) DeletePublicAccessBlock(ctx context.Context, params *DeletePublicAccessBlockInput, optFns ...func(*Options)) (*DeletePublicAccessBlockOutput, error) {
	stack := middleware.NewStack("DeletePublicAccessBlock", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpDeletePublicAccessBlockMiddlewares(stack)
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
	addOpDeletePublicAccessBlockValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePublicAccessBlock(options.Region), middleware.Before)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	addMetadataRetrieverMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)

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
			OperationName: "DeletePublicAccessBlock",
			Err:           err,
		}
	}
	out := result.(*DeletePublicAccessBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePublicAccessBlockInput struct {

	// The Amazon S3 bucket whose PublicAccessBlock configuration you want to delete.
	//
	// This member is required.
	Bucket *string
}

type DeletePublicAccessBlockOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpDeletePublicAccessBlockMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpDeletePublicAccessBlock{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpDeletePublicAccessBlock{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeletePublicAccessBlock(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeletePublicAccessBlock",
	}
}
