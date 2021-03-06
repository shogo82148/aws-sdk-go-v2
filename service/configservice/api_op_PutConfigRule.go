// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds or updates an AWS Config rule for evaluating whether your AWS resources
// comply with your desired configurations. You can use this action for custom AWS
// Config rules and AWS managed Config rules. A custom AWS Config rule is a rule
// that you develop and maintain. An AWS managed Config rule is a customizable,
// predefined rule that AWS Config provides. If you are adding a new custom AWS
// Config rule, you must first create the AWS Lambda function that the rule invokes
// to evaluate your resources. When you use the PutConfigRule action to add the
// rule to AWS Config, you must specify the Amazon Resource Name (ARN) that AWS
// Lambda assigns to the function. Specify the ARN for the SourceIdentifier key.
// This key is part of the Source object, which is part of the ConfigRule object.
// If you are adding an AWS managed Config rule, specify the rule's identifier for
// the SourceIdentifier key. To reference AWS managed Config rule identifiers, see
// About AWS Managed Config Rules
// (https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html).
// For any new rule that you add, specify the ConfigRuleName in the ConfigRule
// object. Do not specify the ConfigRuleArn or the ConfigRuleId. These values are
// generated by AWS Config for new rules. If you are updating a rule that you added
// previously, you can specify the rule by ConfigRuleName, ConfigRuleId, or
// ConfigRuleArn in the ConfigRule data type that you use in this request. The
// maximum number of rules that AWS Config supports is 150.  <p>For information
// about requesting a rule limit increase, see <a
// href="http://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html#limits_config">AWS
// Config Limits</a> in the <i>AWS General Reference Guide</i>.</p> <p>For more
// information about developing and using AWS Config rules, see <a
// href="https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config.html">Evaluating
// AWS Resource Configurations with AWS Config</a> in the <i>AWS Config Developer
// Guide</i>.</p>
func (c *Client) PutConfigRule(ctx context.Context, params *PutConfigRuleInput, optFns ...func(*Options)) (*PutConfigRuleOutput, error) {
	stack := middleware.NewStack("PutConfigRule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutConfigRuleMiddlewares(stack)
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
	addOpPutConfigRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutConfigRule(options.Region), middleware.Before)
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
			OperationName: "PutConfigRule",
			Err:           err,
		}
	}
	out := result.(*PutConfigRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutConfigRuleInput struct {

	// An array of tag object.
	Tags []*types.Tag

	// The rule that you want to add to your account.
	//
	// This member is required.
	ConfigRule *types.ConfigRule
}

type PutConfigRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutConfigRuleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutConfigRule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutConfigRule{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutConfigRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutConfigRule",
	}
}
