// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops GuardDuty monitoring for the specified member accounts. Use the
// StartMonitoringMembers operation to restart monitoring for those accounts.
func (c *Client) StopMonitoringMembers(ctx context.Context, params *StopMonitoringMembersInput, optFns ...func(*Options)) (*StopMonitoringMembersOutput, error) {
	stack := middleware.NewStack("StopMonitoringMembers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStopMonitoringMembersMiddlewares(stack)
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
	addOpStopMonitoringMembersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopMonitoringMembers(options.Region), middleware.Before)
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
			OperationName: "StopMonitoringMembers",
			Err:           err,
		}
	}
	out := result.(*StopMonitoringMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopMonitoringMembersInput struct {

	// A list of account IDs for the member accounts to stop monitoring.
	//
	// This member is required.
	AccountIds []*string

	// The unique ID of the detector associated with the GuardDuty master account that
	// is monitoring member accounts.
	//
	// This member is required.
	DetectorId *string
}

type StopMonitoringMembersOutput struct {

	// A list of objects that contain an accountId for each account that could not be
	// processed, and a result string that indicates why the account was not processed.
	//
	// This member is required.
	UnprocessedAccounts []*types.UnprocessedAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStopMonitoringMembersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStopMonitoringMembers{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStopMonitoringMembers{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopMonitoringMembers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "StopMonitoringMembers",
	}
}
