// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables the integration of an AWS service (the service that is specified by
// ServicePrincipal) with AWS Organizations. When you enable integration, you allow
// the specified service to create a service-linked role
// (http://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html)
// in all the accounts in your organization. This allows the service to perform
// operations on your behalf in your organization and its accounts. We recommend
// that you enable integration between AWS Organizations and the specified AWS
// service by using the console or commands that are provided by the specified
// service. Doing so ensures that the service is aware that it can create the
// resources that are required for the integration. How the service creates those
// resources in the organization's accounts depends on that service. For more
// information, see the documentation for the other AWS service. For more
// information about enabling services to integrate with AWS Organizations, see
// Integrating AWS Organizations with Other AWS Services
// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html)
// in the AWS Organizations User Guide. This operation can be called only from the
// organization's master account and only if the organization has enabled all
// features
// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html).
func (c *Client) EnableAWSServiceAccess(ctx context.Context, params *EnableAWSServiceAccessInput, optFns ...func(*Options)) (*EnableAWSServiceAccessOutput, error) {
	stack := middleware.NewStack("EnableAWSServiceAccess", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpEnableAWSServiceAccessMiddlewares(stack)
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
	addOpEnableAWSServiceAccessValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableAWSServiceAccess(options.Region), middleware.Before)
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
			OperationName: "EnableAWSServiceAccess",
			Err:           err,
		}
	}
	out := result.(*EnableAWSServiceAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableAWSServiceAccessInput struct {

	// The service principal name of the AWS service for which you want to enable
	// integration with your organization. This is typically in the form of a URL, such
	// as  service-abbreviation.amazonaws.com.
	//
	// This member is required.
	ServicePrincipal *string
}

type EnableAWSServiceAccessOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpEnableAWSServiceAccessMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpEnableAWSServiceAccess{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableAWSServiceAccess{}, middleware.After)
}

func newServiceMetadataMiddleware_opEnableAWSServiceAccess(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "EnableAWSServiceAccess",
	}
}
