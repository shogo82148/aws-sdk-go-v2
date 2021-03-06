// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a constraint. A delegated admin is authorized to invoke this command.
func (c *Client) CreateConstraint(ctx context.Context, params *CreateConstraintInput, optFns ...func(*Options)) (*CreateConstraintOutput, error) {
	stack := middleware.NewStack("CreateConstraint", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateConstraintMiddlewares(stack)
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
	addIdempotencyToken_opCreateConstraintMiddleware(stack, options)
	addOpCreateConstraintValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConstraint(options.Region), middleware.Before)
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
			OperationName: "CreateConstraint",
			Err:           err,
		}
	}
	out := result.(*CreateConstraintOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConstraintInput struct {

	// The constraint parameters, in JSON format. The syntax depends on the constraint
	// type as follows: LAUNCH You are required to specify either the RoleArn or the
	// LocalRoleName but can't use both. Specify the RoleArn property as follows:
	// {"RoleArn" : "arn:aws:iam::123456789012:role/LaunchRole"} Specify the
	// LocalRoleName property as follows: {"LocalRoleName": "SCBasicLaunchRole"} If you
	// specify the LocalRoleName property, when an account uses the launch constraint,
	// the IAM role with that name in the account will be used. This allows launch-role
	// constraints to be account-agnostic so the administrator can create fewer
	// resources per shared account. The given role name must exist in the account used
	// to create the launch constraint and the account of the user who launches a
	// product with this launch constraint. You cannot have both a LAUNCH and a
	// STACKSET constraint. You also cannot have more than one LAUNCH constraint on a
	// product and portfolio. NOTIFICATION Specify the NotificationArns property as
	// follows: {"NotificationArns" : ["arn:aws:sns:us-east-1:123456789012:Topic"]}
	// RESOURCE_UPDATE Specify the TagUpdatesOnProvisionedProduct property as follows:
	// {"Version":"2.0","Properties":{"TagUpdateOnProvisionedProduct":"String"}} The
	// TagUpdatesOnProvisionedProduct property accepts a string value of ALLOWED or
	// NOT_ALLOWED. STACKSET Specify the Parameters property as follows: {"Version":
	// "String", "Properties": {"AccountList": [ "String" ], "RegionList": [ "String"
	// ], "AdminRole": "String", "ExecutionRole": "String"}} You cannot have both a
	// LAUNCH and a STACKSET constraint. You also cannot have more than one STACKSET
	// constraint on a product and portfolio. Products with a STACKSET constraint will
	// launch an AWS CloudFormation stack set. TEMPLATE Specify the Rules property. For
	// more information, see Template Constraint Rules
	// (http://docs.aws.amazon.com/servicecatalog/latest/adminguide/reference-template_constraint_rules.html).
	//
	// This member is required.
	Parameters *string

	// The description of the constraint.
	Description *string

	// The type of constraint.
	//
	//     * LAUNCH
	//
	//     * NOTIFICATION
	//
	//     *
	// RESOURCE_UPDATE
	//
	//     * STACKSET
	//
	//     * TEMPLATE
	//
	// This member is required.
	Type *string

	// A unique identifier that you provide to ensure idempotency. If multiple requests
	// differ only by the idempotency token, the same response is returned for each
	// repeated request.
	//
	// This member is required.
	IdempotencyToken *string

	// The portfolio identifier.
	//
	// This member is required.
	PortfolioId *string

	// The product identifier.
	//
	// This member is required.
	ProductId *string

	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string
}

type CreateConstraintOutput struct {

	// The constraint parameters.
	ConstraintParameters *string

	// Information about the constraint.
	ConstraintDetail *types.ConstraintDetail

	// The status of the current request.
	Status types.Status

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateConstraintMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateConstraint{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateConstraint{}, middleware.After)
}

type idempotencyToken_initializeOpCreateConstraint struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateConstraint) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateConstraint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateConstraintInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateConstraintInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateConstraintMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateConstraint{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateConstraint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "CreateConstraint",
	}
}
