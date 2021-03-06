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
)

// Gets an asset property's current value. For more information, see Querying
// Current Property Values
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/query-industrial-data.html#current-values)
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
func (c *Client) GetAssetPropertyValue(ctx context.Context, params *GetAssetPropertyValueInput, optFns ...func(*Options)) (*GetAssetPropertyValueOutput, error) {
	stack := middleware.NewStack("GetAssetPropertyValue", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetAssetPropertyValueMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAssetPropertyValue(options.Region), middleware.Before)
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
			OperationName: "GetAssetPropertyValue",
			Err:           err,
		}
	}
	out := result.(*GetAssetPropertyValueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAssetPropertyValueInput struct {

	// The ID of the asset.
	AssetId *string

	// The property alias that identifies the property, such as an OPC-UA server data
	// stream path (for example, /company/windfarm/3/turbine/7/temperature). For more
	// information, see Mapping Industrial Data Streams to Asset Properties
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/connect-data-streams.html)
	// in the AWS IoT SiteWise User Guide.
	PropertyAlias *string

	// The ID of the asset property.
	PropertyId *string
}

type GetAssetPropertyValueOutput struct {

	// The current asset property value.
	PropertyValue *types.AssetPropertyValue

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetAssetPropertyValueMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetAssetPropertyValue{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAssetPropertyValue{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAssetPropertyValue(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "GetAssetPropertyValue",
	}
}
