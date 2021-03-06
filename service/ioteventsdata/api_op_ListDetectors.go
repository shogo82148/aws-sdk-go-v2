// Code generated by smithy-go-codegen DO NOT EDIT.

package ioteventsdata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ioteventsdata/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists detectors (the instances of a detector model).
func (c *Client) ListDetectors(ctx context.Context, params *ListDetectorsInput, optFns ...func(*Options)) (*ListDetectorsOutput, error) {
	stack := middleware.NewStack("ListDetectors", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListDetectorsMiddlewares(stack)
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
	addOpListDetectorsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDetectors(options.Region), middleware.Before)
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
			OperationName: "ListDetectors",
			Err:           err,
		}
	}
	out := result.(*ListDetectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDetectorsInput struct {

	// A filter that limits results to those detectors (instances) in the given state.
	StateName *string

	// The name of the detector model whose detectors (instances) are listed.
	//
	// This member is required.
	DetectorModelName *string

	// The token for the next set of results.
	NextToken *string

	// The maximum number of results to return at one time.
	MaxResults *int32
}

type ListDetectorsOutput struct {

	// A list of summary information about the detectors (instances).
	DetectorSummaries []*types.DetectorSummary

	// A token to retrieve the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListDetectorsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListDetectors{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListDetectors{}, middleware.After)
}

func newServiceMetadataMiddleware_opListDetectors(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ioteventsdata",
		OperationName: "ListDetectors",
	}
}
