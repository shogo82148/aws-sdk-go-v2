// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an array of Challenge-Handshake Authentication Protocol (CHAP)
// credentials information for a specified iSCSI target, one for each
// target-initiator pair. This operation is supported in the volume and tape
// gateway types.
func (c *Client) DescribeChapCredentials(ctx context.Context, params *DescribeChapCredentialsInput, optFns ...func(*Options)) (*DescribeChapCredentialsOutput, error) {
	stack := middleware.NewStack("DescribeChapCredentials", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeChapCredentialsMiddlewares(stack)
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
	addOpDescribeChapCredentialsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeChapCredentials(options.Region), middleware.Before)
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
			OperationName: "DescribeChapCredentials",
			Err:           err,
		}
	}
	out := result.(*DescribeChapCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing the Amazon Resource Name (ARN) of the iSCSI volume
// target.
type DescribeChapCredentialsInput struct {

	// The Amazon Resource Name (ARN) of the iSCSI volume target. Use the
	// DescribeStorediSCSIVolumes () operation to return to retrieve the TargetARN for
	// specified VolumeARN.
	//
	// This member is required.
	TargetARN *string
}

// A JSON object containing the following fields:
type DescribeChapCredentialsOutput struct {

	// An array of ChapInfo () objects that represent CHAP credentials. Each object in
	// the array contains CHAP credential information for one target-initiator pair. If
	// no CHAP credentials are set, an empty array is returned. CHAP credential
	// information is provided in a JSON object with the following fields:  <ul> <li>
	// <p> <b>InitiatorName</b>: The iSCSI initiator that connects to the target.</p>
	// </li> <li> <p> <b>SecretToAuthenticateInitiator</b>: The secret key that the
	// initiator (for example, the Windows client) must provide to participate in
	// mutual CHAP with the target.</p> </li> <li> <p>
	// <b>SecretToAuthenticateTarget</b>: The secret key that the target must provide
	// to participate in mutual CHAP with the initiator (e.g. Windows client).</p>
	// </li> <li> <p> <b>TargetARN</b>: The Amazon Resource Name (ARN) of the storage
	// volume.</p> </li> </ul>
	ChapCredentials []*types.ChapInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeChapCredentialsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeChapCredentials{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeChapCredentials{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeChapCredentials(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeChapCredentials",
	}
}
