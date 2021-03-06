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

// Deletes one or more remediation exceptions mentioned in the resource keys. AWS
// Config generates a remediation exception when a problem occurs executing a
// remediation action to a specific resource. Remediation exceptions blocks
// auto-remediation until the exception is cleared.
func (c *Client) DeleteRemediationExceptions(ctx context.Context, params *DeleteRemediationExceptionsInput, optFns ...func(*Options)) (*DeleteRemediationExceptionsOutput, error) {
	stack := middleware.NewStack("DeleteRemediationExceptions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteRemediationExceptionsMiddlewares(stack)
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
	addOpDeleteRemediationExceptionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRemediationExceptions(options.Region), middleware.Before)
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
			OperationName: "DeleteRemediationExceptions",
			Err:           err,
		}
	}
	out := result.(*DeleteRemediationExceptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRemediationExceptionsInput struct {

	// An exception list of resource exception keys to be processed with the current
	// request. AWS Config adds exception for each resource key. For example, AWS
	// Config adds 3 exceptions for 3 resource keys.
	//
	// This member is required.
	ResourceKeys []*types.RemediationExceptionResourceKey

	// The name of the AWS Config rule for which you want to delete remediation
	// exception configuration.
	//
	// This member is required.
	ConfigRuleName *string
}

type DeleteRemediationExceptionsOutput struct {

	// Returns a list of failed delete remediation exceptions batch objects. Each
	// object in the batch consists of a list of failed items and failure messages.
	FailedBatches []*types.FailedDeleteRemediationExceptionsBatch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteRemediationExceptionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteRemediationExceptions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteRemediationExceptions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteRemediationExceptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DeleteRemediationExceptions",
	}
}
