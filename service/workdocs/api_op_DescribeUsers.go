// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified users. You can describe all users or filter the results
// (for example, by status or organization). By default, Amazon WorkDocs returns
// the first 24 active or pending users. If there are more results, the response
// includes a marker that you can use to request the next set of results.
func (c *Client) DescribeUsers(ctx context.Context, params *DescribeUsersInput, optFns ...func(*Options)) (*DescribeUsersOutput, error) {
	stack := middleware.NewStack("DescribeUsers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeUsersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUsers(options.Region), middleware.Before)
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
			OperationName: "DescribeUsers",
			Err:           err,
		}
	}
	out := result.(*DescribeUsersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUsersInput struct {

	// A comma-separated list of values. Specify "STORAGE_METADATA" to include the user
	// storage quota and utilization information.
	Fields *string

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string

	// The order for the results.
	Order types.OrderType

	// The maximum number of items to return.
	Limit *int32

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string

	// The state of the users. Specify "ALL" to include inactive users.
	Include types.UserFilterType

	// The sorting criteria.
	Sort types.UserSortType

	// The IDs of the users.
	UserIds *string

	// A query to filter users by user name.
	Query *string

	// The ID of the organization.
	OrganizationId *string
}

type DescribeUsersOutput struct {

	// The total number of users included in the results.
	TotalNumberOfUsers *int64

	// The marker to use when requesting the next set of results. If there are no
	// additional results, the string is empty.
	Marker *string

	// The users.
	Users []*types.User

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeUsersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeUsers{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeUsers{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeUsers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "DescribeUsers",
	}
}
