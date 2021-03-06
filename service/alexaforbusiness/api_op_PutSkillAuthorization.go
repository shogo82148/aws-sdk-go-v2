// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Links a user's account to a third-party skill provider. If this API operation is
// called by an assumed IAM role, the skill being linked must be a private skill.
// Also, the skill must be owned by the AWS account that assumed the IAM role.
func (c *Client) PutSkillAuthorization(ctx context.Context, params *PutSkillAuthorizationInput, optFns ...func(*Options)) (*PutSkillAuthorizationOutput, error) {
	stack := middleware.NewStack("PutSkillAuthorization", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutSkillAuthorizationMiddlewares(stack)
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
	addOpPutSkillAuthorizationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutSkillAuthorization(options.Region), middleware.Before)
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
			OperationName: "PutSkillAuthorization",
			Err:           err,
		}
	}
	out := result.(*PutSkillAuthorizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutSkillAuthorizationInput struct {

	// The unique identifier of a skill.
	//
	// This member is required.
	SkillId *string

	// The room that the skill is authorized for.
	RoomArn *string

	// The authorization result specific to OAUTH code grant output. "Code” must be
	// populated in the AuthorizationResult map to establish the authorization.
	//
	// This member is required.
	AuthorizationResult map[string]*string
}

type PutSkillAuthorizationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutSkillAuthorizationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutSkillAuthorization{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutSkillAuthorization{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutSkillAuthorization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "PutSkillAuthorization",
	}
}
