// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops the running cluster that is specified by DBClusterIdentifier. The cluster
// must be in the available state. For more information, see Stopping and Starting
// an Amazon DocumentDB Cluster
// (https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-stop-start.html).
func (c *Client) StopDBCluster(ctx context.Context, params *StopDBClusterInput, optFns ...func(*Options)) (*StopDBClusterOutput, error) {
	stack := middleware.NewStack("StopDBCluster", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpStopDBClusterMiddlewares(stack)
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
	addOpStopDBClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopDBCluster(options.Region), middleware.Before)
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
			OperationName: "StopDBCluster",
			Err:           err,
		}
	}
	out := result.(*StopDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopDBClusterInput struct {

	// The identifier of the cluster to stop. Example: docdb-2019-05-28-15-24-52
	//
	// This member is required.
	DBClusterIdentifier *string
}

type StopDBClusterOutput struct {

	// Detailed information about a cluster.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpStopDBClusterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpStopDBCluster{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpStopDBCluster{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopDBCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "StopDBCluster",
	}
}
