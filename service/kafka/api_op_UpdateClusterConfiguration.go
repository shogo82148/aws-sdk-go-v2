// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the cluster with the configuration that is specified in the request
// body.
func (c *Client) UpdateClusterConfiguration(ctx context.Context, params *UpdateClusterConfigurationInput, optFns ...func(*Options)) (*UpdateClusterConfigurationOutput, error) {
	stack := middleware.NewStack("UpdateClusterConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateClusterConfigurationMiddlewares(stack)
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
	addOpUpdateClusterConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateClusterConfiguration(options.Region), middleware.Before)
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
			OperationName: "UpdateClusterConfiguration",
			Err:           err,
		}
	}
	out := result.(*UpdateClusterConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateClusterConfigurationInput struct {

	// The Amazon Resource Name (ARN) that uniquely identifies the cluster.
	//
	// This member is required.
	ClusterArn *string

	// The version of the cluster that needs to be updated.
	//
	// This member is required.
	CurrentVersion *string

	// Represents the configuration that you want MSK to use for the brokers in a
	// cluster.
	//
	// This member is required.
	ConfigurationInfo *types.ConfigurationInfo
}

type UpdateClusterConfigurationOutput struct {

	// The Amazon Resource Name (ARN) of the cluster operation.
	ClusterOperationArn *string

	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateClusterConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateClusterConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateClusterConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateClusterConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafka",
		OperationName: "UpdateClusterConfiguration",
	}
}
