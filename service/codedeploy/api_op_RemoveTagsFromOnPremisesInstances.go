// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes one or more tags from one or more on-premises instances.
func (c *Client) RemoveTagsFromOnPremisesInstances(ctx context.Context, params *RemoveTagsFromOnPremisesInstancesInput, optFns ...func(*Options)) (*RemoveTagsFromOnPremisesInstancesOutput, error) {
	stack := middleware.NewStack("RemoveTagsFromOnPremisesInstances", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRemoveTagsFromOnPremisesInstancesMiddlewares(stack)
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
	addOpRemoveTagsFromOnPremisesInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveTagsFromOnPremisesInstances(options.Region), middleware.Before)
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
			OperationName: "RemoveTagsFromOnPremisesInstances",
			Err:           err,
		}
	}
	out := result.(*RemoveTagsFromOnPremisesInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a RemoveTagsFromOnPremisesInstances operation.
type RemoveTagsFromOnPremisesInstancesInput struct {

	// The tag key-value pairs to remove from the on-premises instances.
	//
	// This member is required.
	Tags []*types.Tag

	// The names of the on-premises instances from which to remove tags.
	//
	// This member is required.
	InstanceNames []*string
}

type RemoveTagsFromOnPremisesInstancesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRemoveTagsFromOnPremisesInstancesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRemoveTagsFromOnPremisesInstances{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRemoveTagsFromOnPremisesInstances{}, middleware.After)
}

func newServiceMetadataMiddleware_opRemoveTagsFromOnPremisesInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "RemoveTagsFromOnPremisesInstances",
	}
}
