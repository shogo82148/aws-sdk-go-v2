// Code generated by smithy-go-codegen DO NOT EDIT.

package iotevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotevents/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a detector model.
func (c *Client) CreateDetectorModel(ctx context.Context, params *CreateDetectorModelInput, optFns ...func(*Options)) (*CreateDetectorModelOutput, error) {
	stack := middleware.NewStack("CreateDetectorModel", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateDetectorModelMiddlewares(stack)
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
	addOpCreateDetectorModelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDetectorModel(options.Region), middleware.Before)
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
			OperationName: "CreateDetectorModel",
			Err:           err,
		}
	}
	out := result.(*CreateDetectorModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDetectorModelInput struct {

	// Information about the order in which events are evaluated and how actions are
	// executed.
	EvaluationMethod types.EvaluationMethod

	// The input attribute key used to identify a device or system to create a detector
	// (an instance of the detector model) and then to route each input received to the
	// appropriate detector (instance). This parameter uses a JSON-path expression in
	// the message payload of each input to specify the attribute-value pair that is
	// used to identify the device associated with the input.
	Key *string

	// A brief description of the detector model.
	DetectorModelDescription *string

	// Information that defines how the detectors operate.
	//
	// This member is required.
	DetectorModelDefinition *types.DetectorModelDefinition

	// Metadata that can be used to manage the detector model.
	Tags []*types.Tag

	// The ARN of the role that grants permission to AWS IoT Events to perform its
	// operations.
	//
	// This member is required.
	RoleArn *string

	// The name of the detector model.
	//
	// This member is required.
	DetectorModelName *string
}

type CreateDetectorModelOutput struct {

	// Information about how the detector model is configured.
	DetectorModelConfiguration *types.DetectorModelConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateDetectorModelMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateDetectorModel{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDetectorModel{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDetectorModel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotevents",
		OperationName: "CreateDetectorModel",
	}
}
