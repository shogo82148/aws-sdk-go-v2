// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the role associated with a particular group.
func (c *Client) GetAssociatedRole(ctx context.Context, params *GetAssociatedRoleInput, optFns ...func(*Options)) (*GetAssociatedRoleOutput, error) {
	stack := middleware.NewStack("GetAssociatedRole", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetAssociatedRoleMiddlewares(stack)
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
	addOpGetAssociatedRoleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAssociatedRole(options.Region), middleware.Before)
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
			OperationName: "GetAssociatedRole",
			Err:           err,
		}
	}
	out := result.(*GetAssociatedRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAssociatedRoleInput struct {

	// The ID of the Greengrass group.
	//
	// This member is required.
	GroupId *string
}

type GetAssociatedRoleOutput struct {

	// The ARN of the role that is associated with the group.
	RoleArn *string

	// The time when the role was associated with the group.
	AssociatedAt *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetAssociatedRoleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetAssociatedRole{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAssociatedRole{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAssociatedRole(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetAssociatedRole",
	}
}
