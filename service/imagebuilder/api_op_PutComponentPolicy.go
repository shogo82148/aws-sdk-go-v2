// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Applies a policy to a component. We recommend that you call the RAM API
// CreateResourceShare
// (https://docs.aws.amazon.com/ram/latest/APIReference/API_CreateResourceShare.html)
// to share resources. If you call the Image Builder API PutComponentPolicy, you
// must also call the RAM API PromoteResourceShareCreatedFromPolicy
// (https://docs.aws.amazon.com/ram/latest/APIReference/API_PromoteResourceShareCreatedFromPolicy.html)
// in order for the resource to be visible to all principals with whom the resource
// is shared.
func (c *Client) PutComponentPolicy(ctx context.Context, params *PutComponentPolicyInput, optFns ...func(*Options)) (*PutComponentPolicyOutput, error) {
	stack := middleware.NewStack("PutComponentPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutComponentPolicyMiddlewares(stack)
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
	addOpPutComponentPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutComponentPolicy(options.Region), middleware.Before)
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
			OperationName: "PutComponentPolicy",
			Err:           err,
		}
	}
	out := result.(*PutComponentPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutComponentPolicyInput struct {

	// The policy to apply.
	//
	// This member is required.
	Policy *string

	// The Amazon Resource Name (ARN) of the component that this policy should be
	// applied to.
	//
	// This member is required.
	ComponentArn *string
}

type PutComponentPolicyOutput struct {

	// The request ID that uniquely identifies this request.
	RequestId *string

	// The Amazon Resource Name (ARN) of the component that this policy was applied to.
	ComponentArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutComponentPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutComponentPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutComponentPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutComponentPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "PutComponentPolicy",
	}
}
