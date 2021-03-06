// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the mailbox permissions associated with a user, group, or resource
// mailbox.
func (c *Client) ListMailboxPermissions(ctx context.Context, params *ListMailboxPermissionsInput, optFns ...func(*Options)) (*ListMailboxPermissionsOutput, error) {
	stack := middleware.NewStack("ListMailboxPermissions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListMailboxPermissionsMiddlewares(stack)
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
	addOpListMailboxPermissionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListMailboxPermissions(options.Region), middleware.Before)
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
			OperationName: "ListMailboxPermissions",
			Err:           err,
		}
	}
	out := result.(*ListMailboxPermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMailboxPermissionsInput struct {

	// The identifier of the organization under which the user, group, or resource
	// exists.
	//
	// This member is required.
	OrganizationId *string

	// The identifier of the user, group, or resource for which to list mailbox
	// permissions.
	//
	// This member is required.
	EntityId *string

	// The token to use to retrieve the next page of results. The first call does not
	// contain any tokens.
	NextToken *string

	// The maximum number of results to return in a single call.
	MaxResults *int32
}

type ListMailboxPermissionsOutput struct {

	// The token to use to retrieve the next page of results. The value is "null" when
	// there are no more results to return.
	NextToken *string

	// One page of the user, group, or resource mailbox permissions.
	Permissions []*types.Permission

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListMailboxPermissionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListMailboxPermissions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMailboxPermissions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListMailboxPermissions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "ListMailboxPermissions",
	}
}
