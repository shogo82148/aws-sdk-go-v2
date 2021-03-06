// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates engine-specific attributes on a specified server. The server enters the
// MODIFYING state when this operation is in progress. Only one update can occur at
// a time. You can use this command to reset a Chef server's public key
// (CHEF_PIVOTAL_KEY) or a Puppet server's admin password (PUPPET_ADMIN_PASSWORD).
// This operation is asynchronous. This operation can only be called for servers in
// HEALTHY or UNHEALTHY states. Otherwise, an InvalidStateException is raised. A
// ResourceNotFoundException is thrown when the server does not exist. A
// ValidationException is raised when parameters of the request are not valid.
func (c *Client) UpdateServerEngineAttributes(ctx context.Context, params *UpdateServerEngineAttributesInput, optFns ...func(*Options)) (*UpdateServerEngineAttributesOutput, error) {
	stack := middleware.NewStack("UpdateServerEngineAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateServerEngineAttributesMiddlewares(stack)
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
	addOpUpdateServerEngineAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateServerEngineAttributes(options.Region), middleware.Before)
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
			OperationName: "UpdateServerEngineAttributes",
			Err:           err,
		}
	}
	out := result.(*UpdateServerEngineAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateServerEngineAttributesInput struct {

	// The name of the engine attribute to update.
	//
	// This member is required.
	AttributeName *string

	// The name of the server to update.
	//
	// This member is required.
	ServerName *string

	// The value to set for the attribute.
	AttributeValue *string
}

type UpdateServerEngineAttributesOutput struct {

	// Contains the response to an UpdateServerEngineAttributes request.
	Server *types.Server

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateServerEngineAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateServerEngineAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateServerEngineAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateServerEngineAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks-cm",
		OperationName: "UpdateServerEngineAttributes",
	}
}
