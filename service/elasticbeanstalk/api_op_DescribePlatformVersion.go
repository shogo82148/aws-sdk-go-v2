// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes a platform version. Provides full details. Compare to
// ListPlatformVersions (), which provides summary information about a list of
// platform versions. For definitions of platform version and other
// platform-related terms, see AWS Elastic Beanstalk Platforms Glossary
// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/platforms-glossary.html).
func (c *Client) DescribePlatformVersion(ctx context.Context, params *DescribePlatformVersionInput, optFns ...func(*Options)) (*DescribePlatformVersionOutput, error) {
	stack := middleware.NewStack("DescribePlatformVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribePlatformVersionMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePlatformVersion(options.Region), middleware.Before)
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
			OperationName: "DescribePlatformVersion",
			Err:           err,
		}
	}
	out := result.(*DescribePlatformVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePlatformVersionInput struct {

	// The ARN of the platform version.
	PlatformArn *string
}

type DescribePlatformVersionOutput struct {

	// Detailed information about the platform version.
	PlatformDescription *types.PlatformDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribePlatformVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribePlatformVersion{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribePlatformVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribePlatformVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DescribePlatformVersion",
	}
}
