// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes the specified collection. You can use DescribeCollection to get
// information, such as the number of faces indexed into a collection and the
// version of the model used by the collection for face detection.  <p>For more
// information, see Describing a Collection in the Amazon Rekognition Developer
// Guide.</p>
func (c *Client) DescribeCollection(ctx context.Context, params *DescribeCollectionInput, optFns ...func(*Options)) (*DescribeCollectionOutput, error) {
	stack := middleware.NewStack("DescribeCollection", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeCollectionMiddlewares(stack)
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
	addOpDescribeCollectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCollection(options.Region), middleware.Before)
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
			OperationName: "DescribeCollection",
			Err:           err,
		}
	}
	out := result.(*DescribeCollectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCollectionInput struct {

	// The ID of the collection to describe.
	//
	// This member is required.
	CollectionId *string
}

type DescribeCollectionOutput struct {

	// The number of milliseconds since the Unix epoch time until the creation of the
	// collection. The Unix epoch time is 00:00:00 Coordinated Universal Time (UTC),
	// Thursday, 1 January 1970.
	CreationTimestamp *time.Time

	// The Amazon Resource Name (ARN) of the collection.
	CollectionARN *string

	// The version of the face model that's used by the collection for face detection.
	// <p>For more information, see Model Versioning in the Amazon Rekognition
	// Developer Guide.</p>
	FaceModelVersion *string

	// The number of faces that are indexed into the collection. To index faces into a
	// collection, use IndexFaces ().
	FaceCount *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeCollectionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCollection{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCollection{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeCollection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "DescribeCollection",
	}
}
