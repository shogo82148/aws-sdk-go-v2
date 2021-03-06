// Code generated by smithy-go-codegen DO NOT EDIT.

package accessanalyzer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an archive rule for the specified analyzer. Archive rules automatically
// archive findings that meet the criteria you define when you create the rule.
func (c *Client) CreateArchiveRule(ctx context.Context, params *CreateArchiveRuleInput, optFns ...func(*Options)) (*CreateArchiveRuleOutput, error) {
	stack := middleware.NewStack("CreateArchiveRule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateArchiveRuleMiddlewares(stack)
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
	addOpCreateArchiveRuleValidationMiddleware(stack)
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
			OperationName: "CreateArchiveRule",
			Err:           err,
		}
	}
	out := result.(*CreateArchiveRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates an archive rule.
type CreateArchiveRuleInput struct {

	// The name of the created analyzer.
	//
	// This member is required.
	AnalyzerName *string

	// The name of the rule to create.
	//
	// This member is required.
	RuleName *string

	// The criteria for the rule.
	//
	// This member is required.
	Filter map[string]*types.Criterion

	// A client token.
	ClientToken *string
}

type CreateArchiveRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateArchiveRuleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateArchiveRule{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateArchiveRule{}, middleware.After)
}
