// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates a device from its current room. The device continues to be
// connected to the Wi-Fi network and is still registered to the account. The
// device settings and skills are removed from the room.
func (c *Client) DisassociateDeviceFromRoom(ctx context.Context, params *DisassociateDeviceFromRoomInput, optFns ...func(*Options)) (*DisassociateDeviceFromRoomOutput, error) {
	stack := middleware.NewStack("DisassociateDeviceFromRoom", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDisassociateDeviceFromRoomMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateDeviceFromRoom(options.Region), middleware.Before)
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
			OperationName: "DisassociateDeviceFromRoom",
			Err:           err,
		}
	}
	out := result.(*DisassociateDeviceFromRoomOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateDeviceFromRoomInput struct {

	// The ARN of the device to disassociate from a room. Required.
	DeviceArn *string
}

type DisassociateDeviceFromRoomOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDisassociateDeviceFromRoomMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateDeviceFromRoom{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateDeviceFromRoom{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateDeviceFromRoom(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "DisassociateDeviceFromRoom",
	}
}
