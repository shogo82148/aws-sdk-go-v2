// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Delete a program from a multiplex.
func (c *Client) DeleteMultiplexProgram(ctx context.Context, params *DeleteMultiplexProgramInput, optFns ...func(*Options)) (*DeleteMultiplexProgramOutput, error) {
	stack := middleware.NewStack("DeleteMultiplexProgram", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteMultiplexProgramMiddlewares(stack)
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
	addOpDeleteMultiplexProgramValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMultiplexProgram(options.Region), middleware.Before)
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
			OperationName: "DeleteMultiplexProgram",
			Err:           err,
		}
	}
	out := result.(*DeleteMultiplexProgramOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DeleteMultiplexProgramRequest
type DeleteMultiplexProgramInput struct {

	// The multiplex program name.
	//
	// This member is required.
	ProgramName *string

	// The ID of the multiplex that the program belongs to.
	//
	// This member is required.
	MultiplexId *string
}

// Placeholder documentation for DeleteMultiplexProgramResponse
type DeleteMultiplexProgramOutput struct {

	// The settings for this multiplex program.
	MultiplexProgramSettings *types.MultiplexProgramSettings

	// The MediaLive channel associated with the program.
	ChannelId *string

	// The name of the multiplex program.
	ProgramName *string

	// The packet identifier map for this multiplex program.
	PacketIdentifiersMap *types.MultiplexProgramPacketIdentifiersMap

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteMultiplexProgramMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteMultiplexProgram{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteMultiplexProgram{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteMultiplexProgram(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "DeleteMultiplexProgram",
	}
}
