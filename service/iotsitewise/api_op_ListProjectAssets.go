// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a paginated list of assets associated with an AWS IoT SiteWise Monitor
// project.
func (c *Client) ListProjectAssets(ctx context.Context, params *ListProjectAssetsInput, optFns ...func(*Options)) (*ListProjectAssetsOutput, error) {
	stack := middleware.NewStack("ListProjectAssets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListProjectAssetsMiddlewares(stack)
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
	addOpListProjectAssetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListProjectAssets(options.Region), middleware.Before)
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
			OperationName: "ListProjectAssets",
			Err:           err,
		}
	}
	out := result.(*ListProjectAssetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProjectAssetsInput struct {

	// The maximum number of results to be returned per paginated request.
	MaxResults *int32

	// The ID of the project.
	//
	// This member is required.
	ProjectId *string

	// The token to be used for the next set of paginated results.
	NextToken *string
}

type ListProjectAssetsOutput struct {

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// A list that contains the IDs of each asset associated with the project.
	//
	// This member is required.
	AssetIds []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListProjectAssetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListProjectAssets{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListProjectAssets{}, middleware.After)
}

func newServiceMetadataMiddleware_opListProjectAssets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "ListProjectAssets",
	}
}
