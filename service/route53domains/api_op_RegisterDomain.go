// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation registers a domain. Domains are registered either by Amazon
// Registrar (for .com, .net, and .org domains) or by our registrar associate,
// Gandi (for all other domains). For some top-level domains (TLDs), this operation
// requires extra parameters. When you register a domain, Amazon Route 53 does the
// following:
//
//     * Creates a Route 53 hosted zone that has the same name as the
// domain. Route 53 assigns four name servers to your hosted zone and automatically
// updates your domain registration with the names of these name servers.
//
//     *
// Enables autorenew, so your domain registration will renew automatically each
// year. We'll notify you in advance of the renewal date so you can choose whether
// to renew the registration.
//
//     * Optionally enables privacy protection, so
// WHOIS queries return contact information either for Amazon Registrar (for .com,
// .net, and .org domains) or for our registrar associate, Gandi (for all other
// TLDs). If you don't enable privacy protection, WHOIS queries return the
// information that you entered for the registrant, admin, and tech contacts.
//
//
// * If registration is successful, returns an operation ID that you can use to
// track the progress and completion of the action. If the request is not completed
// successfully, the domain registrant is notified by email.
//
//     * Charges your
// AWS account an amount based on the top-level domain. For more information, see
// Amazon Route 53 Pricing (http://aws.amazon.com/route53/pricing/).
func (c *Client) RegisterDomain(ctx context.Context, params *RegisterDomainInput, optFns ...func(*Options)) (*RegisterDomainOutput, error) {
	stack := middleware.NewStack("RegisterDomain", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRegisterDomainMiddlewares(stack)
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
	addOpRegisterDomainValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterDomain(options.Region), middleware.Before)
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
			OperationName: "RegisterDomain",
			Err:           err,
		}
	}
	out := result.(*RegisterDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The RegisterDomain request includes the following elements.
type RegisterDomainInput struct {

	// The number of years that you want to register the domain for. Domains are
	// registered for a minimum of one year. The maximum period depends on the
	// top-level domain. For the range of valid values for your domain, see Domains
	// that You Can Register with Amazon Route 53
	// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/registrar-tld-list.html)
	// in the Amazon Route 53 Developer Guide. Default: 1
	//
	// This member is required.
	DurationInYears *int32

	// Whether you want to conceal contact information from WHOIS queries. If you
	// specify true, WHOIS ("who is") queries return contact information either for
	// Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If you specify false, WHOIS queries
	// return the information that you entered for the technical contact. Default: true
	PrivacyProtectTechContact *bool

	// Whether you want to conceal contact information from WHOIS queries. If you
	// specify true, WHOIS ("who is") queries return contact information either for
	// Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If you specify false, WHOIS queries
	// return the information that you entered for the registrant contact (the domain
	// owner). Default: true
	PrivacyProtectRegistrantContact *bool

	// Indicates whether the domain will be automatically renewed (true) or not
	// (false). Autorenewal only takes effect after the account is charged. Default:
	// true
	AutoRenew *bool

	// Whether you want to conceal contact information from WHOIS queries. If you
	// specify true, WHOIS ("who is") queries return contact information either for
	// Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If you specify false, WHOIS queries
	// return the information that you entered for the admin contact. Default: true
	PrivacyProtectAdminContact *bool

	// Reserved for future use.
	IdnLangCode *string

	// Provides detailed contact information. For information about the values that you
	// specify for each element, see ContactDetail
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ContactDetail.html).
	//
	// This member is required.
	RegistrantContact *types.ContactDetail

	// Provides detailed contact information. For information about the values that you
	// specify for each element, see ContactDetail
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ContactDetail.html).
	//
	// This member is required.
	AdminContact *types.ContactDetail

	// The domain name that you want to register. The top-level domain (TLD), such as
	// .com, must be a TLD that Route 53 supports. For a list of supported TLDs, see
	// Domains that You Can Register with Amazon Route 53
	// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/registrar-tld-list.html)
	// in the Amazon Route 53 Developer Guide. The domain name can contain only the
	// following characters:
	//
	//     * Letters a through z. Domain names are not case
	// sensitive.
	//
	//     * Numbers 0 through 9.
	//
	//     * Hyphen (-). You can't specify a
	// hyphen at the beginning or end of a label.
	//
	//     * Period (.) to separate the
	// labels in the name, such as the . in example.com.
	//
	// Internationalized domain
	// names are not supported for some top-level domains. To determine whether the TLD
	// that you want to use supports internationalized domain names, see Domains that
	// You Can Register with Amazon Route 53
	// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/registrar-tld-list.html).
	// For more information, see Formatting Internationalized Domain Names
	// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DomainNameFormat.html#domain-name-format-idns).
	//
	// This member is required.
	DomainName *string

	// Provides detailed contact information. For information about the values that you
	// specify for each element, see ContactDetail
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ContactDetail.html).
	//
	// This member is required.
	TechContact *types.ContactDetail
}

// The RegisterDomain response includes the following element.
type RegisterDomainOutput struct {

	// Identifier for tracking the progress of the request. To query the operation
	// status, use GetOperationDetail
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html).
	//
	// This member is required.
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRegisterDomainMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterDomain{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterDomain{}, middleware.After)
}

func newServiceMetadataMiddleware_opRegisterDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "RegisterDomain",
	}
}
