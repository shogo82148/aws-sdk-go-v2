// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Reboots a replication instance. Rebooting results in a momentary outage, until
// the replication instance becomes available again.
func (c *Client) RebootReplicationInstance(ctx context.Context, params *RebootReplicationInstanceInput, optFns ...func(*Options)) (*RebootReplicationInstanceOutput, error) {
	stack := middleware.NewStack("RebootReplicationInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRebootReplicationInstanceMiddlewares(stack)
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
	addOpRebootReplicationInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRebootReplicationInstance(options.Region), middleware.Before)
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
			OperationName: "RebootReplicationInstance",
			Err:           err,
		}
	}
	out := result.(*RebootReplicationInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RebootReplicationInstanceInput struct {

	// The Amazon Resource Name (ARN) of the replication instance.
	//
	// This member is required.
	ReplicationInstanceArn *string

	// If this parameter is true, the reboot is conducted through a Multi-AZ failover.
	// (If the instance isn't configured for Multi-AZ, then you can't specify true.)
	ForceFailover *bool
}

type RebootReplicationInstanceOutput struct {

	// The replication instance that is being rebooted.
	ReplicationInstance *types.ReplicationInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRebootReplicationInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRebootReplicationInstance{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRebootReplicationInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opRebootReplicationInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "RebootReplicationInstance",
	}
}
