// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Connects a volume to an iSCSI connection and then attaches the volume to the
// specified gateway. Detaching and attaching a volume enables you to recover your
// data from one gateway to a different gateway without creating a snapshot. It
// also makes it easier to move your volumes from an on-premises gateway to a
// gateway hosted on an Amazon EC2 instance.
func (c *Client) AttachVolume(ctx context.Context, params *AttachVolumeInput, optFns ...func(*Options)) (*AttachVolumeOutput, error) {
	stack := middleware.NewStack("AttachVolume", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAttachVolumeMiddlewares(stack)
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
	addOpAttachVolumeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAttachVolume(options.Region), middleware.Before)
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
			OperationName: "AttachVolume",
			Err:           err,
		}
	}
	out := result.(*AttachVolumeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// AttachVolumeInput
type AttachVolumeInput struct {

	// The Amazon Resource Name (ARN) of the volume to attach to the specified gateway.
	//
	// This member is required.
	VolumeARN *string

	// The name of the iSCSI target used by an initiator to connect to a volume and
	// used as a suffix for the target ARN. For example, specifying TargetName as
	// myvolume results in the target ARN of
	// arn:aws:storagegateway:us-east-2:111122223333:gateway/sgw-12A3456B/target/iqn.1997-05.com.amazon:myvolume.
	// The target name must be unique across all volumes on a gateway.  <p>If you don't
	// specify a value, Storage Gateway uses the value that was previously used for
	// this volume as the new target name.</p>
	TargetName *string

	// The Amazon Resource Name (ARN) of the gateway that you want to attach the volume
	// to.
	//
	// This member is required.
	GatewayARN *string

	// The unique device ID or other distinguishing data that identifies the local disk
	// used to create the volume. This value is only required when you are attaching a
	// stored volume.
	DiskId *string

	// The network interface of the gateway on which to expose the iSCSI target. Only
	// IPv4 addresses are accepted. Use DescribeGatewayInformation () to get a list of
	// the network interfaces available on a gateway.  <p>Valid Values: A valid IP
	// address.</p>
	//
	// This member is required.
	NetworkInterfaceId *string
}

// AttachVolumeOutput
type AttachVolumeOutput struct {

	// The Amazon Resource Name (ARN) of the volume that was attached to the gateway.
	VolumeARN *string

	// The Amazon Resource Name (ARN) of the volume target, which includes the iSCSI
	// name for the initiator that was used to connect to the target.
	TargetARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAttachVolumeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAttachVolume{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAttachVolume{}, middleware.After)
}

func newServiceMetadataMiddleware_opAttachVolume(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "AttachVolume",
	}
}
