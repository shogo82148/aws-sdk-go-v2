// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns the data points of a specific metric for an Amazon Lightsail content
// delivery network (CDN) distribution.  <p>Metrics report the utilization of your
// resources, and the error counts generated by them. Monitor and collect metric
// data regularly to maintain the reliability, availability, and performance of
// your resources.</p>
func (c *Client) GetDistributionMetricData(ctx context.Context, params *GetDistributionMetricDataInput, optFns ...func(*Options)) (*GetDistributionMetricDataOutput, error) {
	stack := middleware.NewStack("GetDistributionMetricData", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetDistributionMetricDataMiddlewares(stack)
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
	addOpGetDistributionMetricDataValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDistributionMetricData(options.Region), middleware.Before)
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
			OperationName: "GetDistributionMetricData",
			Err:           err,
		}
	}
	out := result.(*GetDistributionMetricDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDistributionMetricDataInput struct {

	// The metric for which you want to return information.  <p>Valid distribution
	// metric names are listed below, along with the most useful
	// <code>statistics</code> to include in your request, and the published
	// <code>unit</code> value.</p> <ul> <li> <p> <b> <code>Requests</code> </b> - The
	// total number of viewer requests received by your Lightsail distribution, for all
	// HTTP methods, and for both HTTP and HTTPS requests.</p> <p>
	// <code>Statistics</code>: The most useful statistic is <code>Sum</code>.</p> <p>
	// <code>Unit</code>: The published unit is <code>None</code>.</p> </li> <li> <p>
	// <b> <code>BytesDownloaded</code> </b> - The number of bytes downloaded by
	// viewers for GET, HEAD, and OPTIONS requests.</p> <p> <code>Statistics</code>:
	// The most useful statistic is <code>Sum</code>.</p> <p> <code>Unit</code>: The
	// published unit is <code>None</code>.</p> </li> <li> <p> <b> <code>BytesUploaded
	// </code> </b> - The number of bytes uploaded to your origin by your Lightsail
	// distribution, using POST and PUT requests.</p> <p> <code>Statistics</code>: The
	// most useful statistic is <code>Sum</code>.</p> <p> <code>Unit</code>: The
	// published unit is <code>None</code>.</p> </li> <li> <p> <b>
	// <code>TotalErrorRate</code> </b> - The percentage of all viewer requests for
	// which the response's HTTP status code was 4xx or 5xx.</p> <p>
	// <code>Statistics</code>: The most useful statistic is <code>Average</code>.</p>
	// <p> <code>Unit</code>: The published unit is <code>Percent</code>.</p> </li>
	// <li> <p> <b> <code>4xxErrorRate</code> </b> - The percentage of all viewer
	// requests for which the response's HTTP status cod was 4xx. In these cases, the
	// client or client viewer may have made an error. For example, a status code of
	// 404 (Not Found) means that the client requested an object that could not be
	// found.</p> <p> <code>Statistics</code>: The most useful statistic is
	// <code>Average</code>.</p> <p> <code>Unit</code>: The published unit is
	// <code>Percent</code>.</p> </li> <li> <p> <b> <code>5xxErrorRate</code> </b> -
	// The percentage of all viewer requests for which the response's HTTP status code
	// was 5xx. In these cases, the origin server did not satisfy the requests. For
	// example, a status code of 503 (Service Unavailable) means that the origin server
	// is currently unavailable.</p> <p> <code>Statistics</code>: The most useful
	// statistic is <code>Average</code>.</p> <p> <code>Unit</code>: The published unit
	// is <code>Percent</code>.</p> </li> </ul>
	//
	// This member is required.
	MetricName types.DistributionMetricName

	// The end of the time interval for which to get metric data. Constraints:
	//
	//     *
	// Specified in Coordinated Universal Time (UTC).
	//
	//     * Specified in the Unix time
	// format. For example, if you wish to use an end time of October 1, 2018, at 9 PM
	// UTC, specify 1538427600 as the end time.
	//
	// You can convert a human-friendly time
	// to Unix time format using a converter like Epoch converter
	// (https://www.epochconverter.com/).
	//
	// This member is required.
	EndTime *time.Time

	// The statistic for the metric. The following statistics are available:
	//
	//     *
	// Minimum - The lowest value observed during the specified period. Use this value
	// to determine low volumes of activity for your application.
	//
	//     * Maximum - The
	// highest value observed during the specified period. Use this value to determine
	// high volumes of activity for your application.
	//
	//     * Sum - All values submitted
	// for the matching metric added together. You can use this statistic to determine
	// the total volume of a metric.
	//
	//     * Average - The value of Sum / SampleCount
	// during the specified period. By comparing this statistic with the Minimum and
	// Maximum values, you can determine the full scope of a metric and how close the
	// average use is to the Minimum and Maximum values. This comparison helps you to
	// know when to increase or decrease your resources.
	//
	//     * SampleCount - The
	// count, or number, of data points used for the statistical calculation.
	//
	// This member is required.
	Statistics []types.MetricStatistic

	// The start of the time interval for which to get metric data. Constraints:
	//
	//     *
	// Specified in Coordinated Universal Time (UTC).
	//
	//     * Specified in the Unix time
	// format. For example, if you wish to use a start time of October 1, 2018, at 8 PM
	// UTC, specify 1538424000 as the start time.
	//
	// You can convert a human-friendly
	// time to Unix time format using a converter like Epoch converter
	// (https://www.epochconverter.com/).
	//
	// This member is required.
	StartTime *time.Time

	// The unit for the metric data request.  <p>Valid units depend on the metric data
	// being requested. For the valid units with each available metric, see the
	// <code>metricName</code> parameter.</p>
	//
	// This member is required.
	Unit types.MetricUnit

	// The granularity, in seconds, for the metric data points that will be returned.
	//
	// This member is required.
	Period *int32

	// The name of the distribution for which to get metric data.  <p>Use the
	// <code>GetDistributions</code> action to get a list of distribution names that
	// you can specify.</p>
	//
	// This member is required.
	DistributionName *string
}

type GetDistributionMetricDataOutput struct {

	// An array of objects that describe the metric data returned.
	MetricData []*types.MetricDatapoint

	// The name of the metric returned.
	MetricName types.DistributionMetricName

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetDistributionMetricDataMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetDistributionMetricData{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDistributionMetricData{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDistributionMetricData(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetDistributionMetricData",
	}
}
