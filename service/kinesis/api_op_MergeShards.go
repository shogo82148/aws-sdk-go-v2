// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Merges two adjacent shards in a Kinesis data stream and combines them into a
// single shard to reduce the stream's capacity to ingest and transport data. Two
// shards are considered adjacent if the union of the hash key ranges for the two
// shards form a contiguous set with no gaps. For example, if you have two shards,
// one with a hash key range of 276...381 and the other with a hash key range of
// 382...454, then you could merge these two shards into a single shard that would
// have a hash key range of 276...454. After the merge, the single child shard
// receives data for all hash key values covered by the two parent shards.
// MergeShards is called when there is a need to reduce the overall capacity of a
// stream because of excess capacity that is not being used. You must specify the
// shard to be merged and the adjacent shard for a stream. For more information
// about merging shards, see Merge Two Shards
// (https://docs.aws.amazon.com/kinesis/latest/dev/kinesis-using-sdk-java-resharding-merge.html)
// in the Amazon Kinesis Data Streams Developer Guide. If the stream is in the
// ACTIVE state, you can call MergeShards. If a stream is in the CREATING,
// UPDATING, or DELETING state, MergeShards returns a ResourceInUseException. If
// the specified stream does not exist, MergeShards returns a
// ResourceNotFoundException. You can use DescribeStream () to check the state of
// the stream, which is returned in StreamStatus. MergeShards is an asynchronous
// operation. Upon receiving a MergeShards request, Amazon Kinesis Data Streams
// immediately returns a response and sets the StreamStatus to UPDATING. After the
// operation is completed, Kinesis Data Streams sets the StreamStatus to ACTIVE.
// Read and write operations continue to work while the stream is in the UPDATING
// state. You use DescribeStream () to determine the shard IDs that are specified
// in the MergeShards request. If you try to operate on too many streams in
// parallel using CreateStream (), DeleteStream (), MergeShards, or SplitShard (),
// you receive a LimitExceededException. MergeShards has a limit of five
// transactions per second per account.
func (c *Client) MergeShards(ctx context.Context, params *MergeShardsInput, optFns ...func(*Options)) (*MergeShardsOutput, error) {
	stack := middleware.NewStack("MergeShards", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpMergeShardsMiddlewares(stack)
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
	addOpMergeShardsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opMergeShards(options.Region), middleware.Before)
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
			OperationName: "MergeShards",
			Err:           err,
		}
	}
	out := result.(*MergeShardsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for MergeShards.
type MergeShardsInput struct {

	// The shard ID of the adjacent shard for the merge.
	//
	// This member is required.
	AdjacentShardToMerge *string

	// The shard ID of the shard to combine with the adjacent shard for the merge.
	//
	// This member is required.
	ShardToMerge *string

	// The name of the stream for the merge.
	//
	// This member is required.
	StreamName *string
}

type MergeShardsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpMergeShardsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpMergeShards{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpMergeShards{}, middleware.After)
}

func newServiceMetadataMiddleware_opMergeShards(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "MergeShards",
	}
}
