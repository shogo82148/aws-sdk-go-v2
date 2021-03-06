// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes one or more of your clusters.
func (c *Client) DescribeClusters(ctx context.Context, params *DescribeClustersInput, optFns ...func(*Options)) (*DescribeClustersOutput, error) {
	stack := middleware.NewStack("DescribeClusters", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeClustersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClusters(options.Region), middleware.Before)
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
			OperationName: "DescribeClusters",
			Err:           err,
		}
	}
	out := result.(*DescribeClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClustersInput struct {

	// Whether to include additional information about your clusters in the response.
	// If this field is omitted, the attachments, statistics, and tags are not
	// included. If ATTACHMENTS is specified, the attachments for the container
	// instances or tasks within the cluster are included. If SETTINGS is specified,
	// the settings for the cluster are included. If STATISTICS is specified, the
	// following additional information, separated by launch type, is included:
	//
	//     *
	// runningEC2TasksCount
	//
	//     * runningFargateTasksCount
	//
	//     *
	// pendingEC2TasksCount
	//
	//     * pendingFargateTasksCount
	//
	//     *
	// activeEC2ServiceCount
	//
	//     * activeFargateServiceCount
	//
	//     *
	// drainingEC2ServiceCount
	//
	//     * drainingFargateServiceCount
	//
	// If TAGS is
	// specified, the metadata tags associated with the cluster are included.
	Include []types.ClusterField

	// A list of up to 100 cluster names or full cluster Amazon Resource Name (ARN)
	// entries. If you do not specify a cluster, the default cluster is assumed.
	Clusters []*string
}

type DescribeClustersOutput struct {

	// Any failures associated with the call.
	Failures []*types.Failure

	// The list of clusters.
	Clusters []*types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeClustersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeClusters{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeClusters{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeClusters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "DescribeClusters",
	}
}
