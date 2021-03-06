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

// Disables the integration of an AWS service (the service that is specified by
// ServicePrincipal) with AWS Organizations. When you disable integration, the
// specified service no longer can create a service-linked role
// (http://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html)
// in new accounts in your organization. This means the service can't perform
// operations on your behalf on any new accounts in your organization. The service
// can still perform operations in older accounts until the service completes its
// clean-up from AWS Organizations. We recommend that you disable integration
// between AWS Organizations and the specified AWS service by using the console or
// commands that are provided by the specified service. Doing so ensures that the
// other service is aware that it can clean up any resources that are required only
// for the integration. How the service cleans up its resources in the
// organization's accounts depends on that service. For more information, see the
// documentation for the other AWS service. After you perform the
// DisableAWSServiceAccess operation, the specified service can no longer perform
// operations in your organization's accounts unless the operations are explicitly
// permitted by the IAM policies that are attached to your roles. For more
// information about integrating other services with AWS Organizations, including
// the list of services that work with Organizations, see Integrating AWS
// Organizations with Other AWS Services
// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html)
// in the AWS Organizations User Guide. This operation can be called only from the
// organization's master account.
func (c *Client) DisableAWSServiceAccess(ctx context.Context, params *DisableAWSServiceAccessInput, optFns ...func(*Options)) (*DisableAWSServiceAccessOutput, error) {
	stack := middleware.NewStack("DisableAWSServiceAccess", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDisableAWSServiceAccessMiddlewares(stack)
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
	addOpDisableAWSServiceAccessValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableAWSServiceAccess(options.Region), middleware.Before)
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
			OperationName: "DisableAWSServiceAccess",
			Err:           err,
		}
	}
	out := result.(*DisableAWSServiceAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableAWSServiceAccessInput struct {

	// The service principal name of the AWS service for which you want to disable
	// integration with your organization. This is typically in the form of a URL, such
	// as  service-abbreviation.amazonaws.com.
	//
	// This member is required.
	ServicePrincipal *string
}

type DisableAWSServiceAccessOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDisableAWSServiceAccessMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDisableAWSServiceAccess{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisableAWSServiceAccess{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisableAWSServiceAccess(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DisableAWSServiceAccess",
	}
}
