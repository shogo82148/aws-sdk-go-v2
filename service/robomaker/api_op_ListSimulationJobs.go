// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of simulation jobs. You can optionally provide filters to
// retrieve specific simulation jobs.
func (c *Client) ListSimulationJobs(ctx context.Context, params *ListSimulationJobsInput, optFns ...func(*Options)) (*ListSimulationJobsOutput, error) {
	stack := middleware.NewStack("ListSimulationJobs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListSimulationJobsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSimulationJobs(options.Region), middleware.Before)
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
			OperationName: "ListSimulationJobs",
			Err:           err,
		}
	}
	out := result.(*ListSimulationJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSimulationJobsInput struct {

	// The nextToken value returned from a previous paginated ListSimulationJobs
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This token should be treated as an opaque
	// identifier that is only used to retrieve the next items in a list and not for
	// other programmatic purposes.
	NextToken *string

	// When this parameter is used, ListSimulationJobs only returns maxResults results
	// in a single page along with a nextToken response element. The remaining results
	// of the initial request can be seen by sending another ListSimulationJobs request
	// with the returned nextToken value. This value can be between 1 and 1000. If this
	// parameter is not used, then ListSimulationJobs returns up to 1000 results and a
	// nextToken value if applicable.
	MaxResults *int32

	// Optional filters to limit results. The filter names status and
	// simulationApplicationName and robotApplicationName are supported. When
	// filtering, you must use the complete value of the filtered item. You can use up
	// to three filters, but they must be for the same named item. For example, if you
	// are looking for items with the status Preparing or the status Running.
	Filters []*types.Filter
}

type ListSimulationJobsOutput struct {

	// A list of simulation job summaries that meet the criteria of the request.
	//
	// This member is required.
	SimulationJobSummaries []*types.SimulationJobSummary

	// The nextToken value to include in a future ListSimulationJobs request. When the
	// results of a ListRobot request exceed maxResults, this value can be used to
	// retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListSimulationJobsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListSimulationJobs{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListSimulationJobs{}, middleware.After)
}

func newServiceMetadataMiddleware_opListSimulationJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "ListSimulationJobs",
	}
}
