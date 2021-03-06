// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation aborts a multipart upload. After a multipart upload is aborted,
// no additional parts can be uploaded using that upload ID. The storage consumed
// by any previously uploaded parts will be freed. However, if any part uploads are
// currently in progress, those part uploads might or might not succeed. As a
// result, it might be necessary to abort a given multipart upload multiple times
// in order to completely free all storage consumed by all parts. To verify that
// all parts have been removed, so you don't get charged for the part storage, you
// should call the ListParts () operation and ensure that the parts list is empty.
// For information about permissions required to use the multipart upload API, see
// Multipart Upload API and Permissions
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html). The
// following operations are related to AbortMultipartUpload:
//
//     *
// CreateMultipartUpload ()
//
//     * UploadPart ()
//
//     * CompleteMultipartUpload
// ()
//
//     * ListParts ()
//
//     * ListMultipartUploads ()
func (c *Client) AbortMultipartUpload(ctx context.Context, params *AbortMultipartUploadInput, optFns ...func(*Options)) (*AbortMultipartUploadOutput, error) {
	stack := middleware.NewStack("AbortMultipartUpload", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpAbortMultipartUploadMiddlewares(stack)
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
	addOpAbortMultipartUploadValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAbortMultipartUpload(options.Region), middleware.Before)
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
			OperationName: "AbortMultipartUpload",
			Err:           err,
		}
	}
	out := result.(*AbortMultipartUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AbortMultipartUploadInput struct {

	// The bucket name to which the upload was taking place. When using this API with
	// an access point, you must direct requests to the access point hostname. The
	// access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// operation using an access point through the AWS SDKs, you provide the access
	// point ARN in place of the bucket name. For more information about access point
	// ARNs, see Using Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) in
	// the Amazon Simple Storage Service Developer Guide.
	//
	// This member is required.
	Bucket *string

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer types.RequestPayer

	// Upload ID that identifies the multipart upload.
	//
	// This member is required.
	UploadId *string

	// Key of the object for which the multipart upload was initiated.
	//
	// This member is required.
	Key *string
}

type AbortMultipartUploadOutput struct {

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpAbortMultipartUploadMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpAbortMultipartUpload{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpAbortMultipartUpload{}, middleware.After)
}

func newServiceMetadataMiddleware_opAbortMultipartUpload(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "AbortMultipartUpload",
	}
}
