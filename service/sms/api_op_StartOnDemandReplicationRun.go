// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts an on-demand replication run for the specified replication job. This
// replication run starts immediately. This replication run is in addition to the
// ones already scheduled. There is a limit on the number of on-demand replications
// runs you can request in a 24-hour period.
func (c *Client) StartOnDemandReplicationRun(ctx context.Context, params *StartOnDemandReplicationRunInput, optFns ...func(*Options)) (*StartOnDemandReplicationRunOutput, error) {
	stack := middleware.NewStack("StartOnDemandReplicationRun", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartOnDemandReplicationRunMiddlewares(stack)
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
	addOpStartOnDemandReplicationRunValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartOnDemandReplicationRun(options.Region), middleware.Before)
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
			OperationName: "StartOnDemandReplicationRun",
			Err:           err,
		}
	}
	out := result.(*StartOnDemandReplicationRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartOnDemandReplicationRunInput struct {

	// The identifier of the replication job.
	//
	// This member is required.
	ReplicationJobId *string

	// The description of the replication run.
	Description *string
}

type StartOnDemandReplicationRunOutput struct {

	// The identifier of the replication run.
	ReplicationRunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartOnDemandReplicationRunMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartOnDemandReplicationRun{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartOnDemandReplicationRun{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartOnDemandReplicationRun(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "StartOnDemandReplicationRun",
	}
}
