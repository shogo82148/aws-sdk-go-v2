// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of security standards controls. For each control, the results
// include information about whether it is currently enabled, the severity, and a
// link to remediation information.
func (c *Client) DescribeStandardsControls(ctx context.Context, params *DescribeStandardsControlsInput, optFns ...func(*Options)) (*DescribeStandardsControlsOutput, error) {
	stack := middleware.NewStack("DescribeStandardsControls", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeStandardsControlsMiddlewares(stack)
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
	addOpDescribeStandardsControlsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStandardsControls(options.Region), middleware.Before)
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
			OperationName: "DescribeStandardsControls",
			Err:           err,
		}
	}
	out := result.(*DescribeStandardsControlsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStandardsControlsInput struct {

	// The maximum number of security standard controls to return.
	MaxResults *int32

	// The ARN of a resource that represents your subscription to a supported standard.
	//
	// This member is required.
	StandardsSubscriptionArn *string

	// The token that is required for pagination. On your first call to the
	// DescribeStandardsControls operation, set the value of this parameter to NULL.
	// For subsequent calls to the operation, to continue listing data, set the value
	// of this parameter to the value returned from the previous response.
	NextToken *string
}

type DescribeStandardsControlsOutput struct {

	// A list of security standards controls.
	Controls []*types.StandardsControl

	// The pagination token to use to request the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeStandardsControlsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeStandardsControls{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeStandardsControls{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeStandardsControls(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "DescribeStandardsControls",
	}
}
