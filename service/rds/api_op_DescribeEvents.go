// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns events related to DB instances, DB security groups, DB snapshots, and DB
// parameter groups for the past 14 days. Events specific to a particular DB
// instance, DB security group, database snapshot, or DB parameter group can be
// obtained by providing the name as a parameter. By default, the past hour of
// events are returned.
func (c *Client) DescribeEvents(ctx context.Context, params *DescribeEventsInput, optFns ...func(*Options)) (*DescribeEventsOutput, error) {
	stack := middleware.NewStack("DescribeEvents", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeEventsMiddlewares(stack)
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
	addOpDescribeEventsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEvents(options.Region), middleware.Before)
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
			OperationName: "DescribeEvents",
			Err:           err,
		}
	}
	out := result.(*DescribeEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeEventsInput struct {

	// This parameter isn't currently supported.
	Filters []*types.Filter

	// The event source to retrieve events for. If no value is specified, all events
	// are returned.
	SourceType types.SourceType

	// An optional pagination token provided by a previous DescribeEvents request. If
	// this parameter is specified, the response includes only records beyond the
	// marker, up to the value specified by MaxRecords.
	Marker *string

	// The identifier of the event source for which events are returned. If not
	// specified, then all sources are included in the response. Constraints:
	//
	//     * If
	// SourceIdentifier is supplied, SourceType must also be provided.
	//
	//     * If the
	// source type is DBInstance, then a DBInstanceIdentifier must be supplied.
	//
	//     *
	// If the source type is DBSecurityGroup, a DBSecurityGroupName must be supplied.
	//
	//
	// * If the source type is DBParameterGroup, a DBParameterGroupName must be
	// supplied.
	//
	//     * If the source type is DBSnapshot, a DBSnapshotIdentifier must
	// be supplied.
	//
	//     * Can't end with a hyphen or contain two consecutive hyphens.
	SourceIdentifier *string

	// A list of event categories that trigger notifications for a event notification
	// subscription.
	EventCategories []*string

	// The beginning of the time interval to retrieve events for, specified in ISO 8601
	// format. For more information about ISO 8601, go to the ISO8601 Wikipedia page.
	// (http://en.wikipedia.org/wiki/ISO_8601) Example: 2009-07-08T18:00Z
	StartTime *time.Time

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// The end of the time interval for which to retrieve events, specified in ISO 8601
	// format. For more information about ISO 8601, go to the ISO8601 Wikipedia page.
	// (http://en.wikipedia.org/wiki/ISO_8601) Example: 2009-07-08T18:00Z
	EndTime *time.Time

	// The number of minutes to retrieve events for. Default: 60
	Duration *int32
}

// Contains the result of a successful invocation of the DescribeEvents action.
type DescribeEventsOutput struct {

	// A list of Event instances.
	Events []*types.Event

	// An optional pagination token provided by a previous Events request. If this
	// parameter is specified, the response includes only records beyond the marker, up
	// to the value specified by MaxRecords .
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeEventsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEvents{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEvents{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeEvents(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeEvents",
	}
}
