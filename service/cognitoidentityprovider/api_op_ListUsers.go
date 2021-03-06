// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the users in the Amazon Cognito user pool.
func (c *Client) ListUsers(ctx context.Context, params *ListUsersInput, optFns ...func(*Options)) (*ListUsersOutput, error) {
	stack := middleware.NewStack("ListUsers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListUsersMiddlewares(stack)
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
	addOpListUsersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListUsers(options.Region), middleware.Before)
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
			OperationName: "ListUsers",
			Err:           err,
		}
	}
	out := result.(*ListUsersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to list users.
type ListUsersInput struct {

	// An array of strings, where each string is the name of a user attribute to be
	// returned for each user in the search results. If the array is null, all
	// attributes are returned.
	AttributesToGet []*string

	// The user pool ID for the user pool on which the search should be performed.
	//
	// This member is required.
	UserPoolId *string

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	PaginationToken *string

	// Maximum number of users to be returned.
	Limit *int32

	// A filter string of the form "AttributeName Filter-Type "AttributeValue"".
	// Quotation marks within the filter string must be escaped using the backslash (\)
	// character. For example, "family_name = \"Reddy\"".
	//
	//     * AttributeName: The
	// name of the attribute to search for. You can only search for one attribute at a
	// time.
	//
	//     * Filter-Type: For an exact match, use =, for example, "given_name =
	// \"Jon\"". For a prefix ("starts with") match, use ^=, for example, "given_name
	// ^= \"Jon\"".
	//
	//     * AttributeValue: The attribute value that must be matched for
	// each user.
	//
	// If the filter string is empty, ListUsers returns all users in the
	// user pool. You can only search for the following standard attributes:
	//
	//     *
	// username (case-sensitive)
	//
	//     * email
	//
	//     * phone_number
	//
	//     * name
	//
	//     *
	// given_name
	//
	//     * family_name
	//
	//     * preferred_username
	//
	//     *
	// cognito:user_status (called Status in the Console) (case-insensitive)
	//
	//     *
	// status (called Enabled in the Console) (case-sensitive)
	//
	//     * sub
	//
	// Custom
	// attributes are not searchable. For more information, see Searching for Users
	// Using the ListUsers API
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/how-to-manage-user-accounts.html#cognito-user-pools-searching-for-users-using-listusers-api)
	// and Examples of Using the ListUsers API
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/how-to-manage-user-accounts.html#cognito-user-pools-searching-for-users-listusers-api-examples)
	// in the Amazon Cognito Developer Guide.
	Filter *string
}

// The response from the request to list users.
type ListUsersOutput struct {

	// The users returned in the request to list users.
	Users []*types.UserType

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	PaginationToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListUsersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListUsers{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListUsers{}, middleware.After)
}

func newServiceMetadataMiddleware_opListUsers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "ListUsers",
	}
}
