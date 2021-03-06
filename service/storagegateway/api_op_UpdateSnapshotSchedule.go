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

// Updates a snapshot schedule configured for a gateway volume. This operation is
// only supported in the cached volume and stored volume gateway types.  <p>The
// default snapshot schedule for volume is once every 24 hours, starting at the
// creation time of the volume. You can use this API to change the snapshot
// schedule configured for the volume.</p> <p>In the request you must identify the
// gateway volume whose snapshot schedule you want to update, and the schedule
// information, including when you want the snapshot to begin on a day and the
// frequency (in hours) of snapshots.</p>
func (c *Client) UpdateSnapshotSchedule(ctx context.Context, params *UpdateSnapshotScheduleInput, optFns ...func(*Options)) (*UpdateSnapshotScheduleOutput, error) {
	stack := middleware.NewStack("UpdateSnapshotSchedule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateSnapshotScheduleMiddlewares(stack)
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
	addOpUpdateSnapshotScheduleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSnapshotSchedule(options.Region), middleware.Before)
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
			OperationName: "UpdateSnapshotSchedule",
			Err:           err,
		}
	}
	out := result.(*UpdateSnapshotScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing one or more of the following fields:  <ul> <li> <p>
// <a>UpdateSnapshotScheduleInput$Description</a> </p> </li> <li> <p>
// <a>UpdateSnapshotScheduleInput$RecurrenceInHours</a> </p> </li> <li> <p>
// <a>UpdateSnapshotScheduleInput$StartAt</a> </p> </li> <li> <p>
// <a>UpdateSnapshotScheduleInput$VolumeARN</a> </p> </li> </ul>
type UpdateSnapshotScheduleInput struct {

	// Frequency of snapshots. Specify the number of hours between snapshots.
	//
	// This member is required.
	RecurrenceInHours *int32

	// The Amazon Resource Name (ARN) of the volume. Use the ListVolumes () operation
	// to return a list of gateway volumes.
	//
	// This member is required.
	VolumeARN *string

	// The hour of the day at which the snapshot schedule begins represented as hh,
	// where hh is the hour (0 to 23). The hour of the day is in the time zone of the
	// gateway.
	//
	// This member is required.
	StartAt *int32

	// A list of up to 50 tags that can be assigned to a snapshot. Each tag is a
	// key-value pair.  <note> <p>Valid characters for key and value are letters,
	// spaces, and numbers representable in UTF-8 format, and the following special
	// characters: + - = . _ : / @. The maximum length of a tag's key is 128
	// characters, and the maximum length for a tag's value is 256.</p> </note>
	Tags []*types.Tag

	// Optional description of the snapshot that overwrites the existing description.
	Description *string
}

// A JSON object containing the Amazon Resource Name (ARN) of the updated storage
// volume.
type UpdateSnapshotScheduleOutput struct {

	// The Amazon Resource Name (ARN) of the volume. Use the ListVolumes () operation
	// to return a list of gateway volumes.
	VolumeARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateSnapshotScheduleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateSnapshotSchedule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateSnapshotSchedule{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateSnapshotSchedule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "UpdateSnapshotSchedule",
	}
}
