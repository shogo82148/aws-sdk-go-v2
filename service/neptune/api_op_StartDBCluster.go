// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts an Amazon Neptune DB cluster that was stopped using the AWS console, the
// AWS CLI stop-db-cluster command, or the StopDBCluster API.
func (c *Client) StartDBCluster(ctx context.Context, params *StartDBClusterInput, optFns ...func(*Options)) (*StartDBClusterOutput, error) {
	stack := middleware.NewStack("StartDBCluster", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpStartDBClusterMiddlewares(stack)
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
	addOpStartDBClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartDBCluster(options.Region), middleware.Before)
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
			OperationName: "StartDBCluster",
			Err:           err,
		}
	}
	out := result.(*StartDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDBClusterInput struct {

	// The DB cluster identifier of the Neptune DB cluster to be started. This
	// parameter is stored as a lowercase string.
	//
	// This member is required.
	DBClusterIdentifier *string
}

type StartDBClusterOutput struct {

	// Contains the details of an Amazon Neptune DB cluster. This data type is used as
	// a response element in the DescribeDBClusters () action.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpStartDBClusterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpStartDBCluster{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpStartDBCluster{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartDBCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "StartDBCluster",
	}
}
