// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a MemberAccounts object that lists the member accounts in the
// administrator's AWS organization. The ListMemberAccounts must be submitted by
// the account that is set as the AWS Firewall Manager administrator.
func (c *Client) ListMemberAccounts(ctx context.Context, params *ListMemberAccountsInput, optFns ...func(*Options)) (*ListMemberAccountsOutput, error) {
	stack := middleware.NewStack("ListMemberAccounts", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListMemberAccountsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListMemberAccounts(options.Region), middleware.Before)
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
			OperationName: "ListMemberAccounts",
			Err:           err,
		}
	}
	out := result.(*ListMemberAccountsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMemberAccountsInput struct {

	// Specifies the number of member account IDs that you want AWS Firewall Manager to
	// return for this request. If you have more IDs than the number that you specify
	// for MaxResults, the response includes a NextToken value that you can use to get
	// another batch of member account IDs.
	MaxResults *int32

	// If you specify a value for MaxResults and you have more account IDs than the
	// number that you specify for MaxResults, AWS Firewall Manager returns a NextToken
	// value in the response that allows you to list another group of IDs. For the
	// second and subsequent ListMemberAccountsRequest requests, specify the value of
	// NextToken from the previous response to get information about another batch of
	// member account IDs.
	NextToken *string
}

type ListMemberAccountsOutput struct {

	// If you have more member account IDs than the number that you specified for
	// MaxResults in the request, the response includes a NextToken value. To list more
	// IDs, submit another ListMemberAccounts request, and specify the NextToken value
	// from the response in the NextToken value in the next request.
	NextToken *string

	// An array of account IDs.
	MemberAccounts []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListMemberAccountsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListMemberAccounts{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMemberAccounts{}, middleware.After)
}

func newServiceMetadataMiddleware_opListMemberAccounts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "ListMemberAccounts",
	}
}
