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

// Returns the available automatic snapshots for an instance or disk. For more
// information, see the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-configuring-automatic-snapshots).
func (c *Client) GetAutoSnapshots(ctx context.Context, params *GetAutoSnapshotsInput, optFns ...func(*Options)) (*GetAutoSnapshotsOutput, error) {
	stack := middleware.NewStack("GetAutoSnapshots", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetAutoSnapshotsMiddlewares(stack)
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
	addOpGetAutoSnapshotsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAutoSnapshots(options.Region), middleware.Before)
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
			OperationName: "GetAutoSnapshots",
			Err:           err,
		}
	}
	out := result.(*GetAutoSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAutoSnapshotsInput struct {

	// The name of the source instance or disk from which to get automatic snapshot
	// information.
	//
	// This member is required.
	ResourceName *string
}

type GetAutoSnapshotsOutput struct {

	// An array of objects that describe the automatic snapshots that are available for
	// the specified source instance or disk.
	AutoSnapshots []*types.AutoSnapshotDetails

	// The resource type (e.g., Instance or Disk).
	ResourceType types.ResourceType

	// The name of the source instance or disk for the automatic snapshots.
	ResourceName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetAutoSnapshotsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetAutoSnapshots{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAutoSnapshots{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAutoSnapshots(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetAutoSnapshots",
	}
}
