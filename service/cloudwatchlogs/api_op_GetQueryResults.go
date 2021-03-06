// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the results from the specified query. Only the fields requested in the
// query are returned, along with a @ptr field which is the identifier for the log
// record. You can use the value of @ptr in a GetLogRecord
// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetLogRecord.html)
// operation to get the full log record. GetQueryResults does not start a query
// execution. To run a query, use StartQuery
// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_StartQuery.html).
// If the value of the Status field in the output is Running, this operation
// returns only partial results. If you see a value of Scheduled or Running for the
// status, you can retry the operation later to see the final results.
func (c *Client) GetQueryResults(ctx context.Context, params *GetQueryResultsInput, optFns ...func(*Options)) (*GetQueryResultsOutput, error) {
	stack := middleware.NewStack("GetQueryResults", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetQueryResultsMiddlewares(stack)
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
	addOpGetQueryResultsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetQueryResults(options.Region), middleware.Before)
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
			OperationName: "GetQueryResults",
			Err:           err,
		}
	}
	out := result.(*GetQueryResultsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetQueryResultsInput struct {

	// The ID number of the query.
	//
	// This member is required.
	QueryId *string
}

type GetQueryResultsOutput struct {

	// The log events that matched the query criteria during the most recent time it
	// ran. The results value is an array of arrays. Each log event is one object in
	// the top-level array. Each of these log event objects is an array of field/value
	// pairs.
	Results [][]*types.ResultField

	// The status of the most recent running of the query. Possible values are
	// Cancelled, Complete, Failed, Running, Scheduled, Timeout, and Unknown. Queries
	// time out after 15 minutes of execution. To avoid having your queries time out,
	// reduce the time range being searched, or partition your query into a number of
	// queries.
	Status types.QueryStatus

	// Includes the number of log events scanned by the query, the number of log events
	// that matched the query criteria, and the total number of bytes in the log events
	// that were scanned.
	Statistics *types.QueryStatistics

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetQueryResultsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetQueryResults{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetQueryResults{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetQueryResults(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "GetQueryResults",
	}
}
