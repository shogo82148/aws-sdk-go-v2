// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation updates the specified domain contact's privacy setting. When
// privacy protection is enabled, contact information such as email address is
// replaced either with contact information for Amazon Registrar (for .com, .net,
// and .org domains) or with contact information for our registrar associate,
// Gandi. This operation affects only the contact information for the specified
// contact type (registrant, administrator, or tech). If the request succeeds,
// Amazon Route 53 returns an operation ID that you can use with GetOperationDetail
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html)
// to track the progress and completion of the action. If the request doesn't
// complete successfully, the domain registrant will be notified by email. By
// disabling the privacy service via API, you consent to the publication of the
// contact information provided for this domain via the public WHOIS database. You
// certify that you are the registrant of this domain name and have the authority
// to make this decision. You may withdraw your consent at any time by enabling
// privacy protection using either UpdateDomainContactPrivacy or the Route 53
// console. Enabling privacy protection removes the contact information provided
// for this domain from the WHOIS database. For more information on our privacy
// practices, see https://aws.amazon.com/privacy/
// (https://aws.amazon.com/privacy/).
func (c *Client) UpdateDomainContactPrivacy(ctx context.Context, params *UpdateDomainContactPrivacyInput, optFns ...func(*Options)) (*UpdateDomainContactPrivacyOutput, error) {
	stack := middleware.NewStack("UpdateDomainContactPrivacy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateDomainContactPrivacyMiddlewares(stack)
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
	addOpUpdateDomainContactPrivacyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDomainContactPrivacy(options.Region), middleware.Before)
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
			OperationName: "UpdateDomainContactPrivacy",
			Err:           err,
		}
	}
	out := result.(*UpdateDomainContactPrivacyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The UpdateDomainContactPrivacy request includes the following elements.
type UpdateDomainContactPrivacyInput struct {

	// Whether you want to conceal contact information from WHOIS queries. If you
	// specify true, WHOIS ("who is") queries return contact information either for
	// Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If you specify false, WHOIS queries
	// return the information that you entered for the admin contact.
	AdminPrivacy *bool

	// The name of the domain that you want to update the privacy setting for.
	//
	// This member is required.
	DomainName *string

	// Whether you want to conceal contact information from WHOIS queries. If you
	// specify true, WHOIS ("who is") queries return contact information either for
	// Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If you specify false, WHOIS queries
	// return the information that you entered for the technical contact.
	TechPrivacy *bool

	// Whether you want to conceal contact information from WHOIS queries. If you
	// specify true, WHOIS ("who is") queries return contact information either for
	// Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If you specify false, WHOIS queries
	// return the information that you entered for the registrant contact (domain
	// owner).
	RegistrantPrivacy *bool
}

// The UpdateDomainContactPrivacy response includes the following element.
type UpdateDomainContactPrivacyOutput struct {

	// Identifier for tracking the progress of the request. To use this ID to query the
	// operation status, use GetOperationDetail.
	//
	// This member is required.
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateDomainContactPrivacyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDomainContactPrivacy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDomainContactPrivacy{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDomainContactPrivacy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "UpdateDomainContactPrivacy",
	}
}
