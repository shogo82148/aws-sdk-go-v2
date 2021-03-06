// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Declines invitations to become a member account.
func (c *Client) DeclineInvitations(ctx context.Context, params *DeclineInvitationsInput, optFns ...func(*Options)) (*DeclineInvitationsOutput, error) {
	stack := middleware.NewStack("DeclineInvitations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeclineInvitationsMiddlewares(stack)
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
	addOpDeclineInvitationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeclineInvitations(options.Region), middleware.Before)
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
			OperationName: "DeclineInvitations",
			Err:           err,
		}
	}
	out := result.(*DeclineInvitationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeclineInvitationsInput struct {

	// The list of account IDs for the accounts from which to decline the invitations
	// to Security Hub.
	//
	// This member is required.
	AccountIds []*string
}

type DeclineInvitationsOutput struct {

	// The list of AWS accounts that were not processed. For each account, the list
	// includes the account ID and the email address.
	UnprocessedAccounts []*types.Result

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeclineInvitationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeclineInvitations{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeclineInvitations{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeclineInvitations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "DeclineInvitations",
	}
}
