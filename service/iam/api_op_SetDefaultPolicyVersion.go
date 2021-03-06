// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the specified version of the specified policy as the policy's default
// (operative) version. This operation affects all users, groups, and roles that
// the policy is attached to. To list the users, groups, and roles that the policy
// is attached to, use the ListEntitiesForPolicy () API. For information about
// managed policies, see Managed Policies and Inline Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
func (c *Client) SetDefaultPolicyVersion(ctx context.Context, params *SetDefaultPolicyVersionInput, optFns ...func(*Options)) (*SetDefaultPolicyVersionOutput, error) {
	stack := middleware.NewStack("SetDefaultPolicyVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpSetDefaultPolicyVersionMiddlewares(stack)
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
	addOpSetDefaultPolicyVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetDefaultPolicyVersion(options.Region), middleware.Before)
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
			OperationName: "SetDefaultPolicyVersion",
			Err:           err,
		}
	}
	out := result.(*SetDefaultPolicyVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetDefaultPolicyVersionInput struct {

	// The Amazon Resource Name (ARN) of the IAM policy whose default version you want
	// to set. For more information about ARNs, see Amazon Resource Names (ARNs) and
	// AWS Service Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	//
	// This member is required.
	PolicyArn *string

	// The version of the policy to set as the default (operative) version. For more
	// information about managed policy versions, see Versioning for Managed Policies
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html)
	// in the IAM User Guide.
	//
	// This member is required.
	VersionId *string
}

type SetDefaultPolicyVersionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpSetDefaultPolicyVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpSetDefaultPolicyVersion{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpSetDefaultPolicyVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetDefaultPolicyVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "SetDefaultPolicyVersion",
	}
}
