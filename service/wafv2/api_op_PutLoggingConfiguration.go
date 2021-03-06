// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from the
// prior release, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// Enables the specified LoggingConfiguration (), to start logging from a web ACL,
// according to the configuration provided. You can access information about all
// traffic that AWS WAF inspects using the following steps:
//
//     * Create an Amazon
// Kinesis Data Firehose. Create the data firehose with a PUT source and in the
// Region that you are operating. If you are capturing logs for Amazon CloudFront,
// always create the firehose in US East (N. Virginia). Give the data firehose a
// name that starts with the prefix aws-waf-logs-. For example,
// aws-waf-logs-us-east-2-analytics. Do not create the data firehose using a
// Kinesis stream as your source.
//
//     * Associate that firehose to your web ACL
// using a PutLoggingConfiguration request.
//
//     <p>When you successfully enable
// logging using a <code>PutLoggingConfiguration</code> request, AWS WAF will
// create a service linked role with the necessary permissions to write logs to the
// Amazon Kinesis Data Firehose. For more information, see <a
// href="https://docs.aws.amazon.com/waf/latest/developerguide/logging.html">Logging
// Web ACL Traffic Information</a> in the <i>AWS WAF Developer Guide</i>.</p>
func (c *Client) PutLoggingConfiguration(ctx context.Context, params *PutLoggingConfigurationInput, optFns ...func(*Options)) (*PutLoggingConfigurationOutput, error) {
	stack := middleware.NewStack("PutLoggingConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutLoggingConfigurationMiddlewares(stack)
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
	addOpPutLoggingConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutLoggingConfiguration(options.Region), middleware.Before)
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
			OperationName: "PutLoggingConfiguration",
			Err:           err,
		}
	}
	out := result.(*PutLoggingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutLoggingConfigurationInput struct {

	//
	//
	// This member is required.
	LoggingConfiguration *types.LoggingConfiguration
}

type PutLoggingConfigurationOutput struct {

	//
	LoggingConfiguration *types.LoggingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutLoggingConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutLoggingConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutLoggingConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutLoggingConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "PutLoggingConfiguration",
	}
}
