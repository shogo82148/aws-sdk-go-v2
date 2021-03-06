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

// Returns temporary SSH keys you can use to connect to a specific virtual private
// server, or instance. The get instance access details operation supports
// tag-based access control via resource tags applied to the resource identified by
// instance name. For more information, see the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) GetInstanceAccessDetails(ctx context.Context, params *GetInstanceAccessDetailsInput, optFns ...func(*Options)) (*GetInstanceAccessDetailsOutput, error) {
	stack := middleware.NewStack("GetInstanceAccessDetails", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetInstanceAccessDetailsMiddlewares(stack)
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
	addOpGetInstanceAccessDetailsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetInstanceAccessDetails(options.Region), middleware.Before)
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
			OperationName: "GetInstanceAccessDetails",
			Err:           err,
		}
	}
	out := result.(*GetInstanceAccessDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInstanceAccessDetailsInput struct {

	// The protocol to use to connect to your instance. Defaults to ssh.
	Protocol types.InstanceAccessProtocol

	// The name of the instance to access.
	//
	// This member is required.
	InstanceName *string
}

type GetInstanceAccessDetailsOutput struct {

	// An array of key-value pairs containing information about a get instance access
	// request.
	AccessDetails *types.InstanceAccessDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetInstanceAccessDetailsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetInstanceAccessDetails{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetInstanceAccessDetails{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetInstanceAccessDetails(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetInstanceAccessDetails",
	}
}
