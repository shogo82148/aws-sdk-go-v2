// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Posts updates to records and adds and deletes records for a dataset and user.
// The sync count in the record patch is your last known sync count for that
// record. The server will reject an UpdateRecords request with a
// ResourceConflictException if you try to patch a record with a new value but a
// stale sync count. For example, if the sync count on the server is 5 for a key
// called highScore and you try and submit a new highScore with sync count of 4,
// the request will be rejected. To obtain the current sync count for a record,
// call ListRecords. On a successful update of the record, the response returns the
// new sync count for that record. You should present that sync count the next time
// you try to update that same record. When the record does not exist, specify the
// sync count as 0. This API can be called with temporary user credentials provided
// by Cognito Identity or with developer credentials.
func (c *Client) UpdateRecords(ctx context.Context, params *UpdateRecordsInput, optFns ...func(*Options)) (*UpdateRecordsOutput, error) {
	stack := middleware.NewStack("UpdateRecords", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateRecordsMiddlewares(stack)
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
	addOpUpdateRecordsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRecords(options.Region), middleware.Before)
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
			OperationName: "UpdateRecords",
			Err:           err,
		}
	}
	out := result.(*UpdateRecordsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to post updates to records or add and delete records for a dataset and
// user.
type UpdateRecordsInput struct {

	// A list of patch operations.
	RecordPatches []*types.RecordPatch

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// This member is required.
	IdentityPoolId *string

	// A string of up to 128 characters. Allowed characters are a-z, A-Z, 0-9, '_'
	// (underscore), '-' (dash), and '.' (dot).
	//
	// This member is required.
	DatasetName *string

	// Intended to supply a device ID that will populate the lastModifiedBy field
	// referenced in other methods. The ClientContext field is not yet implemented.
	ClientContext *string

	// The unique ID generated for this device by Cognito.
	DeviceId *string

	// The SyncSessionToken returned by a previous call to ListRecords for this dataset
	// and identity.
	//
	// This member is required.
	SyncSessionToken *string

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// This member is required.
	IdentityId *string
}

// Returned for a successful UpdateRecordsRequest.
type UpdateRecordsOutput struct {

	// A list of records that have been updated.
	Records []*types.Record

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateRecordsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRecords{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRecords{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateRecords(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "UpdateRecords",
	}
}
