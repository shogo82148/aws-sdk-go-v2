// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Get an aggregation of service statistics defined by a specific time range.
func (c *Client) GetTimeSeriesServiceStatistics(ctx context.Context, params *GetTimeSeriesServiceStatisticsInput, optFns ...func(*Options)) (*GetTimeSeriesServiceStatisticsOutput, error) {
	stack := middleware.NewStack("GetTimeSeriesServiceStatistics", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetTimeSeriesServiceStatisticsMiddlewares(stack)
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
	addOpGetTimeSeriesServiceStatisticsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTimeSeriesServiceStatistics(options.Region), middleware.Before)
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
			OperationName: "GetTimeSeriesServiceStatistics",
			Err:           err,
		}
	}
	out := result.(*GetTimeSeriesServiceStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTimeSeriesServiceStatisticsInput struct {

	// The case-sensitive name of the group for which to pull statistics from.
	GroupName *string

	// Pagination token.
	NextToken *string

	// Aggregation period in seconds.
	Period *int32

	// The ARN of the group for which to pull statistics from.
	GroupARN *string

	// The start of the time frame for which to aggregate statistics.
	//
	// This member is required.
	StartTime *time.Time

	// The end of the time frame for which to aggregate statistics.
	//
	// This member is required.
	EndTime *time.Time

	// A filter expression defining entities that will be aggregated for statistics.
	// Supports ID, service, and edge functions. If no selector expression is
	// specified, edge statistics are returned.
	EntitySelectorExpression *string
}

type GetTimeSeriesServiceStatisticsOutput struct {

	// Pagination token.
	NextToken *string

	// The collection of statistics.
	TimeSeriesServiceStatistics []*types.TimeSeriesServiceStatistics

	// A flag indicating whether or not a group's filter expression has been
	// consistent, or if a returned aggregation may show statistics from an older
	// version of the group's filter expression.
	ContainsOldGroupVersions *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetTimeSeriesServiceStatisticsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetTimeSeriesServiceStatistics{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTimeSeriesServiceStatistics{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetTimeSeriesServiceStatistics(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "GetTimeSeriesServiceStatistics",
	}
}
