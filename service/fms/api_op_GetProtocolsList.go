// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about the specified AWS Firewall Manager protocols list.
func (c *Client) GetProtocolsList(ctx context.Context, params *GetProtocolsListInput, optFns ...func(*Options)) (*GetProtocolsListOutput, error) {
	stack := middleware.NewStack("GetProtocolsList", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetProtocolsListMiddlewares(stack)
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
	addOpGetProtocolsListValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetProtocolsList(options.Region), middleware.Before)
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
			OperationName: "GetProtocolsList",
			Err:           err,
		}
	}
	out := result.(*GetProtocolsListOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetProtocolsListInput struct {

	// The ID of the AWS Firewall Manager protocols list that you want the details for.
	//
	// This member is required.
	ListId *string

	// Specifies whether the list to retrieve is a default list owned by AWS Firewall
	// Manager.
	DefaultList *bool
}

type GetProtocolsListOutput struct {

	// Information about the specified AWS Firewall Manager protocols list.
	ProtocolsList *types.ProtocolsListData

	// The Amazon Resource Name (ARN) of the specified protocols list.
	ProtocolsListArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetProtocolsListMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetProtocolsList{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetProtocolsList{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetProtocolsList(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "GetProtocolsList",
	}
}
