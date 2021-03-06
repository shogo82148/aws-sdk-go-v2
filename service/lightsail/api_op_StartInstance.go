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

// Starts a specific Amazon Lightsail instance from a stopped state. To restart an
// instance, use the reboot instance operation. When you start a stopped instance,
// Lightsail assigns a new public IP address to the instance. To use the same IP
// address after stopping and starting an instance, create a static IP address and
// attach it to the instance. For more information, see the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/lightsail-create-static-ip).
// The start instance operation supports tag-based access control via resource tags
// applied to the resource identified by instance name. For more information, see
// the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) StartInstance(ctx context.Context, params *StartInstanceInput, optFns ...func(*Options)) (*StartInstanceOutput, error) {
	stack := middleware.NewStack("StartInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartInstanceMiddlewares(stack)
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
	addOpStartInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartInstance(options.Region), middleware.Before)
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
			OperationName: "StartInstance",
			Err:           err,
		}
	}
	out := result.(*StartInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartInstanceInput struct {

	// The name of the instance (a virtual private server) to start.
	//
	// This member is required.
	InstanceName *string
}

type StartInstanceOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartInstance{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "StartInstance",
	}
}
