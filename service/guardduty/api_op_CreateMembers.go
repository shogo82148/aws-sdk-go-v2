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

// Creates member accounts of the current AWS account by specifying a list of AWS
// account IDs. This step is a prerequisite for managing the associated member
// accounts either by invitation or through an organization. When using Create
// Members as an organizations delegated administrator this action will enable
// GuardDuty in the added member accounts, with the exception of the organization
// master account, which must enable GuardDuty prior to being added as a member. If
// you are adding accounts by invitation use this action after GuardDuty has been
// enabled in potential member accounts and before using Invite Members
// (https://docs.aws.amazon.com/guardduty/latest/APIReference/API_InviteMembers.html).
func (c *Client) CreateMembers(ctx context.Context, params *CreateMembersInput, optFns ...func(*Options)) (*CreateMembersOutput, error) {
	stack := middleware.NewStack("CreateMembers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateMembersMiddlewares(stack)
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
	addOpCreateMembersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMembers(options.Region), middleware.Before)
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
			OperationName: "CreateMembers",
			Err:           err,
		}
	}
	out := result.(*CreateMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMembersInput struct {

	// A list of account ID and email address pairs of the accounts that you want to
	// associate with the master GuardDuty account.
	//
	// This member is required.
	AccountDetails []*types.AccountDetail

	// The unique ID of the detector of the GuardDuty account that you want to
	// associate member accounts with.
	//
	// This member is required.
	DetectorId *string
}

type CreateMembersOutput struct {

	// A list of objects that include the accountIds of the unprocessed accounts and a
	// result string that explains why each was unprocessed.
	//
	// This member is required.
	UnprocessedAccounts []*types.UnprocessedAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateMembersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateMembers{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMembers{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateMembers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "CreateMembers",
	}
}
