// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about the organization that the user's account belongs to.
// This operation can be called from any account in the organization. Even if a
// policy type is shown as available in the organization, you can disable it
// separately at the root level with DisablePolicyType (). Use ListRoots () to see
// the status of policy types for a specified root.
func (c *Client) DescribeOrganization(ctx context.Context, params *DescribeOrganizationInput, optFns ...func(*Options)) (*DescribeOrganizationOutput, error) {
	stack := middleware.NewStack("DescribeOrganization", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeOrganizationMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrganization(options.Region), middleware.Before)
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
			OperationName: "DescribeOrganization",
			Err:           err,
		}
	}
	out := result.(*DescribeOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrganizationInput struct {
}

type DescribeOrganizationOutput struct {

	// A structure that contains information about the organization. The
	// AvailablePolicyTypes part of the response is deprecated, and you shouldn't use
	// it in your apps. It doesn't include any policy type supported by Organizations
	// other than SCPs. To determine which policy types are enabled in your
	// organization, use the ListRoots () operation.
	Organization *types.Organization

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeOrganizationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeOrganization{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeOrganization{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeOrganization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DescribeOrganization",
	}
}
