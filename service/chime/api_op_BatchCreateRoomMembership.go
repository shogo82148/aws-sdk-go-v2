// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds up to 50 members to a chat room in an Amazon Chime Enterprise account.
// Members can be either users or bots. The member role designates whether the
// member is a chat room administrator or a general chat room member.
func (c *Client) BatchCreateRoomMembership(ctx context.Context, params *BatchCreateRoomMembershipInput, optFns ...func(*Options)) (*BatchCreateRoomMembershipOutput, error) {
	stack := middleware.NewStack("BatchCreateRoomMembership", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpBatchCreateRoomMembershipMiddlewares(stack)
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
	addOpBatchCreateRoomMembershipValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchCreateRoomMembership(options.Region), middleware.Before)
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
			OperationName: "BatchCreateRoomMembership",
			Err:           err,
		}
	}
	out := result.(*BatchCreateRoomMembershipOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchCreateRoomMembershipInput struct {

	// The list of membership items.
	//
	// This member is required.
	MembershipItemList []*types.MembershipItem

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The room ID.
	//
	// This member is required.
	RoomId *string
}

type BatchCreateRoomMembershipOutput struct {

	// If the action fails for one or more of the member IDs in the request, a list of
	// the member IDs is returned, along with error codes and error messages.
	Errors []*types.MemberError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpBatchCreateRoomMembershipMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpBatchCreateRoomMembership{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchCreateRoomMembership{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchCreateRoomMembership(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "BatchCreateRoomMembership",
	}
}
