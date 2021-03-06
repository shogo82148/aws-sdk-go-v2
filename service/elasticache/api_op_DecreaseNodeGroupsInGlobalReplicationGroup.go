// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Decreases the number of node groups in a Global Datastore
func (c *Client) DecreaseNodeGroupsInGlobalReplicationGroup(ctx context.Context, params *DecreaseNodeGroupsInGlobalReplicationGroupInput, optFns ...func(*Options)) (*DecreaseNodeGroupsInGlobalReplicationGroupOutput, error) {
	stack := middleware.NewStack("DecreaseNodeGroupsInGlobalReplicationGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDecreaseNodeGroupsInGlobalReplicationGroupMiddlewares(stack)
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
	addOpDecreaseNodeGroupsInGlobalReplicationGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDecreaseNodeGroupsInGlobalReplicationGroup(options.Region), middleware.Before)
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
			OperationName: "DecreaseNodeGroupsInGlobalReplicationGroup",
			Err:           err,
		}
	}
	out := result.(*DecreaseNodeGroupsInGlobalReplicationGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DecreaseNodeGroupsInGlobalReplicationGroupInput struct {

	// If the value of NodeGroupCount is less than the current number of node groups
	// (shards), then either NodeGroupsToRemove or NodeGroupsToRetain is required.
	// NodeGroupsToRemove is a list of NodeGroupIds to remove from the cluster.
	// ElastiCache for Redis will attempt to remove all node groups listed by
	// NodeGroupsToRemove from the cluster. </p>
	GlobalNodeGroupsToRemove []*string

	// The name of the Global Datastore
	//
	// This member is required.
	GlobalReplicationGroupId *string

	// Indicates that the shard reconfiguration process begins immediately. At present,
	// the only permitted value for this parameter is true.
	//
	// This member is required.
	ApplyImmediately *bool

	// The number of node groups (shards) that results from the modification of the
	// shard configuration
	//
	// This member is required.
	NodeGroupCount *int32

	// If the value of NodeGroupCount is less than the current number of node groups
	// (shards), then either NodeGroupsToRemove or NodeGroupsToRetain is required.
	// NodeGroupsToRemove is a list of NodeGroupIds to remove from the cluster.
	// ElastiCache for Redis will attempt to remove all node groups listed by
	// NodeGroupsToRemove from the cluster. </p>
	GlobalNodeGroupsToRetain []*string
}

type DecreaseNodeGroupsInGlobalReplicationGroupOutput struct {

	// Consists of a primary cluster that accepts writes and an associated secondary
	// cluster that resides in a different AWS region. The secondary cluster accepts
	// only reads. The primary cluster automatically replicates updates to the
	// secondary cluster.  <ul> <li> <p>The <b>GlobalReplicationGroupIdSuffix</b>
	// represents the name of the Global Datastore, which is what you use to associate
	// a secondary cluster.</p> </li> </ul>
	GlobalReplicationGroup *types.GlobalReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDecreaseNodeGroupsInGlobalReplicationGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDecreaseNodeGroupsInGlobalReplicationGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDecreaseNodeGroupsInGlobalReplicationGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opDecreaseNodeGroupsInGlobalReplicationGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DecreaseNodeGroupsInGlobalReplicationGroup",
	}
}
