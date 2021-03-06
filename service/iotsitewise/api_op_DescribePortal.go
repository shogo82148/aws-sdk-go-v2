// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves information about a portal.
func (c *Client) DescribePortal(ctx context.Context, params *DescribePortalInput, optFns ...func(*Options)) (*DescribePortalOutput, error) {
	stack := middleware.NewStack("DescribePortal", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribePortalMiddlewares(stack)
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
	addOpDescribePortalValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePortal(options.Region), middleware.Before)
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
			OperationName: "DescribePortal",
			Err:           err,
		}
	}
	out := result.(*DescribePortalOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePortalInput struct {

	// The ID of the portal.
	//
	// This member is required.
	PortalId *string
}

type DescribePortalOutput struct {

	// The date the portal was created, in Unix epoch time.
	//
	// This member is required.
	PortalCreationDate *time.Time

	// The date the portal was last updated, in Unix epoch time.
	//
	// This member is required.
	PortalLastUpdateDate *time.Time

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the portal, which has the following format.
	// arn:${Partition}:iotsitewise:${Region}:${Account}:portal/${PortalId}
	//
	// This member is required.
	PortalArn *string

	// The portal's logo image, which is available at a URL.
	PortalLogoImageLocation *types.ImageLocation

	// The AWS SSO application generated client ID (used with AWS SSO APIs).
	//
	// This member is required.
	PortalClientId *string

	// The AWS administrator's contact email address.
	//
	// This member is required.
	PortalContactEmail *string

	// The ID of the portal.
	//
	// This member is required.
	PortalId *string

	// The public root URL for the AWS IoT AWS IoT SiteWise Monitor application portal.
	//
	// This member is required.
	PortalStartUrl *string

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the service role that allows the portal's users to access your AWS IoT SiteWise
	// resources on your behalf. For more information, see Using service roles for AWS
	// IoT SiteWise Monitor
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/monitor-service-role.html)
	// in the AWS IoT SiteWise User Guide.
	RoleArn *string

	// The name of the portal.
	//
	// This member is required.
	PortalName *string

	// The current status of the portal, which contains a state and any error message.
	//
	// This member is required.
	PortalStatus *types.PortalStatus

	// The portal's description.
	PortalDescription *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribePortalMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribePortal{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribePortal{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribePortal(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "DescribePortal",
	}
}
