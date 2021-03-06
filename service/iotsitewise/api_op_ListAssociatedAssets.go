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

// Retrieves a paginated list of the assets associated to a parent asset (assetId)
// by a given hierarchy (hierarchyId).
func (c *Client) ListAssociatedAssets(ctx context.Context, params *ListAssociatedAssetsInput, optFns ...func(*Options)) (*ListAssociatedAssetsOutput, error) {
	stack := middleware.NewStack("ListAssociatedAssets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListAssociatedAssetsMiddlewares(stack)
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
	addOpListAssociatedAssetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAssociatedAssets(options.Region), middleware.Before)
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
			OperationName: "ListAssociatedAssets",
			Err:           err,
		}
	}
	out := result.(*ListAssociatedAssetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssociatedAssetsInput struct {

	// The ID of the parent asset.
	//
	// This member is required.
	AssetId *string

	// The maximum number of results to be returned per paginated request.
	MaxResults *int32

	// The hierarchy ID (of the parent asset model) whose associated assets are
	// returned. To find a hierarchy ID, use the DescribeAsset
	// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DescribeAsset.html)
	// or DescribeAssetModel
	// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DescribeAssetModel.html)
	// actions. For more information, see Asset Hierarchies
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-hierarchies.html)
	// in the AWS IoT SiteWise User Guide.
	//
	// This member is required.
	HierarchyId *string

	// The token to be used for the next set of paginated results.
	NextToken *string
}

type ListAssociatedAssetsOutput struct {

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// A list that summarizes the associated assets.
	//
	// This member is required.
	AssetSummaries []*types.AssociatedAssetsSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListAssociatedAssetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListAssociatedAssets{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListAssociatedAssets{}, middleware.After)
}

func newServiceMetadataMiddleware_opListAssociatedAssets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "ListAssociatedAssets",
	}
}
