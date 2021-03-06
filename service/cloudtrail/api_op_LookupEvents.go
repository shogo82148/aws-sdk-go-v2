// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Looks up management events
// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-concepts.html#cloudtrail-concepts-management-events)
// or CloudTrail Insights events
// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-concepts.html#cloudtrail-concepts-insights-events)
// that are captured by CloudTrail. You can look up events that occurred in a
// region within the last 90 days. Lookup supports the following attributes for
// management events:
//
//     * AWS access key
//
//     * Event ID
//
//     * Event name
//
//
// * Event source
//
//     * Read only
//
//     * Resource name
//
//     * Resource type
//
//     *
// User name
//
// Lookup supports the following attributes for Insights events:
//
//     *
// Event ID
//
//     * Event name
//
//     * Event source
//
// All attributes are optional. The
// default number of results returned is 50, with a maximum of 50 possible. The
// response includes a token that you can use to get the next page of results. The
// rate of lookup requests is limited to two per second per account. If this limit
// is exceeded, a throttling error occurs.
func (c *Client) LookupEvents(ctx context.Context, params *LookupEventsInput, optFns ...func(*Options)) (*LookupEventsOutput, error) {
	stack := middleware.NewStack("LookupEvents", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpLookupEventsMiddlewares(stack)
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
	addOpLookupEventsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opLookupEvents(options.Region), middleware.Before)
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
			OperationName: "LookupEvents",
			Err:           err,
		}
	}
	out := result.(*LookupEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains a request for LookupEvents.
type LookupEventsInput struct {

	// Contains a list of lookup attributes. Currently the list can contain only one
	// item.
	LookupAttributes []*types.LookupAttribute

	// Specifies that only events that occur after or at the specified time are
	// returned. If the specified start time is after the specified end time, an error
	// is returned.
	StartTime *time.Time

	// The token to use to get the next page of results after a previous API call. This
	// token must be passed in with the same parameters that were specified in the the
	// original call. For example, if the original call specified an AttributeKey of
	// 'Username' with a value of 'root', the call with NextToken should include those
	// same parameters.
	NextToken *string

	// Specifies the event category. If you do not specify an event category, events of
	// the category are not returned in the response. For example, if you do not
	// specify insight as the value of EventCategory, no Insights events are returned.
	EventCategory types.EventCategory

	// The number of events to return. Possible values are 1 through 50. The default is
	// 50.
	MaxResults *int32

	// Specifies that only events that occur before or at the specified time are
	// returned. If the specified end time is before the specified start time, an error
	// is returned.
	EndTime *time.Time
}

// Contains a response to a LookupEvents action.
type LookupEventsOutput struct {

	// A list of events returned based on the lookup attributes specified and the
	// CloudTrail event. The events list is sorted by time. The most recent event is
	// listed first.
	Events []*types.Event

	// The token to use to get the next page of results after a previous API call. If
	// the token does not appear, there are no more results to return. The token must
	// be passed in with the same parameters as the previous call. For example, if the
	// original call specified an AttributeKey of 'Username' with a value of 'root',
	// the call with NextToken should include those same parameters.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpLookupEventsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpLookupEvents{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpLookupEvents{}, middleware.After)
}

func newServiceMetadataMiddleware_opLookupEvents(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "LookupEvents",
	}
}
