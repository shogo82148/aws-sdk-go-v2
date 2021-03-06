// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes one or more flow logs. To view the information in your flow logs (the
// log streams for the network interfaces), you must use the CloudWatch Logs
// console or the CloudWatch Logs API.
func (c *Client) DescribeFlowLogs(ctx context.Context, params *DescribeFlowLogsInput, optFns ...func(*Options)) (*DescribeFlowLogsOutput, error) {
	stack := middleware.NewStack("DescribeFlowLogs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeFlowLogsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFlowLogs(options.Region), middleware.Before)
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
			OperationName: "DescribeFlowLogs",
			Err:           err,
		}
	}
	out := result.(*DescribeFlowLogsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFlowLogsInput struct {

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// One or more flow log IDs. Constraint: Maximum of 1000 flow log IDs.
	FlowLogIds []*string

	// The token for the next page of results.
	NextToken *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	//     * deliver-log-status - The status of the logs delivery
	// (SUCCESS | FAILED).
	//
	//     * log-destination-type - The type of destination to
	// which the flow log publishes data. Possible destination types include
	// cloud-watch-logs and S3.
	//
	//     * flow-log-id - The ID of the flow log.
	//
	//     *
	// log-group-name - The name of the log group.
	//
	//     * resource-id - The ID of the
	// VPC, subnet, or network interface.
	//
	//     * traffic-type - The type of traffic
	// (ACCEPT | REJECT | ALL).
	//
	//     * tag: - The key/value combination of a tag
	// assigned to the resource. Use the tag key in the filter name and the tag value
	// as the filter value. For example, to find all resources that have a tag with the
	// key Owner and the value TeamA, specify tag:Owner for the filter name and TeamA
	// for the filter value.
	//
	//     * tag-key - The key of a tag assigned to the
	// resource. Use this filter to find all resources assigned a tag with a specific
	// key, regardless of the tag value.
	Filter []*types.Filter
}

type DescribeFlowLogsOutput struct {

	// Information about the flow logs.
	FlowLogs []*types.FlowLog

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeFlowLogsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeFlowLogs{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeFlowLogs{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeFlowLogs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeFlowLogs",
	}
}
