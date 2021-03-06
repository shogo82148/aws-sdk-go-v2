// Code generated by smithy-go-codegen DO NOT EDIT.

package iotdataplane

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the shadow for the specified thing. For more information, see
// GetThingShadow
// (http://docs.aws.amazon.com/iot/latest/developerguide/API_GetThingShadow.html)
// in the AWS IoT Developer Guide.
func (c *Client) GetThingShadow(ctx context.Context, params *GetThingShadowInput, optFns ...func(*Options)) (*GetThingShadowOutput, error) {
	stack := middleware.NewStack("GetThingShadow", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetThingShadowMiddlewares(stack)
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
	addOpGetThingShadowValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetThingShadow(options.Region), middleware.Before)
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
			OperationName: "GetThingShadow",
			Err:           err,
		}
	}
	out := result.(*GetThingShadowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the GetThingShadow operation.
type GetThingShadowInput struct {

	// The name of the thing.
	//
	// This member is required.
	ThingName *string

	// The name of the shadow.
	ShadowName *string
}

// The output from the GetThingShadow operation.
type GetThingShadowOutput struct {

	// The state information, in JSON format.
	Payload []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetThingShadowMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetThingShadow{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetThingShadow{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetThingShadow(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotdata",
		OperationName: "GetThingShadow",
	}
}
