// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes one or more versions of a launch template. You cannot delete the default
// version of a launch template; you must first assign a different version as the
// default. If the default version is the only version for the launch template, you
// must delete the entire launch template using DeleteLaunchTemplate ().
func (c *Client) DeleteLaunchTemplateVersions(ctx context.Context, params *DeleteLaunchTemplateVersionsInput, optFns ...func(*Options)) (*DeleteLaunchTemplateVersionsOutput, error) {
	stack := middleware.NewStack("DeleteLaunchTemplateVersions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDeleteLaunchTemplateVersionsMiddlewares(stack)
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
	addOpDeleteLaunchTemplateVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteLaunchTemplateVersions(options.Region), middleware.Before)
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
			OperationName: "DeleteLaunchTemplateVersions",
			Err:           err,
		}
	}
	out := result.(*DeleteLaunchTemplateVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteLaunchTemplateVersionsInput struct {

	// The ID of the launch template. You must specify either the launch template ID or
	// launch template name in the request.
	LaunchTemplateId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The name of the launch template. You must specify either the launch template ID
	// or launch template name in the request.
	LaunchTemplateName *string

	// The version numbers of one or more launch template versions to delete.
	//
	// This member is required.
	Versions []*string
}

type DeleteLaunchTemplateVersionsOutput struct {

	// Information about the launch template versions that could not be deleted.
	UnsuccessfullyDeletedLaunchTemplateVersions []*types.DeleteLaunchTemplateVersionsResponseErrorItem

	// Information about the launch template versions that were successfully deleted.
	SuccessfullyDeletedLaunchTemplateVersions []*types.DeleteLaunchTemplateVersionsResponseSuccessItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDeleteLaunchTemplateVersionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDeleteLaunchTemplateVersions{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteLaunchTemplateVersions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteLaunchTemplateVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteLaunchTemplateVersions",
	}
}
