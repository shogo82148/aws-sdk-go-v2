// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the known host key or certificate used by the Amazon Lightsail
// browser-based SSH or RDP clients to authenticate an instance. This operation
// enables the Lightsail browser-based SSH or RDP clients to connect to the
// instance after a host key mismatch. Perform this operation only if you were
// expecting the host key or certificate mismatch or if you are familiar with the
// new host key or certificate on the instance. For more information, see
// Troubleshooting connection issues when using the Amazon Lightsail browser-based
// SSH or RDP client
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-troubleshooting-browser-based-ssh-rdp-client-connection).
func (c *Client) DeleteKnownHostKeys(ctx context.Context, params *DeleteKnownHostKeysInput, optFns ...func(*Options)) (*DeleteKnownHostKeysOutput, error) {
	stack := middleware.NewStack("DeleteKnownHostKeys", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteKnownHostKeysMiddlewares(stack)
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
	addOpDeleteKnownHostKeysValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteKnownHostKeys(options.Region), middleware.Before)
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
			OperationName: "DeleteKnownHostKeys",
			Err:           err,
		}
	}
	out := result.(*DeleteKnownHostKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteKnownHostKeysInput struct {

	// The name of the instance for which you want to reset the host key or
	// certificate.
	//
	// This member is required.
	InstanceName *string
}

type DeleteKnownHostKeysOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteKnownHostKeysMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteKnownHostKeys{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteKnownHostKeys{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteKnownHostKeys(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "DeleteKnownHostKeys",
	}
}
