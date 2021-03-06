// Code generated by smithy-go-codegen DO NOT EDIT.

package iotevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotevents/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the detector models you have created. Only the metadata associated with
// each detector model is returned.
func (c *Client) ListDetectorModels(ctx context.Context, params *ListDetectorModelsInput, optFns ...func(*Options)) (*ListDetectorModelsOutput, error) {
	stack := middleware.NewStack("ListDetectorModels", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListDetectorModelsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDetectorModels(options.Region), middleware.Before)
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
			OperationName: "ListDetectorModels",
			Err:           err,
		}
	}
	out := result.(*ListDetectorModelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDetectorModelsInput struct {

	// The token for the next set of results.
	NextToken *string

	// The maximum number of results to return at one time.
	MaxResults *int32
}

type ListDetectorModelsOutput struct {

	// Summary information about the detector models.
	DetectorModelSummaries []*types.DetectorModelSummary

	// A token to retrieve the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListDetectorModelsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListDetectorModels{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListDetectorModels{}, middleware.After)
}

func newServiceMetadataMiddleware_opListDetectorModels(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotevents",
		OperationName: "ListDetectorModels",
	}
}
