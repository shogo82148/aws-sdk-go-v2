// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the configuration of a location used by AWS DataSync.
func (c *Client) DeleteLocation(ctx context.Context, params *DeleteLocationInput, optFns ...func(*Options)) (*DeleteLocationOutput, error) {
	stack := middleware.NewStack("DeleteLocation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteLocationMiddlewares(stack)
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
	addOpDeleteLocationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteLocation(options.Region), middleware.Before)
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
			OperationName: "DeleteLocation",
			Err:           err,
		}
	}
	out := result.(*DeleteLocationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DeleteLocation
type DeleteLocationInput struct {

	// The Amazon Resource Name (ARN) of the location to delete.
	//
	// This member is required.
	LocationArn *string
}

type DeleteLocationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteLocationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteLocation{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteLocation{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteLocation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "DeleteLocation",
	}
}
