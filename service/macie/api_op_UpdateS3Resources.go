// Code generated by smithy-go-codegen DO NOT EDIT.

package macie

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the classification types for the specified S3 resources. If
// memberAccountId isn't specified, the action updates the classification types of
// the S3 resources associated with Amazon Macie Classic for the current master
// account. If memberAccountId is specified, the action updates the classification
// types of the S3 resources associated with Amazon Macie Classic for the specified
// member account.
func (c *Client) UpdateS3Resources(ctx context.Context, params *UpdateS3ResourcesInput, optFns ...func(*Options)) (*UpdateS3ResourcesOutput, error) {
	stack := middleware.NewStack("UpdateS3Resources", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateS3ResourcesMiddlewares(stack)
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
	addOpUpdateS3ResourcesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateS3Resources(options.Region), middleware.Before)
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
			OperationName: "UpdateS3Resources",
			Err:           err,
		}
	}
	out := result.(*UpdateS3ResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateS3ResourcesInput struct {

	// The S3 resources whose classification types you want to update.
	//
	// This member is required.
	S3ResourcesUpdate []*types.S3ResourceClassificationUpdate

	// The AWS ID of the Amazon Macie Classic member account whose S3 resources'
	// classification types you want to update.
	MemberAccountId *string
}

type UpdateS3ResourcesOutput struct {

	// The S3 resources whose classification types can't be updated. An error code and
	// an error message are provided for each failed item.
	FailedS3Resources []*types.FailedS3Resource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateS3ResourcesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateS3Resources{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateS3Resources{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateS3Resources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie",
		OperationName: "UpdateS3Resources",
	}
}
