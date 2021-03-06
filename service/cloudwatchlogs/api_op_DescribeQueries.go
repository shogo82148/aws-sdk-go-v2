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

// Returns a list of CloudWatch Logs Insights queries that are scheduled,
// executing, or have been executed recently in this account. You can request all
// queries, or limit it to queries of a specific log group or queries with a
// certain status.
func (c *Client) DescribeQueries(ctx context.Context, params *DescribeQueriesInput, optFns ...func(*Options)) (*DescribeQueriesOutput, error) {
	stack := middleware.NewStack("DescribeQueries", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeQueriesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeQueries(options.Region), middleware.Before)
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
			OperationName: "DescribeQueries",
			Err:           err,
		}
	}
	out := result.(*DescribeQueriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeQueriesInput struct {

	// Limits the returned queries to only those that have the specified status. Valid
	// values are Cancelled, Complete, Failed, Running, and Scheduled.
	Status types.QueryStatus

	// Limits the returned queries to only those for the specified log group.
	LogGroupName *string

	// Limits the number of returned queries to the specified number.
	MaxResults *int32

	// The token for the next set of items to return. The token expires after 24 hours.
	NextToken *string
}

type DescribeQueriesOutput struct {

	// The list of queries that match the request.
	Queries []*types.QueryInfo

	// The token for the next set of items to return. The token expires after 24 hours.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeQueriesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeQueries{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeQueries{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeQueries(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "DescribeQueries",
	}
}
