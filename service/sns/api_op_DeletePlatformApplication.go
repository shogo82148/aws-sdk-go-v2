// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a platform application object for one of the supported push notification
// services, such as APNS and GCM (Firebase Cloud Messaging). For more information,
// see Using Amazon SNS Mobile Push Notifications
// (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html).
func (c *Client) DeletePlatformApplication(ctx context.Context, params *DeletePlatformApplicationInput, optFns ...func(*Options)) (*DeletePlatformApplicationOutput, error) {
	stack := middleware.NewStack("DeletePlatformApplication", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeletePlatformApplicationMiddlewares(stack)
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
	addOpDeletePlatformApplicationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePlatformApplication(options.Region), middleware.Before)
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
			OperationName: "DeletePlatformApplication",
			Err:           err,
		}
	}
	out := result.(*DeletePlatformApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for DeletePlatformApplication action.
type DeletePlatformApplicationInput struct {

	// PlatformApplicationArn of platform application object to delete.
	//
	// This member is required.
	PlatformApplicationArn *string
}

type DeletePlatformApplicationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeletePlatformApplicationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeletePlatformApplication{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeletePlatformApplication{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeletePlatformApplication(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "DeletePlatformApplication",
	}
}
