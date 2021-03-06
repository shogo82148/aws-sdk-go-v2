// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of cache policies. You can optionally apply a filter to return only
// the managed policies created by AWS, or only the custom policies created in your
// AWS account. You can optionally specify the maximum number of items to receive
// in the response. If the total number of items in the list exceeds the maximum
// that you specify, or the default maximum, the response is paginated. To get the
// next page of items, send a subsequent request that specifies the NextMarker
// value from the current response as the Marker value in the subsequent request.
func (c *Client) ListCachePolicies(ctx context.Context, params *ListCachePoliciesInput, optFns ...func(*Options)) (*ListCachePoliciesOutput, error) {
	stack := middleware.NewStack("ListCachePolicies", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpListCachePoliciesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListCachePolicies(options.Region), middleware.Before)
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
			OperationName: "ListCachePolicies",
			Err:           err,
		}
	}
	out := result.(*ListCachePoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCachePoliciesInput struct {

	// Use this field when paginating results to indicate where to begin in your list
	// of cache policies. The response includes cache policies in the list that occur
	// after the marker. To get the next page of the list, set this field’s value to
	// the value of NextMarker from the current page’s response.
	Marker *string

	// The maximum number of cache policies that you want in the response.
	MaxItems *string

	// A filter to return only the specified kinds of cache policies. Valid values
	// are:
	//
	//     * managed – Returns only the managed policies created by AWS.
	//
	//     *
	// custom – Returns only the custom policies created in your AWS account.
	Type types.CachePolicyType
}

type ListCachePoliciesOutput struct {

	// A list of cache policies.
	CachePolicyList *types.CachePolicyList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpListCachePoliciesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpListCachePolicies{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpListCachePolicies{}, middleware.After)
}

func newServiceMetadataMiddleware_opListCachePolicies(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "ListCachePolicies",
	}
}
