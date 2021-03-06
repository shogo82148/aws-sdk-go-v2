// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds termination settings for the specified Amazon Chime Voice Connector. If
// emergency calling is configured for the Amazon Chime Voice Connector, it must be
// deleted prior to turning off termination settings.
func (c *Client) PutVoiceConnectorTermination(ctx context.Context, params *PutVoiceConnectorTerminationInput, optFns ...func(*Options)) (*PutVoiceConnectorTerminationOutput, error) {
	stack := middleware.NewStack("PutVoiceConnectorTermination", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutVoiceConnectorTerminationMiddlewares(stack)
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
	addOpPutVoiceConnectorTerminationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutVoiceConnectorTermination(options.Region), middleware.Before)
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
			OperationName: "PutVoiceConnectorTermination",
			Err:           err,
		}
	}
	out := result.(*PutVoiceConnectorTerminationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutVoiceConnectorTerminationInput struct {

	// The Amazon Chime Voice Connector ID.
	//
	// This member is required.
	VoiceConnectorId *string

	// The termination setting details to add.
	//
	// This member is required.
	Termination *types.Termination
}

type PutVoiceConnectorTerminationOutput struct {

	// The updated termination setting details.
	Termination *types.Termination

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutVoiceConnectorTerminationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutVoiceConnectorTermination{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutVoiceConnectorTermination{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutVoiceConnectorTermination(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "PutVoiceConnectorTermination",
	}
}
