// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists profiling groups.
func (c *Client) ListProfilingGroups(ctx context.Context, params *ListProfilingGroupsInput, optFns ...func(*Options)) (*ListProfilingGroupsOutput, error) {
	stack := middleware.NewStack("ListProfilingGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListProfilingGroupsMiddlewares(stack)
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
			OperationName: "ListProfilingGroups",
			Err:           err,
		}
	}
	out := result.(*ListProfilingGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The structure representing the listProfilingGroupsRequest.
type ListProfilingGroupsInput struct {

	// A Boolean value indicating whether to include a description.
	IncludeDescription *bool

	// The maximum number of profiling groups results returned by ListProfilingGroups
	// in paginated output. When this parameter is used, ListProfilingGroups only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListProfilingGroups request with the returned nextToken value.
	MaxResults *int32

	// The nextToken value returned from a previous paginated ListProfilingGroups
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This token should be treated as an opaque
	// identifier that is only used to retrieve the next items in a list and not for
	// other programmatic purposes.
	NextToken *string
}

// The structure representing the listProfilingGroupsResponse.
type ListProfilingGroupsOutput struct {

	// The nextToken value to include in a future ListProfilingGroups request. When the
	// results of a ListProfilingGroups request exceed maxResults, this value can be
	// used to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string

	// Information about profiling group names.
	//
	// This member is required.
	ProfilingGroupNames []*string

	// Information about profiling groups.
	ProfilingGroups []*types.ProfilingGroupDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListProfilingGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListProfilingGroups{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListProfilingGroups{}, middleware.After)
}
