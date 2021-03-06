// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describe specified users. Required Permissions: To use this action, an IAM user
// must have an attached policy that explicitly grants permissions. For more
// information about user permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) DescribeUserProfiles(ctx context.Context, params *DescribeUserProfilesInput, optFns ...func(*Options)) (*DescribeUserProfilesOutput, error) {
	stack := middleware.NewStack("DescribeUserProfiles", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeUserProfilesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUserProfiles(options.Region), middleware.Before)
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
			OperationName: "DescribeUserProfiles",
			Err:           err,
		}
	}
	out := result.(*DescribeUserProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUserProfilesInput struct {

	// An array of IAM or federated user ARNs that identify the users to be described.
	IamUserArns []*string
}

// Contains the response to a DescribeUserProfiles request.
type DescribeUserProfilesOutput struct {

	// A Users object that describes the specified users.
	UserProfiles []*types.UserProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeUserProfilesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeUserProfiles{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeUserProfiles{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeUserProfiles(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "DescribeUserProfiles",
	}
}
