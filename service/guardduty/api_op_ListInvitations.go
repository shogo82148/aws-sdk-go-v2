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

// Lists all GuardDuty membership invitations that were sent to the current AWS
// account.
func (c *Client) ListInvitations(ctx context.Context, params *ListInvitationsInput, optFns ...func(*Options)) (*ListInvitationsOutput, error) {
	stack := middleware.NewStack("ListInvitations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListInvitationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListInvitations(options.Region), middleware.Before)
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
			OperationName: "ListInvitations",
			Err:           err,
		}
	}
	out := result.(*ListInvitationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInvitationsInput struct {

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the list action. For subsequent calls to
	// the action, fill nextToken in the request with the value of NextToken from the
	// previous response to continue listing data.
	NextToken *string

	// You can use this parameter to indicate the maximum number of items that you want
	// in the response. The default value is 50. The maximum value is 50.
	MaxResults *int32
}

type ListInvitationsOutput struct {

	// The pagination parameter to be used on the next list operation to retrieve more
	// items.
	NextToken *string

	// A list of invitation descriptions.
	Invitations []*types.Invitation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListInvitationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListInvitations{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListInvitations{}, middleware.After)
}

func newServiceMetadataMiddleware_opListInvitations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "ListInvitations",
	}
}
