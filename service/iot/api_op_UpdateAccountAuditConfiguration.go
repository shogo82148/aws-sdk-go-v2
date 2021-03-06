// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Configures or reconfigures the Device Defender audit settings for this account.
// Settings include how audit notifications are sent and which audit checks are
// enabled or disabled.
func (c *Client) UpdateAccountAuditConfiguration(ctx context.Context, params *UpdateAccountAuditConfigurationInput, optFns ...func(*Options)) (*UpdateAccountAuditConfigurationOutput, error) {
	stack := middleware.NewStack("UpdateAccountAuditConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateAccountAuditConfigurationMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAccountAuditConfiguration(options.Region), middleware.Before)
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
			OperationName: "UpdateAccountAuditConfiguration",
			Err:           err,
		}
	}
	out := result.(*UpdateAccountAuditConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAccountAuditConfigurationInput struct {

	// Specifies which audit checks are enabled and disabled for this account. Use
	// DescribeAccountAuditConfiguration to see the list of all checks, including those
	// that are currently enabled. Some data collection might start immediately when
	// certain checks are enabled. When a check is disabled, any data collected so far
	// in relation to the check is deleted. You cannot disable a check if it is used by
	// any scheduled audit. You must first delete the check from the scheduled audit or
	// delete the scheduled audit itself. On the first call to
	// UpdateAccountAuditConfiguration, this parameter is required and must specify at
	// least one enabled check.
	AuditCheckConfigurations map[string]*types.AuditCheckConfiguration

	// Information about the targets to which audit notifications are sent.
	AuditNotificationTargetConfigurations map[string]*types.AuditNotificationTarget

	// The ARN of the role that grants permission to AWS IoT to access information
	// about your devices, policies, certificates and other items as required when
	// performing an audit.
	RoleArn *string
}

type UpdateAccountAuditConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateAccountAuditConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAccountAuditConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAccountAuditConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateAccountAuditConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "UpdateAccountAuditConfiguration",
	}
}
