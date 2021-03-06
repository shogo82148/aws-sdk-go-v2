// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Gets a VPC link.
func (c *Client) GetVpcLink(ctx context.Context, params *GetVpcLinkInput, optFns ...func(*Options)) (*GetVpcLinkOutput, error) {
	stack := middleware.NewStack("GetVpcLink", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetVpcLinkMiddlewares(stack)
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
	addOpGetVpcLinkValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetVpcLink(options.Region), middleware.Before)
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
			OperationName: "GetVpcLink",
			Err:           err,
		}
	}
	out := result.(*GetVpcLinkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVpcLinkInput struct {

	// The ID of the VPC link.
	//
	// This member is required.
	VpcLinkId *string
}

type GetVpcLinkOutput struct {

	// A list of security group IDs for the VPC link.
	SecurityGroupIds []*string

	// A list of subnet IDs to include in the VPC link.
	SubnetIds []*string

	// The name of the VPC link.
	Name *string

	// A message summarizing the cause of the status of the VPC link.
	VpcLinkStatusMessage *string

	// The ID of the VPC link.
	VpcLinkId *string

	// Tags for the VPC link.
	Tags map[string]*string

	// The status of the VPC link.
	VpcLinkStatus types.VpcLinkStatus

	// The timestamp when the VPC link was created.
	CreatedDate *time.Time

	// The version of the VPC link.
	VpcLinkVersion types.VpcLinkVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetVpcLinkMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetVpcLink{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetVpcLink{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetVpcLink(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetVpcLink",
	}
}
