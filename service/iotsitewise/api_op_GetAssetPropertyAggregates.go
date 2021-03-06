// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Gets aggregated values for an asset property. For more information, see Querying
// Aggregated Property Values
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/query-industrial-data.html#aggregates)
// in the AWS IoT SiteWise User Guide. To identify an asset property, you must
// specify one of the following:
//
//     * The assetId and propertyId of an asset
// property.
//
//     * A propertyAlias, which is a data stream alias (for example,
// /company/windfarm/3/turbine/7/temperature). To define an asset property's alias,
// see UpdateAssetProperty
// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetProperty.html).
func (c *Client) GetAssetPropertyAggregates(ctx context.Context, params *GetAssetPropertyAggregatesInput, optFns ...func(*Options)) (*GetAssetPropertyAggregatesOutput, error) {
	stack := middleware.NewStack("GetAssetPropertyAggregates", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetAssetPropertyAggregatesMiddlewares(stack)
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
	addOpGetAssetPropertyAggregatesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAssetPropertyAggregates(options.Region), middleware.Before)
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
			OperationName: "GetAssetPropertyAggregates",
			Err:           err,
		}
	}
	out := result.(*GetAssetPropertyAggregatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAssetPropertyAggregatesInput struct {

	// The ID of the asset property.
	PropertyId *string

	// The exclusive start of the range from which to query historical data, expressed
	// in seconds in Unix epoch time.
	//
	// This member is required.
	StartDate *time.Time

	// The maximum number of results to be returned per paginated request.
	MaxResults *int32

	// The inclusive end of the range from which to query historical data, expressed in
	// seconds in Unix epoch time.
	//
	// This member is required.
	EndDate *time.Time

	// The data aggregating function.
	//
	// This member is required.
	AggregateTypes []types.AggregateType

	// The ID of the asset.
	AssetId *string

	// The property alias that identifies the property, such as an OPC-UA server data
	// stream path (for example, /company/windfarm/3/turbine/7/temperature). For more
	// information, see Mapping Industrial Data Streams to Asset Properties
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/connect-data-streams.html)
	// in the AWS IoT SiteWise User Guide.
	PropertyAlias *string

	// The time interval over which to aggregate data.
	//
	// This member is required.
	Resolution *string

	// The token to be used for the next set of paginated results.
	NextToken *string

	// The chronological sorting order of the requested information.
	TimeOrdering types.TimeOrdering

	// The quality by which to filter asset data.
	Qualities []types.Quality
}

type GetAssetPropertyAggregatesOutput struct {

	// The requested aggregated values.
	//
	// This member is required.
	AggregatedValues []*types.AggregatedValue

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetAssetPropertyAggregatesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetAssetPropertyAggregates{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAssetPropertyAggregates{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAssetPropertyAggregates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "GetAssetPropertyAggregates",
	}
}
