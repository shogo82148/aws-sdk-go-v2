// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified version from the specified application. You cannot delete
// an application version that is associated with a running environment.
func (c *Client) DeleteApplicationVersion(ctx context.Context, params *DeleteApplicationVersionInput, optFns ...func(*Options)) (*DeleteApplicationVersionOutput, error) {
	stack := middleware.NewStack("DeleteApplicationVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteApplicationVersionMiddlewares(stack)
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
	addOpDeleteApplicationVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteApplicationVersion(options.Region), middleware.Before)
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
			OperationName: "DeleteApplicationVersion",
			Err:           err,
		}
	}
	out := result.(*DeleteApplicationVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to delete an application version.
type DeleteApplicationVersionInput struct {

	// Set to true to delete the source bundle from your storage bucket. Otherwise, the
	// application version is deleted only from Elastic Beanstalk and the source bundle
	// remains in Amazon S3.
	DeleteSourceBundle *bool

	// The name of the application to which the version belongs.
	//
	// This member is required.
	ApplicationName *string

	// The label of the version to delete.
	//
	// This member is required.
	VersionLabel *string
}

type DeleteApplicationVersionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteApplicationVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteApplicationVersion{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteApplicationVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteApplicationVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DeleteApplicationVersion",
	}
}
