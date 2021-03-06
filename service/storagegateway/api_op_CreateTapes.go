// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates one or more virtual tapes. You write data to the virtual tapes and then
// archive the tapes. This operation is only supported in the tape gateway type.
// <note> <p>Cache storage must be allocated to the gateway before you can create
// virtual tapes. Use the <a>AddCache</a> operation to add cache storage to a
// gateway.</p> </note>
func (c *Client) CreateTapes(ctx context.Context, params *CreateTapesInput, optFns ...func(*Options)) (*CreateTapesOutput, error) {
	stack := middleware.NewStack("CreateTapes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateTapesMiddlewares(stack)
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
	addOpCreateTapesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTapes(options.Region), middleware.Before)
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
			OperationName: "CreateTapes",
			Err:           err,
		}
	}
	out := result.(*CreateTapesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateTapesInput
type CreateTapesInput struct {

	// The unique Amazon Resource Name (ARN) that represents the gateway to associate
	// the virtual tapes with. Use the ListGateways () operation to return a list of
	// gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string

	// The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used for
	// Amazon S3 server-side encryption. Storage Gateway does not support asymmetric
	// CMKs. This value can only be set when KMSEncrypted is true. Optional.
	KMSKey *string

	// The ID of the pool that you want to add your tape to for archiving. The tape in
	// this pool is archived in the S3 storage class that is associated with the pool.
	// When you use your backup application to eject the tape, the tape is archived
	// directly into the storage class (S3 Glacier or S3 Glacier Deep Archive) that
	// corresponds to the pool.  <p>Valid Values: <code>GLACIER</code> |
	// <code>DEEP_ARCHIVE</code> </p>
	PoolId *string

	// A prefix that you append to the barcode of the virtual tape you are creating.
	// This prefix makes the barcode unique.  <note> <p>The prefix must be 1 to 4
	// characters in length and must be one of the uppercase letters from A to Z.</p>
	// </note>
	//
	// This member is required.
	TapeBarcodePrefix *string

	// Set to true to use Amazon S3 server-side encryption with your own AWS KMS key,
	// or false to use a key managed by Amazon S3. Optional.  <p>Valid Values:
	// <code>true</code> | <code>false</code> </p>
	KMSEncrypted *bool

	// The number of virtual tapes that you want to create.
	//
	// This member is required.
	NumTapesToCreate *int32

	// The size, in bytes, of the virtual tapes that you want to create.  <note> <p>The
	// size must be aligned by gigabyte (1024*1024*1024 bytes).</p> </note>
	//
	// This member is required.
	TapeSizeInBytes *int64

	// A unique identifier that you use to retry a request. If you retry a request, use
	// the same ClientToken you specified in the initial request.  <note> <p>Using the
	// same <code>ClientToken</code> prevents creating the tape multiple times.</p>
	// </note>
	//
	// This member is required.
	ClientToken *string

	// A list of up to 50 tags that can be assigned to a virtual tape. Each tag is a
	// key-value pair.  <note> <p>Valid characters for key and value are letters,
	// spaces, and numbers representable in UTF-8 format, and the following special
	// characters: + - = . _ : / @. The maximum length of a tag's key is 128
	// characters, and the maximum length for a tag's value is 256.</p> </note>
	Tags []*types.Tag
}

// CreateTapeOutput
type CreateTapesOutput struct {

	// A list of unique Amazon Resource Names (ARNs) that represents the virtual tapes
	// that were created.
	TapeARNs []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateTapesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTapes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTapes{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateTapes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "CreateTapes",
	}
}
