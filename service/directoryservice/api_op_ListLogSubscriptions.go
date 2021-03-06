// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the active log subscriptions for the AWS account.
func (c *Client) ListLogSubscriptions(ctx context.Context, params *ListLogSubscriptionsInput, optFns ...func(*Options)) (*ListLogSubscriptionsOutput, error) {
	stack := middleware.NewStack("ListLogSubscriptions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListLogSubscriptionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListLogSubscriptions(options.Region), middleware.Before)
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
			OperationName: "ListLogSubscriptions",
			Err:           err,
		}
	}
	out := result.(*ListLogSubscriptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLogSubscriptionsInput struct {

	// The token for the next set of items to return.
	NextToken *string

	// If a DirectoryID is provided, lists only the log subscription associated with
	// that directory. If no DirectoryId is provided, lists all log subscriptions
	// associated with your AWS account. If there are no log subscriptions for the AWS
	// account or the directory, an empty list will be returned.
	DirectoryId *string

	// The maximum number of items returned.
	Limit *int32
}

type ListLogSubscriptionsOutput struct {

	// The token for the next set of items to return.
	NextToken *string

	// A list of active LogSubscription () objects for calling the AWS account.
	LogSubscriptions []*types.LogSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListLogSubscriptionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListLogSubscriptions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListLogSubscriptions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListLogSubscriptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "ListLogSubscriptions",
	}
}
