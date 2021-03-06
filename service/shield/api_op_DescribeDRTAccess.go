// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the current role and list of Amazon S3 log buckets used by the DDoS
// Response Team (DRT) to access your AWS account while assisting with attack
// mitigation.
func (c *Client) DescribeDRTAccess(ctx context.Context, params *DescribeDRTAccessInput, optFns ...func(*Options)) (*DescribeDRTAccessOutput, error) {
	stack := middleware.NewStack("DescribeDRTAccess", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeDRTAccessMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDRTAccess(options.Region), middleware.Before)
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
			OperationName: "DescribeDRTAccess",
			Err:           err,
		}
	}
	out := result.(*DescribeDRTAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDRTAccessInput struct {
}

type DescribeDRTAccessOutput struct {

	// The Amazon Resource Name (ARN) of the role the DRT used to access your AWS
	// account.
	RoleArn *string

	// The list of Amazon S3 buckets accessed by the DRT.
	LogBucketList []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeDRTAccessMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDRTAccess{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDRTAccess{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDRTAccess(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "DescribeDRTAccess",
	}
}
