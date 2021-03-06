// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Changes the size of the cluster. You can change the cluster's type, or change
// the number or type of nodes. The default behavior is to use the elastic resize
// method. With an elastic resize, your cluster is available for read and write
// operations more quickly than with the classic resize method. Elastic resize
// operations have the following restrictions:
//
//     * You can only resize clusters
// of the following types:
//
//         * dc2.large
//
//         * dc2.8xlarge
//
//         *
// ds2.xlarge
//
//         * ds2.8xlarge
//
//         * ra3.4xlarge
//
//         *
// ra3.16xlarge
//
//     * The type of nodes that you add must match the node type for
// the cluster.
func (c *Client) ResizeCluster(ctx context.Context, params *ResizeClusterInput, optFns ...func(*Options)) (*ResizeClusterOutput, error) {
	stack := middleware.NewStack("ResizeCluster", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpResizeClusterMiddlewares(stack)
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
	addOpResizeClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResizeCluster(options.Region), middleware.Before)
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
			OperationName: "ResizeCluster",
			Err:           err,
		}
	}
	out := result.(*ResizeClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResizeClusterInput struct {

	// The new node type for the nodes you are adding. If not specified, the cluster's
	// current node type is used.
	NodeType *string

	// A boolean value indicating whether the resize operation is using the classic
	// resize process. If you don't provide this parameter or set the value to false,
	// the resize type is elastic.
	Classic *bool

	// The new number of nodes for the cluster.
	NumberOfNodes *int32

	// The unique identifier for the cluster to resize.
	//
	// This member is required.
	ClusterIdentifier *string

	// The new cluster type for the specified cluster.
	ClusterType *string
}

type ResizeClusterOutput struct {

	// Describes a cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpResizeClusterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpResizeCluster{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpResizeCluster{}, middleware.After)
}

func newServiceMetadataMiddleware_opResizeCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "ResizeCluster",
	}
}
